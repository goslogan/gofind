package find

import (
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
)

// Name returns a Matcher which returns true if the name of the file (or directory)
// matches the glob pattern provided.
func Name(finder *Finder, glob string) Matcher {

	return func(path string, info fs.FileInfo) (bool, error) {
		matched, err := filepath.Match(glob, info.Name())
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, &FinderError{Matcher: "Name", Err: err, Path: path, Info: glob, Entry: info}
		}
		return matched, nil
	}
}

// IName returns a Matcher which returns true if the name of the file (or directory) case-insensitively
// matches the glob (by mapping both to lower case - this may not be 100% reliable for all glob patterns)
func IName(finder *Finder, glob string) Matcher {

	lglob := strings.ToLower(glob)
	return func(path string, info fs.FileInfo) (bool, error) {
		matched, err := filepath.Match(lglob, strings.ToLower(info.Name()))
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, &FinderError{Matcher: "IName", Err: err, Path: path, Info: glob, Entry: info}
		}
		return matched, nil
	}
}

// Path returns a Matcher which returns true if the entire path matches the glob pattern provided.
func Path(finder *Finder, glob string) Matcher {

	return func(path string, info fs.FileInfo) (bool, error) {
		matched, err := filepath.Match(glob, path)
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, &FinderError{Matcher: "Name", Err: err, Path: path, Info: glob, Entry: info}
		}
		return matched, nil
	}
}

// IPath returns  a Matcher which returns true if the entire path case-insensitively
// matches the glob (by mapping both to lower case - this may not be 100% reliable for all glob patterns)
func IPath(finder *Finder, glob string) Matcher {

	lglob := strings.ToLower(glob)
	return func(path string, info fs.FileInfo) (bool, error) {
		matched, err := filepath.Match(lglob, strings.ToLower(path))
		if err = finder.CallInternalErrorHandler(err); err != nil {
			return false, &FinderError{Matcher: "IName", Err: err, Path: path, Info: glob, Entry: info}
		}
		return matched, nil
	}
}

// Regex returns a Matcher which returns true if the full path matches the supplied regular expression
// Note that there is no equivalent to the -iregex flag in find - just make the regular expression case insensitive.
func Regex(finder *Finder, regex *regexp.Regexp) Matcher {
	return func(path string, info fs.FileInfo) (bool, error) {
		return regex.MatchString(path), nil
	}
}

// Name appends a matcher which selects paths based on a glob of their name (where name is the final component in the path)
// `find . -name x`
func (finder *Finder) Name(glob string) *Finder {
	return finder.appendMatcher(Name(finder, glob))
}

// IName appends a matcher which selects paths based on a case-insensitive glob of their name (where name is the final component in the path)
func (finder *Finder) IName(glob string) *Finder {
	return finder.appendMatcher(IName(finder, glob))
}

// Name appends a matcher which selects paths based on a glob of their entire path
// `find . -name x`
func (finder *Finder) Path(glob string) *Finder {
	return finder.appendMatcher(Path(finder, glob))
}

// IName appends a matcher which selects paths based on a case-insensitive glob of their entire path
func (finder *Finder) IPath(glob string) *Finder {
	return finder.appendMatcher(IPath(finder, glob))
}

// Regex appends a matcher which selects paths based on a regular expression of their entire path
func (finder *Finder) Regex(regex *regexp.Regexp) *Finder {
	return finder.appendMatcher(Regex(finder, regex))
}
