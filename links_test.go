package bitly

import (
	"reflect"
	"testing"
)

func TestLinks_Expand(t *testing.T) {
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

	links, err := client.Links.Expand("http://bit.ly/1RmnUT")
	if err != nil {
		t.Fatalf("Links.Expand returned error: %v", err)
	}

	want := Link{
		GlobalHash: "1RmnUT",
		LongURL:    "http://google.com",
		ShortURL:   "http://bit.ly/1RmnUT",
		UserHash:   "1RmnUT",
	}
	if !reflect.DeepEqual(links[0], want) {
		t.Errorf("Links.Expand returned %+v, want %+v", links[0], want)
	}
}

func TestLinks_Info(t *testing.T) {
	setup()
	defer teardown()

	newEndpoint("/info", jsonRes(
		`{
			"info": [
	      {
	        "created_at": 1212926400,
	        "created_by": null,
	        "global_hash": "1RmnUT",
	        "short_url": "http://bit.ly/1RmnUT",
	        "title": "Google",
	        "user_hash": "1RmnUT"
	      }
	    ]
		}`,
		200,
		"OK",
	))

	links, err := client.Links.Info("http://bit.ly/1RmnUT")
	if err != nil {
		t.Fatalf("Links.Info returned error: %v", err)
	}

	want := Link{
		CreatedAt:  1212926400,
		GlobalHash: "1RmnUT",
		ShortURL:   "http://bit.ly/1RmnUT",
		Title:      "Google",
		UserHash:   "1RmnUT",
	}
	if !reflect.DeepEqual(links[0], want) {
		t.Errorf("Links.Info returned %+v, want %+v", links[0], want)
	}
}

func TestLinks_Lookup_single(t *testing.T) {
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

	links, err := client.Links.Lookup("http://www.google.com/")
	if err != nil {
		t.Fatalf("Links.Lookup returned error: %v", err)
	}

	want := Link{
		AggregateLink: "http://bit.ly/2V6CFi",
		URL:           "http://www.google.com/",
	}
	if !reflect.DeepEqual(links[0], want) {
		t.Errorf("Links.Lookup returned %+v, want %+v", links[0], want)
	}
}

func TestLinks_Lookup_multiple(t *testing.T) {
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

	links, err := client.Links.Lookup("http://www.google.com/", "http://www.facebook.com/")
	if err != nil {
		t.Fatalf("Links.Lookup returned error: %v", err)
	}

	want := []Link{
		Link{
			AggregateLink: "http://bit.ly/2V6CFi",
			URL:           "http://www.google.com/",
		},
		Link{
			AggregateLink: "http://bit.ly/4VGeu",
			URL:           "http://www.facebook.com/",
		},
	}
	if !reflect.DeepEqual(links, want) {
		t.Errorf("Links.Lookup returned %#v, want %#v", links, want)
	}
}
