package version

import (
	"github.com/blang/semver/v4"
	"github.com/samber/mo"
)

// Match checks if a version is matched by a given version filter spec
// The versionFilter syntax is a range from:
// https://pkg.go.dev/github.com/blang/semver/v4#Range
func Match(version, versionFilter string) (bool, error) {
	v, err := semver.Make(version)
	if err != nil {
		return false, err
	}

	if versionFilter == "*" {
		versionFilter = ">=0.0.0 <=0.0.0"
	}

	r, err := semver.ParseRange(versionFilter)
	if err != nil {
		return false, err
	}

	result := r(v)
	return result, nil
}

func NewVersionFilterMap[T any]() *VersionFilterMap[T] {
	return &VersionFilterMap[T]{
		x: make(map[string]T),
	}
}

// VersionFilterMap is a map that holds items based on a SemVer version string,
// and returns an item which matches a version range spec.
// Supports range syntax from https://pkg.go.dev/github.com/blang/semver/v4#Range
type VersionFilterMap[T any] struct { //nolint:revive
	x map[string]T
}

func (vm VersionFilterMap[T]) Add(versionFilter string, obj T) {
	vm.x[versionFilter] = obj
}

func (vm VersionFilterMap[T]) Get(version string) mo.Option[T] {
	for versionFilter, obj := range vm.x {
		match, err := Match(version, versionFilter)
		if err != nil {
			return mo.None[T]()
		}
		if match {
			return mo.Some[T](obj)
		}
	}
	return mo.None[T]()
}
