package config

import (
	"bufio"
	"os"
	"strings"
)

// LoadIniFile reads a simple ini file then return it as a key-value map
func LoadIniFile(filepath string) (map[string]string, error) {
	config := make(map[string]string)

	file, err := os.Open(filepath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		index := strings.IndexRune(text, '=')
		if index < 1 {
			continue
		}

		key, value := strings.TrimSpace(text[0:index]), strings.TrimSpace(text[index:])
		if key == "" || value == "" || key[0] == '#' || key[0] == ';' {
			continue
		}

		config[key] = value
	}

	return config, nil
}
