package httppconntime

import (
	"net"
	"net/http"
	"net/http/httptrace"
	"time"
)

type transport struct {
	roundTripper *http.Transport
	Reused bool
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.roundTripper == nil {
		t.roundTripper = &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       10 * time.Minute,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
	}
	return t.roundTripper.RoundTrip(req)
}

func (t *transport) GotConn(info httptrace.GotConnInfo) {
	t.Reused = info.Reused
}
