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

	client = NewClient("123")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

// buildJSONRes builds a json string indentical to the Bitly json responses.
// Tests can use this to generate mock json responses for the test server.
func buildJSONRes(data string, code int, txt string) string {
	return fmt.Sprintf(`{
		"data": %s,
		"status_code": %d,
		"status_txt": "%s"
	}`, data, code, txt)
}

func TestClient_get_badURL(t *testing.T) {
	c := NewClient("123")
	_, err := c.Get(":", nil)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestClient_get_badParams(t *testing.T) {
	c := NewClient("123")
	_, err := c.Get("/link", "123")
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
}

func TestClient_get_badJSON(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `not_json`)
	})

	_, err := client.Get("/foo", nil)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*json.SyntaxError); !ok {
		t.Errorf("Expected JSON syntax error, got %+v", err)
	}
}

func TestClient_get_errorRes(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, buildJSONRes("{}", 404, "NOT_FOUND"))
	})

	_, err := client.Get("/foo", nil)
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*Error); !ok {
		t.Errorf("Expected response error, got %+v", err)
	}

	got := err.Error()
	want := "404: NOT_FOUND"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error returned %v, want %v", got, want)
	}
}
