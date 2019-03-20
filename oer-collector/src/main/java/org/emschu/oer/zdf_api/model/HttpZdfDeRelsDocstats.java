
package org.emschu.oer.zdf_api.model;

/*-
 * #%L
 * oer-server
 * %%
 * Copyright (C) 2019 emschu[aet]mailbox.org
 * %%
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 * #L%
 */

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class HttpZdfDeRelsDocstats {

    @SerializedName("http://zdf.de/rels/target")
    @Expose
    private String httpZdfDeRelsTarget;
    @SerializedName("profile")
    @Expose
    private String profile;
    @SerializedName("self")
    @Expose
    private String self;
    @SerializedName("avgRating")
    @Expose
    private Double avgRating;
    @SerializedName("viewCount")
    @Expose
    private Integer viewCount;
    @SerializedName("ratingCount")
    @Expose
    private Integer ratingCount;
    @SerializedName("avgRatingCount")
    @Expose
    private Integer avgRatingCount;

    public String getHttpZdfDeRelsTarget() {
        return httpZdfDeRelsTarget;
    }

    public void setHttpZdfDeRelsTarget(String httpZdfDeRelsTarget) {
        this.httpZdfDeRelsTarget = httpZdfDeRelsTarget;
    }

    public String getProfile() {
        return profile;
    }

    public void setProfile(String profile) {
        this.profile = profile;
    }

    public String getSelf() {
        return self;
    }

    public void setSelf(String self) {
        this.self = self;
    }

    public Double getAvgRating() {
        return avgRating;
    }

    public void setAvgRating(Double avgRating) {
        this.avgRating = avgRating;
    }

    public Integer getViewCount() {
        return viewCount;
    }

    public void setViewCount(Integer viewCount) {
        this.viewCount = viewCount;
    }

    public Integer getRatingCount() {
        return ratingCount;
    }

    public void setRatingCount(Integer ratingCount) {
        this.ratingCount = ratingCount;
    }

    public Integer getAvgRatingCount() {
        return avgRatingCount;
    }

    public void setAvgRatingCount(Integer avgRatingCount) {
        this.avgRatingCount = avgRatingCount;
    }
}
