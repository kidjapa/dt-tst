package config

import (
	"errors"
	"os"
	"strconv"
)

var (
	ErrEnvVarEmpty = errors.New("getEnv: environment variable empty")
)

// GetPortEnv Return the Environment variable "ENV_PORT"
func GetPortEnv() int64 {
	if os.Getenv("ENV_PORT") != "" {
		if port, err := getEnvInt64("ENV_PORT"); err == nil {
			return port
		} else {
			return 9090
		}
	}
	return 9090
}


func getEnvInt64(key string) (int64, error) {
	v := os.Getenv(key)
	if v == "" {
		return -1, ErrEnvVarEmpty
	}
	if i, e := strconv.ParseInt(v, 10, 64); e == nil {
		return i, nil
	} else {
		return -1, e
	}
}
