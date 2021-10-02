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

A micro-helper-library for working with Go struct tags.

## Installation

Install with `go get`:

```sh
go get github.com/a-poor/tags
```

## Example

Here's a quick example of working with the `tags` library.

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
data, _ := json.MarshalIndent(fields, "", "  ")
fmt.Println(string(data))
// Output: {
//   "Email": [
//     "user_email",
//     "omitempty"
//   ],
//   "ID": [
//     "user_id"
//   ],
//   "ImEmpty": [
//     ""
//   ],
//   "Name": [
//     "",
//     "omitempty"
//   ]
// }
```

## Usage

The `tags` library is _very_ small. At least for now.

There's only one struct, `TagParser`, which has one field, `TagName`, and one method, `Parse`.

Say, for example, we have a struct that looks like this:

```go
type User struct {
    ID       int     `myTag:"user_id"`
    Name     string  `myTag:"name" otherTag"abc123"`
    Balance  float32 `myTag:"balance,omitempty"`
    IsActive bool    `myTag:",hello"`
}
```

If we wanted to get the struct tag values for `myTag`, we could create a new `TagParser` like this:

```go
tp := tags.TagParser{TagName: "myTag"}
```

and then parse the struct's tags like this:

```go
u := User{} // Create a blank user
ut := tp.Parse(u)
```

`ut` is of the type `map[string][]string`, where each of the map's keys are fields of the struct (with tags), and the map's values are arrays of tag values corresponding to the chosen tag, split on commas.

In our example, we would have the following result (formatted as JSON):

```json
{
  "Balance": [
    "balance",
    "omitempty"
  ],
  "ID": [
    "user_id"
  ],
  "IsActive": [
    "",
    "hello"
  ],
  "Name": [
    "name"
  ]
}
```

## To Do

- Should untagged fields appear in the returned result?
- Add more error checks 
  - ie Catch panics caused by `tags` and return them rather than letting the panic propagate
  - Check that the passed value is a struct (not a basic type)
- Be able to pass a pointer to a struct (without panicing)
- Fill a struct with struct tag values? 
  - ie struct would have fields `name`, `omitempty`, etc. and would be filled by position or value (like flags).

## License

[MIT](./LICENSE)

## Contributing

Go ahead and create an issue or submit a pull request! I'd love to hear from you.

