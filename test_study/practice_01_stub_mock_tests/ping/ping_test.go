package ping

import (
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

// Head - parse url and return response with error
func (c *MockClient) Head(url string) (resp *http.Response, err error) {
	// parse url
	parts := strings.Split(url, "/")
	statusCode, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return nil, err
	}
	resp = &http.Response{StatusCode: statusCode}
	return resp, nil
}

func TestPing(t *testing.T) {
	client := &MockClient{}
	pinger := Pinger{client}
	result := pinger.Ping("https://example.com/200")
	assert.True(t, result, "True statuscode")
}
