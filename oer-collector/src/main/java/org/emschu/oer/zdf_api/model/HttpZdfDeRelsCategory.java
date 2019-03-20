
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

public class HttpZdfDeRelsCategory extends AbstractZdfApiElement {

    @SerializedName("title")
    @Expose
    private String title;
    @SerializedName("http://zdf.de/rels/target")
    @Expose
    private HttpZdfDeRelsTarget httpZdfDeRelsTarget;

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public HttpZdfDeRelsTarget getHttpZdfDeRelsTarget() {
        return httpZdfDeRelsTarget;
    }

    public void setHttpZdfDeRelsTarget(HttpZdfDeRelsTarget httpZdfDeRelsTarget) {
        this.httpZdfDeRelsTarget = httpZdfDeRelsTarget;
    }
}
