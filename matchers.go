package find

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/goslogan/gofind/internal"
)

// Name generates a Matcher which returns true if the name of the file (or directory)
// matches the glob pattern provided.
func Name(finder *Finder, glob string) Matcher {

	return func(path string, info fs.DirEntry) (bool, error) {
		matched, err := filepath.Match(glob, info.Name())
		if err != nil && !finder.InternalErrorHandler(err) {
			return false, &FinderError{Err: err, Path: path, Info: glob, Entry: info}
		}
		return matched, nil
	}
}

// Dir returns a Matcher which returns true if the path is a directory, false if not
// `find . -type d`
func Dir(finder *Finder) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		return info.IsDir(), nil
	}
}

// File returns a Matcher which returns true if the path is regular file, false if not
// `find . -type f`
func File(finder *Finder) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		return info.Type().IsRegular(), nil
	}
}

// Depth returns a Matcher which returns true if the current path (final component) is
// `depth` deep compared to the starting point (depth zero)
// `find . -depth n`
func Depth(finder *Finder, depth int) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		return finder.pathDepth(path) == depth, nil
	}
}

// Mindepth returns a Matcher which returns true if the current path is at least
// depth deep compared to the starting point.
func MinDepth(finder *Finder, depth int) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		return finder.pathDepth(path) >= depth, nil
	}
}

// MaxDepth returns a Matcher which returns true if the current path is at most
// depth deep compared to the starting point.
func MaxDepth(finder *Finder, depth int) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		return finder.pathDepth(path) <= depth, nil
	}
}

// Owner returns a Matcher which returns true if the file is owned by the named user. If that
// fails we try to treat the name as a number and see if matches the owner uid instead. The test
// is case insensitive.
func Owner(finder *Finder, name string) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		user, err := internal.FileOwnerUser(path, info)
		if err != nil && !finder.InternalErrorHandler(err) {
			return false, err
		}
		if strings.EqualFold(name, user.Name) {
			return true, nil
		} else {
			return name == user.Uid, nil
		}
	}
}

// Group returns a Matcher which returns true if the file is owned by the named group. If that
// fails we try to treat the name as a number and see if it matches the owner gid instead. The test
// is case insensitive.
func Group(finder *Finder, name string) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		group, err := internal.FileOwnerGroup(path, info)
		if err != nil && !finder.InternalErrorHandler(err) {
			return false, err
		}
		if strings.EqualFold(name, group.Name) {
			return true, nil
		} else {
			return name == group.Gid, nil
		}
	}
}

// Or generates a Matcher which returns true if any of the Matchers provided returns true.
func Or(finder *Finder, matchers ...Matcher) Matcher {

	return func(path string, info fs.DirEntry) (bool, error) {
		for _, matcher := range matchers {
			// internal error handler has already been called and err returned if err is not nil
			cont, err := matcher(path, info)
			if err != nil {
				return false, err
			} else if !cont {
				return false, nil
			}
		}
		return true, nil
	}
}
