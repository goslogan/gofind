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

		if compareTimeSpec == nil {
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
		pathTime, err = finder.getTime(path, info, pathTimeSpec, timeType)
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, err
		}

		compareTime, err = finder.getTime(compare, info, compareTimeSpec, compareTimeType)
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, err
		}

		return pathTime.After(compareTime), nil
	}
}

func (finder *Finder) getTime(path string, info fs.DirEntry, source times.Timespec, timeType FileTimeType) (time.Time, error) {
	switch timeType {
	case Created:
		if source.HasBirthTime() {
			return source.BirthTime(), nil
		} else {
			err := finder.CallInternalErrorHandler(&FinderError{Path: path, Info: "filesystem does not support birth time", Entry: info})
			return time.Time{}, err
		}
	case Modified:
		return source.ModTime(), nil
	case Accessed:
		return source.AccessTime(), nil
	case Changed:
		if source.HasChangeTime() {
			return source.ChangeTime(), nil
		} else {
			err := finder.CallInternalErrorHandler(&FinderError{Matcher: "Newer", Path: path, Info: "filesystem does not support change time", Entry: info})
			return time.Time{}, err
		}
	default:
		err := finder.CallInternalErrorHandler(&FinderError{Matcher: "Newer", Path: path, Info: fmt.Sprintf("impossible time type: %s", strconv.QuoteRune(rune(timeType)))})
		return time.Time{}, err
	}
}

// Newer appends a Matcher which returns true if the current file is newer than the supplied file.
// "newer" can be in terms of creation, modification, change or access time. This
// maps the `find . -newerXY` except that we don't support the 't' parameter here.
func (finder *Finder) Newer(timeType FileTimeType, compare string, compareTimeType FileTimeType) *Finder {
	return finder.appendMatcher(Newer(finder, timeType, compare, compareTimeType))
}

// Amin returns a Matcher which returns true if the current difference between the time find was started
// and the access time of the file is the number of minutes specified. The second argument defines whether
// the exact time of less than or greater than is required.
func Amin(finder *Finder, minutes time.Duration, compare TimeCompareType) Matcher {
	return xmin(finder, minutes, compare, Accessed)
}

// Amin appends a Matcher which returns true if the current difference between the time find was started
// and the access time of the file is the number of minutes specified. The second argument defines whether
// the exact time of less than or greater than is required.
func (finder *Finder) Amin(minutes time.Duration, compare TimeCompareType) *Finder {
	return finder.appendMatcher(Amin(finder, minutes, compare))
}

// Amin returns a Matcher which returns true if the current difference between the time find was started
// and the changed time of the file is the number of minutes specified. The second argument defines whether
// the exact time of less than or greater than is required.
func Cmin(finder *Finder, minutes time.Duration, compare TimeCompareType) Matcher {
	return xmin(finder, minutes, compare, Changed)
}

// Amin appends a Matcher which returns true if the current difference between the time find was started
// and the changed time of the file is the number of minutes specified. The second argument defines whether
// the exact time of less than or greater than is required.
func (finder *Finder) Cmin(minutes time.Duration, compare TimeCompareType) *Finder {
	return finder.appendMatcher(Cmin(finder, minutes, compare))
}

// Mmin returns a Matcher which returns true if the current difference between the time find was started
// and the modified time of the file is the number of minutes specified. The second argument defines whether
// the exact time of less than or greater than is required.
func Mmin(finder *Finder, minutes time.Duration, compare TimeCompareType) Matcher {
	return xmin(finder, minutes, compare, Modified)
}

// Mmin appends a Matcher which returns true if the current difference between the time find was started
// and the modified time of the file is the number of minutes specified. The second argument defines whether
// the exact time of less than or greater than is required.
func (finder *Finder) Mmin(minutes time.Duration, compare TimeCompareType) *Finder {
	return finder.appendMatcher(Mmin(finder, minutes, compare))
}

// Bmin returns a Matcher which returns true if the current difference between the time find was started
// and the creation (birth) time of the file is the number of minutes specified. The second argument defines whether
// the exact time of less than or greater than is required.
func Bmin(finder *Finder, minutes time.Duration, compare TimeCompareType) Matcher {
	return xmin(finder, minutes, compare, Created)
}

// Bmin appends a Matcher which returns true if the current difference between the time find was started
// and the creation time of the file is the number of minutes specified. The second argument defines whether
// the exact time of less than or greater than is required.
func (finder *Finder) Bmin(minutes time.Duration, compare TimeCompareType) *Finder {
	return finder.appendMatcher(Bmin(finder, minutes, compare))
}

// xmin is the underlying impelementation of the Amin, Bmin, Mmin and Cmin matchers
func xmin(finder *Finder, minutes time.Duration, compare TimeCompareType, timeType FileTimeType) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		fInfo, err := info.Info()
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, err
		}
		timeSpec := times.Get(fInfo)
		fileTime, err := finder.getTime(path, info, timeSpec, timeType)
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, err
		}

		difference := finder.started.Sub(fileTime).Abs().Round(time.Minute)

		switch compare {
		case Equal:
			return minutes.Round(time.Minute) == difference, nil
		case LessThan:
			return difference < minutes.Round(time.Minute), nil
		case GreaterThan:
			return difference > minutes.Round(time.Minute), nil
		default:
			err := finder.CallInternalErrorHandler(&FinderError{Matcher: "Xmin", Path: path, Info: fmt.Sprintf("impossible time compare type: %s", strconv.QuoteRune(rune(compare)))})
			return false, err
		}
	}
}
