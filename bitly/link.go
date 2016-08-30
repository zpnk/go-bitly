package bitly

import "github.com/mitchellh/mapstructure"

// Link handles communication with the link related Bitly API endpoints.
type Link struct {
	*Client
}

// linkLookupReq specifies the optional parameters to the /link/lookup endpoint.
type linkLookupReq struct {
	URLs []string `url:"url"`
}

// LinkLookupRes specifies the response of the Link.Lookup method.
type LinkLookupRes struct {
	AggregateLink string `mapstructure:"aggregate_link"`
	URL           string `mapstructure:"url"`
}

// Lookup queries for bitlink(s) mapping to the given url(s).
//
// Bitly API docs: https://dev.bitly.com/links.html#v3_link_lookup
func (c *Link) Lookup(urls ...string) (res []LinkLookupRes, err error) {
	params := linkLookupReq{URLs: urls}
	r, err := c.Get("/link/lookup", params)
	if err != nil {
		return
	}

	err = mapstructure.Decode(r.Data["link_lookup"], &res)
	if err != nil {
		return
	}

	return
}
