package tags

import (
	"reflect"
	"strings"
)

// ParseStructTags parses a struct's tags and returns
// a map of the struct's fields to an array of the tag values.
//
// This is a convenience function that creates a new `TagParse`
// instance and calles its `Parse` method.
func ParseStructTags(tagName string, d interface{}) map[string][]string {
	return TagParser{TagName: tagName}.Parse(d)
}

// TagParser is a struct used for parsing struct tags
// with the key `TagName`.
type TagParser struct {
	TagName string // They key to look for in the struct tags
}

// Parse parses the struct tags of of the given struct
// and returns a map from the struct's fields to an array
// of the tag values.
//
// Note: `d` is expected to be passed by value, not by reference.
func (p TagParser) Parse(d interface{}) map[string][]string {
	// Map to store the struct tags
	var st map[string][]string

	// Get the type of the struct
	dt := reflect.TypeOf(d)

	// Get the number of (visible) fields in the struct
	n := dt.NumField()
	if n > 0 {
		st = make(map[string][]string)
	}

	// Iterate through the fields
	for i := 0; i < n; i++ {
		field := dt.Field(i)

		val, ok := field.Tag.Lookup(p.TagName)
		if ok {
			// Split the tag into its parts and store
			st[field.Name] = strings.Split(val, ",")
		}

	}

	// Return the results map
	return st
}
