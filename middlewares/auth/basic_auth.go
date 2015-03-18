package auth

import (
	"encoding/base64"
	"net/http"
)

type Middleware struct {
	hashes    []string
	whitelist []string
}

func NewMiddleware(users map[string]string, whitelist []string) *Middleware {
	var hashes []string
	for user, pass := range users {
		str := []byte(user + ":" + pass)
		encoded := base64.StdEncoding.EncodeToString(str)
		hashes = append(hashes, "Basic "+encoded)
	}
	return &Middleware{hashes, whitelist}
}

func (m *Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	isAuthed := false

	if len(m.whitelist) == 0 && len(m.hashes) == 0 {
		isAuthed = true
	}

	if !isAuthed {
		for _, ip := range m.whitelist {
			if request.RemoteAddr == ip {
				isAuthed = true
				continue
			}
		}
	}

	if !isAuthed {
		got := request.Header.Get("Authorization")
		for _, expected := range m.hashes {
			if got == expected {
				isAuthed = true
				continue
			}
		}
	}

	if !isAuthed {
		response.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization\"")
		http.Error(response, "Thou shall not pass", http.StatusUnauthorized)
	} else {
		next(response, request)
	}
}
