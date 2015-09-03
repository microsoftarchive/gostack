package jsonlogger

import (
	"encoding/json"
	"fmt"
	"github.com/wunderlist/gostack/common"
	"net/http"
	"time"
)

type Middleware struct {
	app string
	env string
}

func NewMiddleware(app string, env string) *Middleware {
	return &Middleware{app: app, env: env}
}

type Log struct {
	Type           string `json:"type"`
	V              string `json:"v"`
	App            string `json:"app"`
	ClientId       string `json:"client_id,omitempty"`
	DeviceId       string `json:"device_id,omitempty"`
	Env            string `json:"env"`
	ForwardedFor   string `json:"forwarded_for,omitempty"`
	Host           string `json:"host"`
	Method         string `json:"method"`
	Platform       string `json:"platform,omitempty"`
	Product        string `json:"product,omitempty"`
	ProductVersion string `json:"product_version,omitempty"`
	RemoteAddress  string `json:"remote_address"`
	Status         int    `json:"status"`
	System         string `json:"system,omitempty"`
	SystemVersion  string `json:"system_version,omitempty"`
	TotalTime      int64  `json:"total_time"`
	URI            string `json:"uri"`
	UserAgent      string `json:"user_agent,omitempty"`
	UserId         string `json:"user_id,omitempty"`
	UUID           string `json:"uuid,omitempty"`
}

func (m *Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	start := time.Now()
	log := &Log{
		Type:          "log",
		V:             "1",
		App:           m.app,
		Env:           m.env,
		Method:        request.Method,
		URI:           request.RequestURI,
		Host:          request.Host,
		RemoteAddress: request.RemoteAddr,
		ForwardedFor:  request.Header.Get("X-Forwarded-For"),
		UserAgent:     request.UserAgent(),
	}

	next(response, request)

	log.Status = response.(common.ResponseWriter).Status()
	log.TotalTime = int64(time.Since(start) / time.Millisecond)

	go (func() {
		data, _ := json.Marshal(&log)
		fmt.Println(string(data), request.Header)
	})()
}
