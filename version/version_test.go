package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersionFilter_Match(t *testing.T) {
	tests := []struct {
		name          string
		version       string
		versionFilter string
		want          bool
	}{
		{"exact match", "1.2.3", "1.2.3", true},
		{"wildcard match patch", "1.2.3", "1.2.x", true},
		{"wildcard match minor", "1.2.3", "1.x", true},
		{"wildcard no match", "1.2.3", "1.3.x", false},
		{"exact no match", "1.2.3", "3.2.1", false},
		{"range match", "1.2.3", ">1.0.0", true},
		{"range match 2", "1.2.3", ">1.0.0 <2.0.0", true},
		{"range no match", "1.2.3", "<1.0.0", false},
		{"range no match 2", "1.2.3", ">2.0.0 <3.0.0", false},
		{"match all", "1.2.3", ">=0.0.0", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Match(tt.version, tt.versionFilter)
			require.NoError(t, err)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestVersionFilterMap(t *testing.T) {
	vfm := NewVersionFilterMap[string]()
	vfm.Add("1.x", "one")
	vfm.Add("2.x", "two")
	vfm.Add(">=4.0.0 <5.0.0", "four range")
	vfm.Add("3.2.1", "three")

	type testCase struct {
		name        string
		version     string
		want        string
		shouldMatch bool
	}
	tests := []testCase{
		{"v1.2.3", "1.2.3", "one", true},
		{"v2.3.4", "2.3.4", "two", true},
		{"v3.3.3", "3.3.3", "", false},
		{"v4.5.6", "4.5.6", "four range", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := vfm.Get(tt.version)
			assert.Equal(t, tt.shouldMatch, result.IsPresent())
			val, _ := result.Get()
			assert.Equal(t, tt.want, val)
		})
	}
}
