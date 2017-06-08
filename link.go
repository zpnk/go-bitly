package bitly

import (
	"encoding/json"
	"net/url"
)

// Link handles communication with the link related Bitly API endpoints.
type Link struct {
	*Client
}

// LinkLookup represents the results of the Link.Lookup method.
type LinkLookup struct {
	URL           string `json:"url"`
	AggregateLink string `json:"aggregate_link"`
}

// Lookup queries for bitlink(s) mapping to the given url(s).
//
// Bitly API docs: https://dev.bitly.com/links.html#v3_link_lookup
func (client *Link) Lookup(urls ...string) (linkLookup []LinkLookup, err error) {
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
