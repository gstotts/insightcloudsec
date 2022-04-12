package insightcloudsec

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the API client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setup() {
	// Test Server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Configure client to use with test server

	test_config := Config{
		BaseURL:    server.URL,
		ApiKey:     "whoami",
		Headers:    map[string][]string{},
		HTTPClient: &http.Client{},
	}
	client, _ = NewClient(&test_config)
}

func teardown() {
	server.Close()
}

func getJSONFile(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestClient_Headers(t *testing.T) {

	// Default headers should be set appropriately
	setup()
	mux.HandleFunc("/v2/public/user/info", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'POST', got %s", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "whoami", r.Header.Get("Api-Key"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))
		assert.Equal(t, "insightcloudsec-client", r.Header.Get("User-Agent"))
	})
	client.Users.CurrentUserInfo()
	teardown()
}
