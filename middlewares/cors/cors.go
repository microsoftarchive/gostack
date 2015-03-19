package cors

import (
	"net/http"
	"strings"
)

func setCorsHeaders(response http.ResponseWriter, headers map[string]string) {
	for key, value := range headers {
		if value != "" {
			key = "Access-Control-" + key
			response.Header().Set(key, value)
		}
	}
}

type Middleware struct {
	config *Config
}

func NewMiddleware(config *Config) *Middleware {
	return &Middleware{config}
}

func (middleware *Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	origin := request.Header.Get("Origin")
	method := request.Header.Get("Access-Control-Request-Method")
	config := middleware.config

	isAllowedOrigin := origin != "" && config.isOriginAllowed(origin)
	isAllowedMethod := method != "" || config.isMethodAllowed(method)
	if isAllowedOrigin && isAllowedMethod {
		setCorsHeaders(response, map[string]string{
			"Allow-Credentials": "true",
			"Allow-Headers":     strings.Join(config.AllowedHeaders, ","),
			"Allow-Methods":     strings.Join(config.AllowedMethods, ","),
			"Allow-Origin":      origin,
			"Expose-Headers":    strings.Join(config.ExposedHeaders, ","),
			"Max-Age":           config.MaxAge,
		})
	}

	if request.Method == "OPTIONS" {
		response.WriteHeader(http.StatusNoContent)
		response.Write([]byte(""))
	} else {
		next(response, request)
	}
}
