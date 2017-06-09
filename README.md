# go-bitly

[![Build Status](https://travis-ci.org/zpnk/go-bitly.svg?branch=master)](https://travis-ci.org/zpnk/go-bitly)
[![Coverage Status](https://coveralls.io/repos/github/zpnk/go-bitly/badge.svg?branch=master)](https://coveralls.io/github/zpnk/go-bitly?branch=master)
[![GoDoc](https://godoc.org/github.com/zpnk/go-bitly/bitly?status.svg)](https://godoc.org/github.com/zpnk/go-bitly)

Simple Bitly API v3 client for Go.

***Work in progress***

## Usage

```
$ go get github.com/zpnk/go-bitly
```

```go
import "github.com/zpnk/go-bitly"

b := bitly.New("<token>")

shortURL, err := b.Links.Shorten("https://golang.org/")

// bitly.Link{URL:"https://bit.ly/2scFxid", ... }
```

## Roadmap

The goal of this library is to provide a simple, well tested means of working
with the complete Bitly API. Coverage and consistency with the [official docs](https://dev.bitly.com/api.html)
is a top priority. To that end, contributions are welcome and encouraged!
Calling and testing patterns have been established so adding support for endpoints
should be relatively straightforward.

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.
