// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://www.mugomes.com.br

package mgsettings

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type MGSETTINGS struct {
	pathfile string
	data     map[string]any
}

func Load(nameApp string, defaultPathHome bool) *MGSETTINGS {
	mgsettings := &MGSETTINGS{}
	mgsettings.data = make(map[string]any)

	var pathHome string
	var pathConfig string
	if defaultPathHome {
		pathHome, _ = os.UserHomeDir()
		pathConfig = path.Join(pathHome, fmt.Sprintf(".%s", nameApp))
	} else {
		pathConfig = nameApp
	}

	_, errDir := os.Stat(pathConfig)
	if errDir != nil {
		os.Mkdir(pathConfig, os.ModePerm)
	}

	mgsettings.pathfile = path.Join(pathConfig, "config.json")
	_, errFile := os.Stat(mgsettings.pathfile)

	if errFile == nil {
		var data map[string]any
		loadFile, _ := os.ReadFile(mgsettings.pathfile)
		err := json.Unmarshal(loadFile, &data)
		if err == nil {
			mgsettings.data = data
		}
	}

	return mgsettings
}

func (mgsettings *MGSETTINGS) Set(name string, value any) {
	mgsettings.data[name] = value
}

func (mgsettings *MGSETTINGS) Get(name string, defaultValue any) any {
	valor, ok := mgsettings.data[name]
	if ok {
		return valor
	} else {
		switch r := defaultValue.(type) {
		case int, int64:
			return float64(r.(int))
		case float32, float64:
			return float64(r.(int))
		default:
			return r
		}
	}
}

func (mgsettings *MGSETTINGS) Save() {
	json, _ := json.Marshal(mgsettings.data)
	os.WriteFile(mgsettings.pathfile, json, os.ModePerm)
}
