package bugsnag

import (
	"github.com/bugsnag/bugsnag-go"
	"github.com/bugsnag/bugsnag-go/errors"
	"net/http"
)

type Middleware struct {
	notifier *bugsnag.Notifier
}

func NewMiddleware(key string, env string) *Middleware {
	config := bugsnag.Configuration{
		APIKey:              key,
		ReleaseStage:        env,
		NotifyReleaseStages: []string{"production"},
		PanicHandler:        func() {},
	}
	notifier := bugsnag.New(config)
	return &Middleware{notifier}
}

func (m *Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			m.notifier.Notify(errors.New(err, 0), bugsnag.SeverityError, request)
		}
	}()

	next(response, request)
}
