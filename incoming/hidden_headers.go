package incoming

import (
	"net/http"
	"strings"
)

type HiddenHeadersMiddleware struct{}

func NewHiddenHeadersMiddleware() *HiddenHeadersMiddleware {
	return &HiddenHeadersMiddleware{}
}

func (m *HiddenHeadersMiddleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	headers := request.Header

	// remove private or SPDY headers in the incoming request
	for key, _ := range headers {
		if strings.HasPrefix(key, "X-Hidden-") || strings.HasPrefix(key, ":") {
			headers.Del(key)
		}
	}

	next(response, request)
}
