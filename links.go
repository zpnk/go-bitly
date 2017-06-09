package bitly

import (
	"encoding/json"
	"net/url"
)

// Links handles communication with the link related Bitly API endpoints.
type Links struct {
	*Client
}

// Link represents the data returned from link endpoints.
type Link struct {
	ShortURL      string `json:"short_url"`
	LongURL       string `json:"long_url"`
	GlobalHash    string `json:"global_hash"`
	UserHash      string `json:"user_hash"`
	Hash          string `json:"hash"`
	NewHash       int    `json:"new_hash"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	AggregateLink string `json:"aggregate_link"`
	CreatedAt     int    `json:"created_at"`
	CreatedBy     string `json:"created_by"`
}

// req wraps Client#get and unpacks the response specifically for Links methods.
func (client *Links) req(path string, params url.Values, key string) (links []Link, err error) {
	req, err := client.get(path, params)
	if err != nil {
		return
	}

	res := map[string][]Link{}
	err = json.Unmarshal(req.Data, &res)
	if err != nil {
		return
	}

	return res[key], err
}

// Expand returns the long urls for a given set short urls.
// Returns slice of:
//
//	bitly.Link{
//		GlobalHash: "1RmnUT",
//		LongURL:    "http://google.com",
//		ShortURL:   "http://bit.ly/1RmnUT",
//		UserHash:   "1RmnUT",
//	}
//
// Bitly API docs: http://dev.bitly.com/links.html#v3_expand
func (client *Links) Expand(urls ...string) (links []Link, err error) {
	return client.req("/expand", url.Values{"shortUrl": urls}, "expand")
}

// Info returns the page title and other metadata for a given set of short urls.
// Returns slice of:
//
//	bitly.Link{
//		GlobalHash: "1RmnUT",
//		ShortURL:   "http://bit.ly/1RmnUT",
//		Title:      "Google",
//		UserHash:   "1RmnUT",
//		CreatedAt:  1212926400,
//	}
//
// Bitly API docs: http://dev.bitly.com/links.html#v3_info
func (client *Links) Info(urls ...string) (links []Link, err error) {
	return client.req("/info", url.Values{"shortUrl": urls}, "info")
}

// Lookup queries for bitlink(s) mapping to the given url(s).
// Returns slice of:
//
//	bitly.Link{
//		URL:           "http://www.google.com/",
//		AggregateLink: "http://www.google.com/",
//	}
//
// Bitly API docs: https://dev.bitly.com/links.html#v3_link_lookup
func (client *Links) Lookup(urls ...string) (links []Link, err error) {
	return client.req("/link/lookup", url.Values{"url": urls}, "link_lookup")
}

// Shorten returns a short url from a given long url.
// Returns:
//
//	bitly.Link{
//		LongURL:    "http://google.com/",
//		GlobalHash: "900913",
//		Hash:       "ze6poY",
//		NewHash:    0,
//		URL:        "http://bit.ly/ze6poY",
//	}
//
// Bitly API docs: http://dev.bitly.com/links.html#v3_shorten
func (client *Links) Shorten(longURL string) (link Link, err error) {
	req, err := client.get("/shorten", url.Values{
		"longUrl": []string{longURL},
	})
	if err != nil {
		return
	}

	err = json.Unmarshal(req.Data, &link)

	return
}
