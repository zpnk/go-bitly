package bitly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux
	// client is the GitHub client being tested.
	client *Client
	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a bitly.Client that is
// configured to talk to that test server.  Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = New("123")
	url, _ := url.Parse(server.URL)
	client.APIURL = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

// jsonRes builds a json string indentical to the Bitly json responses.
// Tests can use this to generate mock json responses for the test server.
func jsonRes(data string, code int, txt string) string {
	return fmt.Sprintf(`{
		"data": %s,
		"status_code": %d,
		"status_txt": "%s"
	}`, data, code, txt)
}

func newEndpoint(endpoint string, response string) {
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	})
}

func TestClient_Get_badJSON(t *testing.T) {
	setup()
	defer teardown()

	newEndpoint("/foo", `not_json`)

	_, err := client.Get("/foo", url.Values{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*json.SyntaxError); !ok {
		t.Errorf("Expected JSON syntax error, got %+v", err)
	}
}

func TestClient_Get_serverError(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.Get("/foo123", url.Values{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*Error); !ok {
		t.Errorf("Expected response error, got %+v", err)
	}

	got := err.Error()
	want := "404: 404 Not Found"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error returned %v, want %v", got, want)
	}
}

func TestClient_Get_bitlyError(t *testing.T) {
	setup()
	defer teardown()

	newEndpoint("/foo", jsonRes("{}", 403, "RATE_LIMIT_EXCEEDED"))

	_, err := client.Get("/foo", url.Values{})
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*Error); !ok {
		t.Errorf("Expected response error, got %+v", err)
	}

	got := err.Error()
	want := "403: RATE_LIMIT_EXCEEDED"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error returned %v, want %v", got, want)
	}
}
