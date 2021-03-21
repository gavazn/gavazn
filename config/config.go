package config

import (
	"bufio"
	"os"
	"strings"
)

var conf = load()

func load() map[string]interface{} {
	file, err := os.Open(".env")
	if err != nil {
		return nil
	}
	defer file.Close()

	configMap := map[string]interface{}{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")

		key := strings.Trim(line[0], " ")
		value := strings.Trim(line[1], " ")

		configMap[key] = value
	}

	return configMap
}

// Get environment variable
func Get(key string) string {
	value, exist := conf[key]
	if !exist {
		return ""
	}

	return value.(string)
}
