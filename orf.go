//
// oerc, alias oer-collector
// Copyright (C) 2021 emschu[aet]mailbox.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public
// License along with this program.
// If not, see <https://www.gnu.org/licenses/>.
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/alitto/pond"
	"github.com/gocolly/colly/v2"
	"gorm.io/gorm"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	orfHost                  = "tvthek.orf.at"
	orfHostWithPrefix        = "https://" + orfHost
	orfProgramHost           = "tv.orf.at"
	orfProgramHostWithPrefix = "https://" + orfProgramHost
)

var (
	orfTvShowLinkMatcher       = regexp.MustCompile(`^https?://tvthek.orf.at/.*`)
	orfProgramEntryLinkMatcher = regexp.MustCompile(`^https?://(tv|okidoki).orf.at/.*`)
	orfDailyProgramURLMatcher  = regexp.MustCompile(`^/program/[a-zA-Z0-9]+/index.*\.html$`)
)

// ParseORF central method to parse ORF tv show and program data
func ParseORF() {
	db, _ := getDb()

	// get channel family db record
	var channelFamily = getChannelFamily(db, "ORF")
	if channelFamily.ID == 0 {
		log.Fatalln("ORF channelFamily was not found!")
		return
	}

	// import tv shows
	if GetAppConf().EnableTVShowCollection {
		fetchTvShowsORF(db, channelFamily)
	}

	times := *generateDateRange(GetAppConf().DaysInPast, GetAppConf().DaysInFuture)

	// import program entries for the configured date range
	if GetAppConf().EnableProgramEntryCollection {
		pool := pond.New(4, 1000, pond.IdleTimeout(60*60*time.Second))
		for _, channel := range getChannelsOfFamily(db, channelFamily) {
			for _, day := range times {
				family := *channelFamily
				chn := channel
				dayToFetch := day

				pool.Submit(func() {
					handleDayORF(db, family, chn, dayToFetch)
				})
			}
		}
		// wait for finish
		pool.StopAndWait()
	}

	if verboseGlobal {
		log.Println("ORF parsed successfully")
	}
}

// fetchTvShowORF: This method checks all the tv shows
func fetchTvShowsORF(db *gorm.DB, family *ChannelFamily) {
	if !GetAppConf().EnableTVShowCollection || isRecentlyFetched() {
		log.Printf("Skip update of tv shows, due to recent fetch. Use 'forceUpdate' = true to ignore this.")
		return
	}

	c := orfCollector()

	c.OnHTML(".b-teaser", func(element *colly.HTMLElement) {
		var title, url string
		a := element.DOM.Find("a")
		titleInput, ex := a.Attr("title")
		if ex {
			title = trimAndSanitizeString(titleInput)
		}
		urlInput, ex2 := a.Attr("href")
		if ex2 {
			if !orfTvShowLinkMatcher.Match([]byte(urlInput)) {
				appLog(fmt.Sprintf("Unexpected orf tv show url detected for title '%s'. Was set to empty.", title))
				urlInput = ""
			}
			url = urlInput
		}
		if title == "" {
			appLog(fmt.Sprint("Empty title for orf tv show detected"))
			return
		}
		var hash = buildHash([]string{
			fmt.Sprintf("%d", int(family.ID)),
			title,
			"tv-show",
		})
		var tvShow = &TvShow{
			ManagedRecord: ManagedRecord{
				Title:           title,
				URL:             url,
				Hash:            hash,
				Homepage:        url,
				ChannelFamily:   *family,
				ChannelFamilyID: family.ID, // 4 = orf
			},
		}

		show := TvShow{}
		db.Where("hash = ?", hash).Find(&show)
		if show.ID != 0 {
			tvShow.ID = show.ID
		}
		saveTvShowRecord(db, tvShow)
	})

	err := c.Visit(orfHostWithPrefix + "/profiles")
	if err != nil {
		appLog(fmt.Sprintf("ORF parser fetch url error: %v\n", err))
	}
	c.Wait()
}

// handleDayORF method to fetch a single day of ORF
func handleDayORF(db *gorm.DB, family ChannelFamily, channel Channel, day time.Time) {
	queryURL := fmt.Sprintf("https://tv.orf.at/program/%s", channel.TechnicalID)

	c := orfCollector()
	dateDayDiff := getDaysBetween(day, time.Now())

	var programDetailURLPerDay string
	c.OnHTML("li.lane-item", func(dayElement *colly.HTMLElement) {
		if programDetailURLPerDay != "" {
			return
		}
		attr := dayElement.Attr("data-ds-day-index")
		dayIdx, err := strconv.ParseInt(attr, 10, 64)
		if err != nil {
			return
		}
		if int(dayIdx) == dateDayDiff {
			dayElement.ForEach("a", func(i int, linkElement *colly.HTMLElement) {
				s := linkElement.Attr("href")
				if orfDailyProgramURLMatcher.MatchString(s) {
					programDetailURLPerDay = s
				} else {
					appLog(fmt.Sprintf("Unexpected orf program url '%s'", trimAndSanitizeString(s)))
					return
				}
			})
		}
	})
	err := c.Visit(queryURL)
	if err != nil {
		appLog(fmt.Sprintf("Problem fetching orf url '%s'. Error: %v.", queryURL, err))
		return
	}
	c.Wait()

	if len(programDetailURLPerDay) == 0 {
		appLog("No daily program orf url found")
		return
	}

	c.OnHTML("li.broadcast", func(element *colly.HTMLElement) {
		title := trimAndSanitizeString(element.DOM.Find("div.series-title").Text())
		subTitle := trimAndSanitizeString(element.DOM.Find("div.episode-title").Text())
		detailLink := element.DOM.Find("div.series-title a")
		startTimeStr := trimAndSanitizeString(element.Attr("data-start-time"))
		endTimeStr := trimAndSanitizeString(element.Attr("data-end-time"))
		description := trimAndSanitizeString(element.DOM.Find("div.meta-data").Text())

		if len(subTitle) > 0 {
			title += " - " + subTitle
		}
		if title == "" {
			log.Printf("title empty, skip entry")
			return
		}
		if startTimeStr == "" || endTimeStr == "" {
			log.Printf("starttime or enddatetime empty, skipping entry")
			return
		}

		var url string
		url, _ = detailLink.Attr("href")
		url = fmt.Sprintf("%s%s", orfProgramHostWithPrefix, url)

		if len(url) > 0 && !orfProgramEntryLinkMatcher.MatchString(url) {
			appLog(fmt.Sprintf("Unexpected/Invalid orf program entry on page '%s'. Skipping entry with url '%s'.",
				trimAndSanitizeString(programDetailURLPerDay), trimAndSanitizeString(url)))
			return
		}

		var programEntry ProgramEntry
		programEntry.URL = trimAndSanitizeString(url)

		// handle start date time
		startDateTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			appLog(fmt.Sprint("Problem with parsing start date time in orf program entry.\n"))
			return
		}
		if startDateTime.IsZero() {
			appLog(fmt.Sprint("Problem with parsing start date time in orf program entry.\n"))
			return
		}

		// handle end date time
		endDateTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			appLog(fmt.Sprint("Problem with parsing start date time in orf program entry.\n"))
			return
		}
		if endDateTime.IsZero() {
			appLog(fmt.Sprint("Problem with parsing start date time in orf program entry.\n"))
			return
		}
		programEntry.StartDateTime = &startDateTime
		programEntry.EndDateTime = &endDateTime
		programEntry.DurationMinutes = int16(programEntry.EndDateTime.Sub(*programEntry.StartDateTime).Minutes())

		hash := buildHash([]string{
			startDateTime.String(),
			endDateTime.String(),
			title,
			url,
			strconv.Itoa(int(channel.ID)),
			strconv.Itoa(int(family.ID)),
			"program-entry",
		})
		programEntry.Hash = hash
		programEntry.TechnicalID = hash

		entry := ProgramEntry{}
		db.Model(&entry).Where("hash = ?", programEntry.Hash).Where("channel_id = ?", channel.ID).Preload("ImageLinks").Find(&entry)
		if entry.ID != 0 {
			if isRecentlyUpdated(&entry) {
				status.TotalSkippedPE++
				return
			}
			programEntry = entry
		}

		programEntry.Title = title
		programEntry.Description = description + "<br/>"
		programEntry.ChannelFamily = family
		programEntry.ChannelFamilyID = family.ID
		programEntry.Channel = channel
		programEntry.ChannelID = channel.ID

		if url != "" && !strings.Contains(url, "okidoki.orf.at") {
			requestHeaders := map[string]string{"Origin": "tv.orf.at", "Host": "tv.orf.at", "Accept": "text/html"}
			response, err := doGetRequest(url, requestHeaders, 3)
			if err != nil || response == nil {
				appLog(fmt.Sprintf("Problem fetching orf URL '%s'\n", url))
				return
			}
			// Load the HTML document
			reader := strings.NewReader(*response)
			doc, err := goquery.NewDocumentFromReader(reader)
			if err != nil {
				appLog(fmt.Sprintf("error fetching orf program entry: %v", err))
				return
			}
			desc := trimAndSanitizeString(doc.Find("div.document-content p.broadcast-programtext").Text())
			if len(desc) > 0 {
				programEntry.Description += desc
			}
			programEntry.Homepage = url

			doc.Find("div.broadcast-data a.broadcast-category").Each(func(i int, selection *goquery.Selection) {
				genre := trimAndSanitizeString(selection.Text())
				if len(genre) > 0 && len(genre) < 48 {
					considerTagExists(&programEntry, &genre)
				}
			})
		}

		// handle item images
		imageLinks := make([]ImageLink, 0)
		element.ForEach("figure.broadcast-image img", func(i int, imageElement *colly.HTMLElement) {
			imgLink := imageElement.Attr("src")
			imgLink = fmt.Sprintf("%s%s", orfProgramHostWithPrefix, imgLink)

			for _, existingImgLinkEntry := range programEntry.ImageLinks {
				if existingImgLinkEntry.URL == imgLink {
					// already existing
					return
				}
			}
			imageLinks = append(imageLinks, ImageLink{URL: imgLink})
		})
		programEntry.ImageLinks = append(programEntry.ImageLinks, imageLinks...)

		saveProgramEntryRecord(db, &programEntry)
	})

	err = c.Visit(fmt.Sprintf("%s%s", orfProgramHostWithPrefix, programDetailURLPerDay))
	if err != nil {
		appLog(fmt.Sprintf("Error of orf collector in url '%s': %v\n", queryURL, err))
	}
	c.Wait()
}

func getDaysBetween(day time.Time, now time.Time) int {
	first := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC)
	second := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	return int(first.Sub(second).Hours() / 24)
}

// helper method to get a collector instance
func orfCollector() *colly.Collector {
	collector := baseCollector([]string{orfHost, orfProgramHost})
	collector.Async = false
	return collector
}
