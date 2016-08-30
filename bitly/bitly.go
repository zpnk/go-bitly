package bitly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const defaultBaseURL = "https://api-ssl.bitly.com/v3"

// Client manages communication with the Bitly API..
type Client struct {
	AccessToken string
	BaseURL     *url.URL
	HTTPClient  *http.Client

	Link *Link
}

// NewClient returns a new Bitly client. A Bitly API access token must be
// provided. To create a token see https://dev.bitly.com/authentication.html.
func NewClient(accessToken string) (c *Client) {
	baseURL, _ := url.Parse(defaultBaseURL)

	c = &Client{AccessToken: accessToken, BaseURL: baseURL}
	c.HTTPClient = http.DefaultClient
	c.Link = &Link{c}

	return
}

// Response is a Bitly API response. It is used for decoding the raw json
// response. Data is specificed as map[string]interface{} because the
// corresponding json field returns arbitrary values based on the API endpoint.
// It is left up to the endpoint methods to decode Data further.
type Response struct {
	StatusCode int                    `json:"status_code"`
	StatusTxt  string                 `json:"status_txt"`
	Data       map[string]interface{} `json:"data"`
}

// Error is a Bitly API error. It is used for reporting api error responses.
type Error Response

func (e *Error) Error() string {
	if e.StatusCode == 0 {
		return fmt.Sprintf("No status code given")
	}
	return fmt.Sprintf("%d: %s", e.StatusCode, e.StatusTxt)
}

// Params are the query parameters to be appended to the request url. It is
// specified as an empty interface given the differnt queries allowed per
// endpoint. It is left up to the endpoint methods to marshall user input into
// a struct cotaining the allowed values.
type Params interface{}

// Get sends an API request and returns the API response. The response is JSON
// decoded and stored in a Response struct. The status code contained in the
// response is checked and in the event of an error code, the response is
// stored in an Error struct and returned.
func (c *Client) Get(path string, params Params) (res *Response, err error) {
	url, err := c.buildURL(path, params)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := c.HTTPClient.Do(req)
	if err != nil {
		return
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		err = (*Error)(res)
		return
	}

	return
}

// buildURL adds a given path and query params to the client's base url. The
// client's access token is always appended to the query. The resulting url.URL
// is returned for use in the request.
func (c *Client) buildURL(path string, params Params) (u *url.URL, err error) {
	ref, err := url.Parse(path)
	if err != nil {
		return
	}

	c.BaseURL.Path = c.BaseURL.EscapedPath() + ref.EscapedPath()
	u = c.BaseURL

	query, err := query.Values(params)
	if err != nil {
		return
	}
	query.Set("access_token", c.AccessToken)

	u.RawQuery = query.Encode()

	return
}
