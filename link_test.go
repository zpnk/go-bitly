package bitly

import (
	"reflect"
	"testing"
)

func TestLink_Expand(t *testing.T) {
	setup()
	defer teardown()

	newEndpoint("/expand", jsonRes(
		`{
		"expand": [
				{
					"global_hash": "1RmnUT",
					"long_url": "http://google.com",
					"short_url": "http://bit.ly/1RmnUT",
					"user_hash": "1RmnUT"
				}
			]
		}`,
		200,
		"OK",
	))

	links, err := client.Link.Expand("http://bit.ly/1RmnUT")
	if err != nil {
		t.Fatalf("Link.Expand returned error: %v", err)
	}

	want := ExpandedLink{
		GlobalHash: "1RmnUT",
		LongURL:    "http://google.com",
		ShortURL:   "http://bit.ly/1RmnUT",
		UserHash:   "1RmnUT",
	}
	if !reflect.DeepEqual(links[0], want) {
		t.Errorf("Link.Expand returned %+v, want %+v", links[0], want)
	}
}

func TestLink_Lookup_single(t *testing.T) {
	setup()
	defer teardown()

	newEndpoint("/link/lookup", jsonRes(
		`{
			"link_lookup": [
				{
					"aggregate_link": "http://bit.ly/2V6CFi",
					"url": "http://www.google.com/"
				}
			]
		}`,
		200,
		"OK",
	))

	links, err := client.Link.Lookup("http://www.google.com/")
	if err != nil {
		t.Fatalf("Link.Lookup returned error: %v", err)
	}

	want := LinkLookup{AggregateLink: "http://bit.ly/2V6CFi", URL: "http://www.google.com/"}
	if !reflect.DeepEqual(links[0], want) {
		t.Errorf("Link.Lookup returned %+v, want %+v", links[0], want)
	}
}

func TestLink_Lookup_multiple(t *testing.T) {
	setup()
	defer teardown()

	newEndpoint("/link/lookup", jsonRes(
		`{
			"link_lookup": [
				{
					"aggregate_link": "http://bit.ly/2V6CFi",
					"url": "http://www.google.com/"
				},
				{
					"aggregate_link": "http://bit.ly/4VGeu",
					"url": "http://www.facebook.com/"
				}
			]
		}`,
		200,
		"OK",
	))

	links, err := client.Link.Lookup("http://www.google.com/", "http://www.facebook.com/")
	if err != nil {
		t.Fatalf("Link.Lookup returned error: %v", err)
	}

	want := []LinkLookup{
		LinkLookup{AggregateLink: "http://bit.ly/2V6CFi", URL: "http://www.google.com/"},
		LinkLookup{AggregateLink: "http://bit.ly/4VGeu", URL: "http://www.facebook.com/"},
	}
	if !reflect.DeepEqual(links, want) {
		t.Errorf("Link.Lookup returned %#v, want %#v", links, want)
	}
}
