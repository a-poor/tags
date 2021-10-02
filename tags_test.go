package tags_test

import (
	"testing"

	"github.com/a-poor/tags"
)

// TODO: Fix this test
// func TestTagParseErr(t *testing.T) {
// 	tags.ParseStructTags("hello", 1)
// }

func TestTagParser(t *testing.T) {
	// Create a tag parser instance
	p := tags.TagParser{
		TagName: "test",
	}

	// Create a struct to parse
	u := struct {
		ID        int     `test:"user_id"`
		FirstName string  `test:"first_name" json:"not_me"`
		LastName  string  `test:"last_name,omitempty"`
		notHidden string  `test:"not_hidden"`
		ANumber   float32 `test:""`
		IsCool    bool
	}{
		1,
		"Tim",
		"Tomson",
		"secret",
		2.3,
		true,
	}

	// Parse the struct tags
	structTags := p.Parse(u)

	// Create a slice of test cases
	testCases := []struct {
		name   string
		exists bool
		expect []string
	}{
		{
			"ID",
			true,
			[]string{"user_id"},
		},
		{
			"FirstName",
			true,
			[]string{"first_name"},
		},
		{
			"LastName",
			true,
			[]string{"last_name", "omitempty"},
		},
		{
			"notHidden",
			true,
			[]string{"not_hidden"},
		},
		{
			"ANumber",
			true,
			[]string{""},
		},
		{
			"IsCool",
			false,
			[]string{},
		},
	}

	for _, tc := range testCases {
		r, ok := structTags[tc.name]

		// Check that tag existance was the same
		if ok != tc.exists {
			t.Errorf("tag %q exists = %t, not %t", tc.name, ok, tc.exists)
		}

		// Check that result length is the same
		if len(r) != len(tc.expect) {
			t.Errorf("tag %q array was %d, not %d", tc.name, len(r), len(tc.expect))

			// Skip the rest of the tests (which will cause a panic)
			continue
		}

		// Check that result is the same
		for i, v := range r {
			e := tc.expect[i]
			if v != e {
				t.Errorf("tag %q array[%d] was %q, not %q", tc.name, i, v, e)
			}
		}
	}

}
