package config

import (
	"bufio"
	"os"
	"strings"
)

// LoadFile reads an ini file then return it as a map (associated array)
func LoadFile(filepath string) (map[string]string, error) {
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

		name, value := strings.TrimSpace(text[0:index]), strings.TrimSpace(text[index:])
		if value == "" || name == "" || name[0] == '#' || name[0] == ';' {
			continue
		}

		config[name] = value
	}

	return config, nil
}
