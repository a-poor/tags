# tags

[![Go Reference](https://pkg.go.dev/badge/github.com/a-poor/tags.svg)](https://pkg.go.dev/github.com/a-poor/tags)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/a-poor/tags?style=flat-square)](https://pkg.go.dev/github.com/a-poor/tags)
[![Go Test](https://github.com/a-poor/tags/actions/workflows/go.yml/badge.svg)](https://github.com/a-poor/tags/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/a-poor/tags)](https://goreportcard.com/report/github.com/a-poor/tags)
[![GitHub](https://img.shields.io/github/license/a-poor/tags?style=flat-square)](https://github.com/a-poor/tags/blob/main/LICENSE)
![GitHub last commit](https://img.shields.io/github/last-commit/a-poor/tags?style=flat-square)
[![Sourcegraph](https://sourcegraph.com/github.com/a-poor/tags/-/badge.svg)](https://sourcegraph.com/github.com/a-poor/tags?badge)
[![CodeFactor](https://www.codefactor.io/repository/github/a-poor/tags/badge/main)](https://www.codefactor.io/repository/github/a-poor/tags/overview/main)


_created by Austin Poor_

A micro helper library for working with Go struct tags.

## Installation

Install with `go get`:

```sh
go get github.com/a-poor/tags
```

## Example

```go
// Define a struct that we'll be getting the tags from
user := struct {
    ID      int    `app:"user_id"`
    Name    string `app:",omitempty"`
    Email   string `app:"user_email,omitempty"`
    NotMe   bool
    ImEmpty bool `app:""`
}{}

// Parse the struct's tags
fields := tags.ParseStructTags("app", user)

// Print out the results as JSON
data, _ := json.Marshal(fields)
fmt.Println(string(data))
// Output: {"Email":["user_email","omitempty"],"ID":["user_id"],"ImEmpty":[""],"Name":["","omitempty"]}
```

