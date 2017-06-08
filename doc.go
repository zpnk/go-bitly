/*
Package bitly provides a simple client for using the Bitly v3 API.

Installation:

	go get github.com/zpnk/go-bitly

Usage:

	import "github.com/zpnk/go-bitly"

	client := bitly.New("<token>")

	links, err := bitly.Link.Lookup("http://golang.org/", "http://google.com/")

*/
package bitly
