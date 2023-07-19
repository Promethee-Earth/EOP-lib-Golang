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
		if key == "" || value == "" || key[0] == '#' || key[0] == ';' {
			continue
		}

		cfg.values[key] = value
	}

	return cfg, nil
}

// GetBoolean returns true if the value is "true", "yes", "on" or "1"
func (c Config) GetBoolean(key string) bool {
	var value = strings.ToLower(c.values[key])
	return value == "true" || value == "yes" || value == "on" || value == "1"
}

// GetValue returns the value of the key or the default value
func (c Config) GetValue(key string, defaultValue ...string) string {
	var value = c.values[key]
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}

// IsBetween checks if a key value is between min and max
func (c Config) IsBetween(key string, min, max int64) bool {
	var value, err = strconv.ParseInt(c.values[key], 10, 32)
	return err == nil && value >= min && value <= max
}

func (c Config) HasValue(key string, values ...string) bool {
	for _, value := range values {
		if c.values[key] == value {
			return true
		}
	}
	return false
}
