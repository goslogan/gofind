package find

import (
	"fmt"
	"io/fs"
	"strconv"
	"strings"

	"github.com/djherbis/times"
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

// Newer returns a Matcher which returns true if the current file is newer than the supplied
// filename. "newer" can be in terms of creation, modification, change or access time. This
// maps the `find . -newerXY` except that we don't support the 't' parameter here.
func Newer(finder *Finder, timeType FileTimeType, compare string) Matcher {

	if finder.CacheCmpFile {
		finder.cmpFileTime, _ = times.Stat(compare)
	}

	return func(path string, info fs.DirEntry) (bool, error) {

		var cmpTimes times.Timespec
		var err error

		cmpTimes = finder.cmpFileTime

		if cmpTimes == nil {
			cmpTimes, err = times.Stat(compare)
			if err = finder.CallInternalErrorHandler(err); err != nil {
				return false, err
			}
			if finder.CacheCmpFile {
				finder.cmpFileTime = cmpTimes
			}
		}

		fInfo, err := info.Info()
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, err
		}

		pathTimes := times.Get(fInfo)

		switch timeType {
		case Created:
			if cmpTimes.HasBirthTime() && pathTimes.HasBirthTime() {
				return pathTimes.BirthTime().After(cmpTimes.BirthTime()), nil
			} else {
				return false, &FinderError{Path: path, Info: "filesystem does not support birth time", Entry: info}
			}
		case Modified:
			return pathTimes.ModTime().After(cmpTimes.ModTime()), nil
		case Accessed:
			return pathTimes.AccessTime().After(cmpTimes.AccessTime()), nil
		case Changed:
			if cmpTimes.HasChangeTime() && pathTimes.HasChangeTime() {
				return pathTimes.ChangeTime().After(cmpTimes.ChangeTime()), nil
			} else {
				return false, &FinderError{Path: path, Info: "filesystem does not support change time", Entry: info}
			}
		default:
			return false, &FinderError{Path: path, Info: fmt.Sprintf("impossible time type: %s", strconv.QuoteRune(rune(timeType)))}
		}
	}
}
