package ping

import (
	"net/http"
)

// create own client to use in Pinger, cause Go statically-typed language
// and Go wait http.Client
type HTTPClient interface {
	Head(url string) (resp *http.Response, err error)
}

type Pinger struct {
	client HTTPClient
}

// Ping - ping url and check if url accessible
func (p Pinger) Ping(url string) bool {
	resp, err := p.client.Head(url)
	if err != nil {
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}
	return true
}
