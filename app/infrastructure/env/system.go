package env

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// GetEnvOrDefault gives env variable name or default value
func GetEnvOrDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return defaultValue
}

// GetEnvOrDefault gives env name
func GetEnvName() string {
	env := GetEnvOrDefault("ENV", "PROD")
	return strings.ToLower(env)
}

// FailOnError logs Fatal when erorr=
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
