package find

import (
	"fmt"
	"io/fs"
	"strconv"
	"time"

	"github.com/djherbis/times"
)

// Newer returns a Matcher which returns true if the current file is newer than the supplied
// filename. "newer" can be in terms of creation, modification, change or access time. This
// maps the `find . -newerXY` except that we don't support the 't' parameter here.
func Newer(finder *Finder, timeType FileTimeType, compare string, compareTimeType FileTimeType) Matcher {

	if finder.CacheCmpFile {
		finder.cmpFileTime, _ = times.Stat(compare)
	}

	return func(path string, info fs.DirEntry) (bool, error) {

		var compareTimeSpec times.Timespec
		var pathTime, compareTime time.Time
		var err error

		compareTimeSpec = finder.cmpFileTime

		if cmpTimes == nil {
			cmpStat, err := finder.rootFS.Stat(compare)
			if err = finder.CallInternalErrorHandler(err); err != nil {
				return false, err
			}
			compareTimeSpec = times.Get(cmpStat)
			if finder.CacheCmpFile {
				finder.cmpFileTime = compareTimeSpec
			}
		}

		fInfo, err := info.Info()
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, err
		}

		pathTimeSpec := times.Get(fInfo)

		switch timeType {
		case Created:
			if pathTimeSpec.HasBirthTime() {
				pathTime = pathTimeSpec.BirthTime()
			} else {
				err := finder.CallInternalErrorHandler(&FinderError{Path: path, Info: "filesystem does not support birth time", Entry: info})
				return false, err
			}
		case Modified:
			pathTime = pathTimeSpec.ModTime()
		case Accessed:
			return pathTimes.AccessTime().After(cmpTimes.AccessTime()), nil
		case Changed:
			if cmpTimes.HasChangeTime() && pathTimes.HasChangeTime() {
				return pathTimes.ChangeTime().After(cmpTimes.ChangeTime()), nil
			} else {
				err := finder.CallInternalErrorHandler(&FinderError{Matcher: "Newer", Path: path, Info: "filesystem does not support change time", Entry: info})
				return false, err
			}
		default:
			err := finder.CallInternalErrorHandler(&FinderError{Matcher: "Newer", Path: path, Info: fmt.Sprintf("impossible time type: %s", strconv.QuoteRune(rune(timeType)))})
			return false, err
		}
	}
}

// Newer appends a Matcher which returns true if the current file is newer than the supplied file.
// "newer" can be in terms of creation, modification, change or access time. This
// maps the `find . -newerXY` except that we don't support the 't' parameter here.
func (finder *Finder) Newer(timeType FileTimeType, compare string) *Finder {
	return finder.appendMatcher(Newer(finder, timeType, compare))
}
