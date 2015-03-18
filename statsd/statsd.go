package statsd

import (
	"fmt"
	quipo "github.com/quipo/statsd"
	"github.com/wunderlist/gostack/common"
	"net/http"
	"time"
)

type Middleware struct {
	client *quipo.StatsdClient
}

func NewMiddleware(app string, env string) *Middleware {
	prefix := fmt.Sprintf("%s.%s.", app, env)
	client := quipo.NewStatsdClient("localhost:8125", prefix)
	client.CreateSocket()

	return &Middleware{
		client: client,
	}
}

func (m *Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(response, request)
	status := response.(common.ResponseWriter).Status()

	go (func() {
		duration := int64(time.Since(start) / time.Millisecond)
		m.client.Timing("request.time", duration)
		m.client.Incr("request.count", 1)
		m.client.Incr(fmt.Sprintf("%dxx.count", status/100), 1)
	})()
}
