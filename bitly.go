package bitly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client manages communication with the Bitly API..
type Client struct {
	AccessToken string
	APIURL      *url.URL
	HTTPClient  *http.Client

	Links *Links
}

// Response is a representation of the standard Bitly API reponse.
// It is left to the endpoint methods to decode Data as needed.
type Response struct {
	Data       json.RawMessage `json:"data"`
	StatusCode int             `json:"status_code"`
	StatusText string          `json:"status_txt"`
}

// Error response.
type Error struct {
	StatusCode int
	Summary    string
}

// Error string.
func (e *Error) Error() string {
	if e.StatusCode == 0 {
		return fmt.Sprintf("No status code given")
	}
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Summary)
}

// New returns a new Bitly client. A Bitly API access token must be
// provided. To create a token see https://dev.bitly.com/authentication.html.
func New(accessToken string) (client *Client) {
	apiURL, _ := url.Parse("https://api-ssl.bitly.com/v3")

	client = &Client{
		AccessToken: accessToken,
		APIURL:      apiURL,
	}

	client.HTTPClient = http.DefaultClient

	client.Links = &Links{client}

	return
}

// get sends a request to the Bitly API and checks for errors.
func (c *Client) get(path string, params url.Values) (res Response, err error) {
	params.Set("access_token", c.AccessToken)

	reqURL := *c.APIURL
	reqURL.Path += path
	reqURL.RawQuery = params.Encode()

	req, err := http.Get(reqURL.String())
	if err != nil {
		return
	}

	defer req.Body.Close()

	if req.StatusCode != 200 {
		err = &Error{
			StatusCode: req.StatusCode,
			Summary:    req.Status,
		}
		return
	}

	if err = json.NewDecoder(req.Body).Decode(&res); err != nil {
		return res, err
	}

	if res.StatusCode != 200 {
		err = &Error{
			StatusCode: res.StatusCode,
			Summary:    res.StatusText,
		}
		return
	}

	return
}
