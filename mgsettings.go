// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://www.mugomes.com.br

package mgsettings

import (
	"encoding/json"
	"os"
	"path"
)

type MGSETTINGS struct {
	pathfile string
	data     map[string]json.RawMessage
}

func Load(nameApp string, defaultPathHome bool) (*MGSETTINGS, error) {
	m := &MGSETTINGS{
		data: make(map[string]json.RawMessage),
	}

	var pathConfig string
	if defaultPathHome {
		home, _ := os.UserHomeDir()
		pathConfig = path.Join(home, "."+nameApp)
	} else {
		pathConfig = nameApp
	}

	if err := os.MkdirAll(pathConfig, 0755); err != nil {
		return nil, err
	}

	m.pathfile = path.Join(pathConfig, "config.json")

	if raw, err := os.ReadFile(m.pathfile); err == nil {
		_ = json.Unmarshal(raw, &m.data)
	}

	return m, nil
}

func (m *MGSETTINGS) set(key string, v any) {
	raw, _ := json.Marshal(v)
	m.data[key] = raw
}

func (m *MGSETTINGS) get(key string, out any) bool {
	raw, ok := m.data[key]
	if !ok {
		return false
	}
	return json.Unmarshal(raw, out) == nil
}

func (m *MGSETTINGS) SetString(key, value string) {
	m.set(key, value)
}

func (m *MGSETTINGS) SetInt(key string, value int) {
	m.set(key, value)
}

func (m *MGSETTINGS) SetBool(key string, value bool) {
	m.set(key, value)
}

func (m *MGSETTINGS) SetStringSlice(key string, value []string) {
	m.set(key, value)
}

func (m *MGSETTINGS) GetString(key string, def string) string {
	var v string
	if m.get(key, &v) {
		return v
	}
	return def
}

func (m *MGSETTINGS) GetInt(key string, def int) int {
	var v int
	if m.get(key, &v) {
		return v
	}
	return def
}

func (m *MGSETTINGS) GetBool(key string, def bool) bool {
	var v bool
	if m.get(key, &v) {
		return v
	}
	return def
}

func (m *MGSETTINGS) GetStringSlice(key string, def []string) []string {
	var v []string
	if m.get(key, &v) {
		return v
	}
	return def
}

func (m *MGSETTINGS) Save() error {
	raw, err := json.MarshalIndent(m.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(m.pathfile, raw, 0644)
}
