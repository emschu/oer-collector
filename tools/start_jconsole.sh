#!/usr/bin/env bash

###
# #%L
# oer-collector-project
# %%
# Copyright (C) 2019 emschu[aet]mailbox.org
# %%
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
# 
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
# 
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.
# #L%
###

JDK_BIN_LOCATION=~/software/jdk-11/bin

if [ "$(jps -l | grep -c 'OerCollector')" -gt 0 ]
 then
    echo "Show JConsole of OerCollector $(jps -vl | grep -c 'OerCollector') collector processes"
    jps -l | grep 'OerCollector' | awk '{print $1}' | xargs "$JDK_BIN_LOCATION/jconsole" &
else
    echo "No running collector processes found"
fi