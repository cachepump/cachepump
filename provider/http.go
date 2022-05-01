package provider

import (
	"github.com/cachepump/cachepump/cache"
	"fmt"
	"io"
	"net/http"
	"strings"

	logger "github.com/AntonYurchenko/log-go"
)

// HttpSource define available fields for source with type = http.
type HttpSource struct {
	Endpoint string      `yaml:"endpoint"`
	Method   string      `yaml:"method"`
	Auth     Auth        `yaml:"auth"`
	Header   http.Header `yaml:"header"`
	Body     string      `yaml:"body"`
}

// Auth is a structure for basic authentication.
type Auth struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// String serialisation a structure Auth to string.
func (a Auth) String() string {
	return fmt.Sprintf(`{User:%[1]s Password:%[2]s}`, a.User, strings.Repeat("*", len(a.Password)))
}

// HttpPump generate an job function for updating data from http sources.
func HttpPump(src *HttpSource, name string) func() {

	if src == nil {
		logger.Error("Invalid source with name %q. Section is not defined", name)
		return func() {}
	}

	return func() {

		req, err := newRequest(src.Method, src.Endpoint, src.Body, name, src.Auth, src.Header)
		if err != nil {
			logger.ErrorF("I cannot create http request, source name: %[1]q error: %[2]v", name, err)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.ErrorF("I cannot recieve response, source name: %[1]q, error: %[2]v", name, err)
			return
		}
		defer resp.Body.Close()

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.ErrorF("I cannot read response, source name: %[1]q, error: %[2]v", name, err)
			return
		}

		err = cache.Set(name, bytes)
		if err != nil {
			logger.ErrorF("I cannot save response body to cache, source name: %[1]q, error: %[2]v", name, err)
			return
		}
		logger.InfoF("Data from source %q has been updated.", name)
	}
}

// newRequest creates a new http request.
func newRequest(method, endpoint, body, sourceName string, auth Auth, header http.Header) (req *http.Request, err error) {
	req, err = http.NewRequest(method, endpoint, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
		logger.DebugF("Set header: %+[1]v for source %[2]q", header, sourceName)
	}
	if auth != (Auth{}) {
		req.SetBasicAuth(auth.User, auth.Password)
		logger.DebugF("Set basic auth: '%[1]v' for source %[2]q", auth, sourceName)
	}
	return req, nil
}
