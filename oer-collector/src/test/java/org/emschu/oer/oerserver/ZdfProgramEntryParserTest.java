package org.emschu.oer.oerserver;

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

import junit.framework.TestCase;
import org.junit.Assert;
import org.junit.Test;
import org.emschu.oer.core.model.Channel;
import org.emschu.oer.collector.reader.parser.zdf.ProgramEntryParser;

public class ZdfProgramEntryParserTest extends TestCase {

    @Test
    public void testProgramParser() {
        ProgramEntryParser.ZDFScraper zdfScraper = new ProgramEntryParser.ZDFScraper(getDummyZdfChannel());
        final String zdfApiKey = zdfScraper.retrieveApiKey();
        Assert.assertNotNull(zdfApiKey);
        Assert.assertNotEquals("", zdfApiKey);
        Assert.assertFalse(zdfApiKey.contains("'"));
    }

    protected Channel getDummyZdfChannel() {
        Channel channel = new Channel();
        channel.setAdapterFamily(Channel.AdapterFamily.ZDF);
        channel.setChannelKey(Channel.ChannelKey.ZDF);
        channel.setTechnicalId("zdf");
        channel.setName("ZDF");
        return channel;
    }
}