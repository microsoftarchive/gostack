package cors

import (
	"strings"
)

type Config struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
	ExposedHeaders []string
	MaxAge         string
}

func (config *Config) isOriginAllowed(origin string) bool {
	for _, option := range config.AllowedOrigins {
		switch option {
		case "*":
			return true
		case origin:
			return true
		}
	}
	return false
}

func (config *Config) isMethodAllowed(method string) bool {
	methods := config.AllowedMethods
	method = strings.ToUpper(method)
	if len(methods) == 0 || method == "OPTIONS" {
		return true
	}
	for _, option := range methods {
		if option == method {
			return true
		}
	}
	return false
}
