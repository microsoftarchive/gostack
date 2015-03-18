package redirect

import (
	"fmt"
	"net/http"
)

type HttpsMiddleware struct {
	host string
}

func NewHttpsMiddleware(host string) *HttpsMiddleware {
	return &HttpsMiddleware{host}
}

func (m *HttpsMiddleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	host := request.Host
	url := request.URL
	status := 0
	proto := request.Header.Get("X-Forwarded-Proto")

	if host != m.host || (request.TLS == nil && proto != "https") {
		status = http.StatusMovedPermanently
	}

	if status > 0 {
		redirect_url := "https://" + m.host + url.RequestURI()
		response.Header().Set("Location", redirect_url)
		response.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		response.WriteHeader(status)
		if request.Method == "GET" {
			fmt.Fprintln(response, "<a href=\""+redirect_url+"\">redirecting</a>.\n")
		}
	} else {
		next(response, request)
	}
}
