package config

import "os"

func Config(key string, fallback interface{}) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback.(string)
	}

	return value
}
