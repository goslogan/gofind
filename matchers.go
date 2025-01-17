package find

import (
	"fmt"
	"io/fs"
	"strconv"
	"strings"
	"syscall"

	"github.com/goslogan/go-bytesize"
	"github.com/goslogan/gofind/internal"
)

// Depth returns a Matcher which returns true if the current path (final component) is
// `depth` deep compared to the starting point (depth zero)
// `find . -depth n`
func Depth(finder *Finder, depth int) Matcher {
	return func(path string, info fs.FileInfo) (bool, error) {
		return finder.pathDepth(path) == depth, nil
	}
}

// Mindepth returns a Matcher which returns true if the current path is at least
// depth deep compared to the starting point.
func MinDepth(finder *Finder, depth int) Matcher {
	return func(path string, info fs.FileInfo) (bool, error) {
		return finder.pathDepth(path) >= depth, nil
	}
}

// MaxDepth returns a Matcher which returns true if the current path is at most
// depth deep compared to the starting point.
func MaxDepth(finder *Finder, depth int) Matcher {
	return func(path string, info fs.FileInfo) (bool, error) {
		return finder.pathDepth(path) <= depth, nil
	}
}

// Owner returns a Matcher which returns true if the file is owned by the named user. If that
// fails we try to treat the name as a number and see if matches the owner uid instead. The test
// is case insensitive.
func Owner(finder *Finder, name string) Matcher {
	return func(path string, info fs.FileInfo) (bool, error) {
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
	return func(path string, info fs.FileInfo) (bool, error) {
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
	return func(path string, info fs.FileInfo) (bool, error) {
		return false, fs.SkipDir
	}
}

// Empty returns a Matcher which returns true if the current path (file or directory) is empty.
func Empty(finder *Finder) Matcher {

	return func(path string, info fs.FileInfo) (bool, error) {
		if info.IsDir() {
			dirContent, err := finder.rootFS.ReadDir(path)
			if finder.CallInternalErrorHandler(err); err != nil {
				return false, err
			} else {
				return len(dirContent) == 0, nil
			}
		} else if info.Mode().IsRegular() {
			return info.Size() == 0, nil
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

// Size returns a Matcher which returns true if the current file has the given size using the
// provided units. The file size is rounded up to the next unit before testing.
func Size(finder *Finder, size int64, units TimeSizeType) Matcher {

	scaler := map[TimeSizeType]bytesize.ByteSize{
		Bytes:     bytesize.B,
		Blocks:    bytesize.KB * 2,
		Kilobytes: bytesize.KB,
		Megabytes: bytesize.MB,
		Gigabytes: bytesize.GB,
		Terabytes: bytesize.TB,
		Petabytes: bytesize.PB,
	}[units]

	bs := bytesize.ByteSize(size) * scaler

	return func(path string, info fs.FileInfo) (bool, error) {
		if scaler != 0 {
			scaled := bytesize.ByteSize(info.Size()).Round(scaler)
			return scaled == bs, nil
		} else {
			return false, fmt.Errorf("gofind: unknown time size units - %s", strconv.QuoteRune(rune(units)))
		}
	}
}

// Size appends a Matcher which returns true if the current file has the given size using the
// provided units. The file size is rounded up to the next unit before testing.
func (finder *Finder) Size(size int64, units TimeSizeType) *Finder {
	return finder.appendMatcher(Size(finder, size, units))
}

// Sparse returns true if the file is a sparse file (that is, the file size in blocks indicates the
// file is smaller than the the Size would indicate).
func Sparse(finder *Finder) Matcher {
	return func(path string, info fs.FileInfo) (bool, error) {
		blocks := info.Sys().(*syscall.Stat_t).Blocks
		return blocks*512 < info.Size(), nil
	}
}

// Sparse appends a Matcher which returns true if the current file is a sparse file.
func (finder *Finder) Sparse() *Finder {
	return finder.appendMatcher(Sparse(finder))
}
