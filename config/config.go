package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	values map[string]string
}

// LoadIniFile reads a simple ini file then return a key-value config
func LoadIniFile(filepath string) (Config, error) {
	var cfg = Config{values: make(map[string]string)}

	var file, err = os.Open(filepath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var text = scanner.Text()
		var index = strings.IndexRune(text, '=')
		if index < 1 {
			continue
		}

		var key, value = strings.TrimSpace(text[0:index]), strings.TrimSpace(text[index+1:])
		if key != "" && value != "" && key[0] != '#' && key[0] != ';' && key[0] != '[' {
			cfg.values[key] = value
		}
	}

	return cfg, nil
}

// Get returns the value of the key or the default value
func (c Config) Get(key string, defaultValue ...string) string {
	var value = c.values[key]
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

// GetBool returns true if the value is "yes", "on", "1" or "true" (case insensitive)
func (c Config) GetBool(key string) bool {
	var value = strings.ToLower(c.values[key])
	return value == "true" || value == "yes" || value == "on" || value == "1"
}

// IsBetween checks if a key value number is between min and max (included)
func (c Config) IsBetween(key string, min, max int64) bool {
	var value, err = strconv.ParseInt(c.values[key], 10, 32)
	return err == nil && value >= min && value <= max
}
