/*
Package bitly provides a simple client for using the Bitly v3 API.

Installation:

	go get github.com/zpnk/go-bitly

Usage:

	import "github.com/zpnk/go-bitly"

	b := bitly.New("<token>")

	shortURL, err := b.Links.Shorten("https://golang.org/")

	// bitly.Link{URL:"https://bit.ly/2scFxid", ... }
*/
package bitly
