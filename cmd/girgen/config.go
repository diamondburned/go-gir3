/*
 * Copyright (C) 2019 ~ 2020 Uniontech Software Technology Co.,Ltd
 *
 * Author:
 *
 * Maintainer:
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	Black     []string `json:"black"`
	CIncludes []string `json:"cIncludes"`
}

func loadConfig(filename string, cfg *config) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, cfg)
}

type genState struct {
	PrevNamespace string
	FuncNextId    int
	GetTypeNextId int
}

func loadGenState(filename string, genState *genState) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, genState)
}

func saveGenState(filename string, genState *genState) error {
	data, err := json.Marshal(genState)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
