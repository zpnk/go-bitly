package bitly

import (
	"encoding/json"
	"net/url"
)

// Link handles communication with the link related Bitly API endpoints.
type Link struct {
	*Client
}

// ExpandedLink represents the results of the Link.Expand method.
type ExpandedLink struct {
	GlobalHash string `json:"global_hash"`
	LongURL    string `json:"long_url"`
	ShortURL   string `json:"short_url"`
	UserHash   string `json:"user_hash"`
}

// LinkLookup represents the results of the Link.Lookup method.
type LinkLookup struct {
	URL           string `json:"url"`
	AggregateLink string `json:"aggregate_link"`
}

// Expand returns the long urls for a given set short urls.
//
// Bitly API docs: http://dev.bitly.com/links.html#v3_expand
func (client *Link) Expand(urls ...string) (links []ExpandedLink, err error) {
	req, err := client.Get("/expand", url.Values{"shortUrl": urls})
	if err != nil {
		return
	}

	res := map[string][]ExpandedLink{}
	err = json.Unmarshal(req.Data, &res)
	if err != nil {
		return
	}

	return res["expand"], err
}

// Lookup queries for bitlink(s) mapping to the given url(s).
//
// Bitly API docs: https://dev.bitly.com/links.html#v3_link_lookup
func (client *Link) Lookup(urls ...string) (links []LinkLookup, err error) {
	req, err := client.Get("/link/lookup", url.Values{"url": urls})
	if err != nil {
		return
	}

	res := map[string][]LinkLookup{}
	err = json.Unmarshal(req.Data, &res)
	if err != nil {
		return
	}

	return res["link_lookup"], err
}
