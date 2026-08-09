// Required Notice: Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

package mgsettings

import (
	"encoding/json"
	"os"
	"path"
)

type mgsettings struct {
	pathfile string
	data     map[string]json.RawMessage
}

func Load(nameApp string, defaultPathHome bool) (*mgsettings, error) {
	m := &mgsettings{
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

func (m *mgsettings) set(key string, v any) {
	raw, _ := json.Marshal(v)
	m.data[key] = raw
}

func (m *mgsettings) get(key string, out any) bool {
	raw, ok := m.data[key]
	if !ok {
		return false
	}
	return json.Unmarshal(raw, out) == nil
}

func (m *mgsettings) SetString(key, value string) {
	m.set(key, value)
}

func (m *mgsettings) SetInt(key string, value int) {
	m.set(key, value)
}

func (m *mgsettings) SetBool(key string, value bool) {
	m.set(key, value)
}

func (m *mgsettings) SetStringSlice(key string, value []string) {
	m.set(key, value)
}

func (m *mgsettings) GetString(key string, def string) string {
	var v string
	if m.get(key, &v) {
		return v
	}
	return def
}

func (m *mgsettings) GetInt(key string, def int) int {
	var v int
	if m.get(key, &v) {
		return v
	}
	return def
}

func (m *mgsettings) GetBool(key string, def bool) bool {
	var v bool
	if m.get(key, &v) {
		return v
	}
	return def
}

func (m *mgsettings) GetStringSlice(key string, def []string) []string {
	var v []string
	if m.get(key, &v) {
		return v
	}
	return def
}

func (m *mgsettings) Save() error {
	raw, err := json.MarshalIndent(m.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(m.pathfile, raw, 0644)
}
