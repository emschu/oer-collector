
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

public class ZdfImage {

    @SerializedName("altText")
    @Expose
    private String altText;
    @SerializedName("profile")
    @Expose
    private String profile;
    @SerializedName("source")
    @Expose
    private String source;
    @SerializedName("layouts")
    @Expose
    private Layouts layouts;

    public String getAltText() {
        return altText;
    }

    public void setAltText(String altText) {
        this.altText = altText;
    }

    public ZdfImage withAltText(String altText) {
        this.altText = altText;
        return this;
    }

    public String getProfile() {
        return profile;
    }

    public void setProfile(String profile) {
        this.profile = profile;
    }

    public ZdfImage withProfile(String profile) {
        this.profile = profile;
        return this;
    }

    public String getSource() {
        return source;
    }

    public void setSource(String source) {
        this.source = source;
    }

    public ZdfImage withSource(String source) {
        this.source = source;
        return this;
    }

    public Layouts getLayouts() {
        return layouts;
    }

    public void setLayouts(Layouts layouts) {
        this.layouts = layouts;
    }

    public ZdfImage withLayouts(Layouts layouts) {
        this.layouts = layouts;
        return this;
    }

}
