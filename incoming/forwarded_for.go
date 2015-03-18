package incoming

import (
	"net"
	"net/http"
	"strings"
)

type ForwardedForMiddleware struct{}

func NewForwardedForMiddleware() *ForwardedForMiddleware {
	return &ForwardedForMiddleware{}
}

func (m *ForwardedForMiddleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	headers := request.Header
	address := headers.Get("X-Forwarded-For")

	if address != "" {
		address = strings.TrimSpace(strings.Split(address, ",")[0])
		request.RemoteAddr = strings.TrimSpace(strings.Split(address, ":")[0])
	} else {
		request.RemoteAddr, _, _ = net.SplitHostPort(request.RemoteAddr)
	}

	next(response, request)
}
