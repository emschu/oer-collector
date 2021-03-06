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

// database struct(ure) models

import (
	"time"
)

// BaseModel base model for all entities
type BaseModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ChannelFamily entity
type ChannelFamily struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Title string `gorm:"size:32" json:"title"`
}

// ManagedRecord base entity part, meta-struct used by several other db related structs
type ManagedRecord struct {
	Title           string        `gorm:"size:500" json:"title"`
	URL             string        `gorm:"size:1500" json:"url"`
	Hash            string        `gorm:"type:varchar(32);unique_index;not null" json:"hash"`
	TechnicalID     string        `gorm:"type:varchar(64);null" json:"technical_id"`
	Homepage        string        `gorm:"size:1500" json:"homepage"`
	ChannelFamily   ChannelFamily `json:"-"`
	ChannelFamilyID uint          `gorm:"index" json:"channel_family_id"`
}

// Channel entity
type Channel struct {
	BaseModel
	ManagedRecord
}

// TvShow entity
type TvShow struct {
	BaseModel
	ManagedRecord

	RelatedProgramEntries []ProgramEntry `gorm:"many2many:tv_show_program_entries" json:"-"`
}

// ProgramEntry entity
type ProgramEntry struct {
	BaseModel
	ManagedRecord

	StartDateTime   *time.Time  `gorm:"index,not null" json:"start_date_time"`
	EndDateTime     *time.Time  `gorm:"index,not null" json:"end_date_time"`
	LastCheck       *time.Time  `json:"last_check"`
	DurationMinutes int16       `json:"duration_in_minutes"`
	Description     string      `gorm:"size:30000" json:"description"`
	Channel         Channel     `json:"-"`
	ChannelID       uint        `gorm:"index,not null" json:"channel_id"`
	Tags            string      `gorm:"size:256" json:"tags"`
	ImageLinks      []ImageLink `gorm:"foreignKey:ProgramEntryID" json:"image_links"`
}

// ImageLink entity
type ImageLink struct {
	BaseModel

	URL            string `gorm:"size:1024" json:"url"`
	ProgramEntryID uint   `gorm:"index" json:"program_entry_id"`
}

// LogEntry used to store parsing errors
type LogEntry struct {
	BaseModel

	Message string `gorm:"size:1536" json:"message"`
}

// Settings used to store values
type Settings struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	SettingKey string `gorm:"type:varchar(64)" json:"key"`
	Value      string `gorm:"type:varchar(1024)" json:"value"`
}

// Recommendation store program recommendations
type Recommendation struct {
	ID             uint         `gorm:"primary_key" json:"id"`
	CreatedAt      *time.Time   `json:"created_at"`
	ProgramEntry   ProgramEntry `json:"program_entry"`
	ProgramEntryID uint         `gorm:"unique_index;not null" json:"program_entry_id"`
	ChannelID      uint         `json:"channel_id"`
	StartDateTime  *time.Time   `json:"start_date_time"`
	Keywords       string       `gorm:"size:500" json:"keywords"`
}
