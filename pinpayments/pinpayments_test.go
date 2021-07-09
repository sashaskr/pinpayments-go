package pinpayments

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	tMux    *http.ServeMux
	tServer *httptest.Server
	tClient *Client
	tConf   *Config
)

func setup() {
	tMux = http.NewServeMux()
	tServer = httptest.NewServer(tMux)
	tConf = NewConfig(true, APIEnvSecret, APIEnvKey)
	tClient, _ = NewClient(nil, tConf)
	u, _ := url.Parse(tServer.URL + "/")
	tClient.BaseURL = u
}

func teardown() {
	tServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}