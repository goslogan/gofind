package find

import (
	"io/fs"
	"strings"

	"github.com/goslogan/gofind/internal"
)

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
		if err = finder.CallInternalErrorHandler(err); err != nil {
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
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, err
		}
		if strings.EqualFold(name, group.Name) {
			return true, nil
		} else {
			return name == group.Gid, nil
		}
	}
}

// Prune returns a Matcher which simply stops the walk down the current path.
func Prune(finder *Finder) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		return false, fs.SkipDir
	}
}

// Empty returns a Matcher which returns true if the current path (file or directory) is empty.
func Empty(finder *Finder) Matcher {

	return func(path string, info fs.DirEntry) (bool, error) {
		if info.IsDir() {
			dirContent, err := finder.rootFS.ReadDir(path)
			if finder.CallInternalErrorHandler(err); err != nil {
				return false, err
			} else {
				return len(dirContent) == 0, nil
			}
		} else if info.Type().IsRegular() {
			fileInfo, err := info.Info()
			if finder.CallInternalErrorHandler(err); err != nil {
				return false, err
			} else {
				return fileInfo.Size() == 0, nil
			}
		} else {
			// no meaningful behaviour defined for other types
			return false, nil
		}
	}

}

// Empty appends a Matcher which returns true if the current path (file or drectory) is empty.
func (finder *Finder) Empty() *Finder {
	return finder.appendMatcher(Empty(finder))
}
