/*
Package bitly provides a simple client for using the Bitly v3 API.

Usage:

	import "github.com/zpnk/go-bitly"

	client := bitly.NewClient("<token>")

	links, err := bitly.Link.Lookup("http://golang.org/", "http://google.com/")

*/
package bitly
