package regex

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var githubRegex = regexp.MustCompile("^(?P<owner>.*)/(?P<repo>.*)@(?P<version>.*)$")

func TestMatch(t *testing.T) {
	tests := []struct {
		re    *regexp.Regexp
		input string
		want  []map[string]string
	}{
		{
			re:    githubRegex,
			input: "foo/bar@baz",
			want: []map[string]string{
				{
					"":        "foo/bar@baz",
					"owner":   "foo",
					"repo":    "bar",
					"version": "baz",
				},
			},
		},
		{
			re:    githubRegex,
			input: "foo",
			want:  nil,
		},
		{
			re:    regexp.MustCompile(`^(?P<file>.*)\.(?P<ext>.*)$`),
			input: "foo.bar",
			want: []map[string]string{
				{
					"":     "foo.bar",
					"file": "foo",
					"ext":  "bar",
				},
			},
		},
	}

	for _, tt := range tests {
		result := Match(tt.re, tt.input)
		assert.Equal(t, tt.want, result)
	}
}
