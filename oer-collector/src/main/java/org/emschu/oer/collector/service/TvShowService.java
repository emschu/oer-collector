package org.emschu.oer.collector.service;

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

import org.emschu.oer.core.model.ProgramEntry;
import org.emschu.oer.core.model.TvShow;
import org.emschu.oer.core.model.repository.TvShowRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class TvShowService {

    private final TvShowRepository tvShowRepository;

    @Autowired
    public TvShowService(TvShowRepository tvShowRepository) {
        this.tvShowRepository = tvShowRepository;
    }

    public synchronized void linkProgramEntryWithZdfBrandId(String brandId, ProgramEntry entry) {
        TvShow tvShow = tvShowRepository.findByAdditionalId(brandId);
        // only add if not already linked!
        if (tvShow != null) {
            if (!tvShow.getRelatedProgramEntries().contains(entry)) {
                tvShow.getRelatedProgramEntries().add(entry);
                tvShowRepository.save(tvShow);
            }
        }
    }
}
