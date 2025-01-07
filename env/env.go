package env

import (
	"os"
	"strings"
)

var globalEnvPrefix = ""

func SetEnvPrefix(prefix string) {
	globalEnvPrefix = prefix
}

func GetString(key string) (string, bool) {
	return os.LookupEnv(globalEnvPrefix + key)
}

func GetStringDefault(key, defaultValue string) string {
	if value, ok := GetString(key); ok {
		return value
	}
	return defaultValue
}

func GetBool(key string) (bool, bool) {
	if val, ok := GetString(key); ok {
		switch strings.ToLower(val) {
		case "true", "1", "yes", "on":
			return true, true
		case "false", "0", "no", "off":
			return false, true
		}
	}
	return false, false
}

func GetBoolDefault(key string, defaultValue bool) bool {
	if value, ok := GetBool(key); ok {
		return value
	}
	return defaultValue
}
