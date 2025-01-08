package find

import "io/fs"

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

// Type returns a Matcher which returns true if the path is of the type provided.
// `find . -type X` where X is the type.
func Type(finder *Finder, t fs.FileMode) Matcher {
	return func(path string, info fs.DirEntry) (bool, error) {
		return info.Type() == t, nil
	}
}

// Dir appends a matcher which selects paths which are directories only
// `find . -type d`
func (finder *Finder) Dir() *Finder {
	return finder.appendMatcher(Dir(finder))
}

// File appends a matcher which selects paths which are regular files only
// `find . -type f`
func (finder *Finder) File() *Finder {
	return finder.appendMatcher(File(finder))
}

// Type appends a matcher which selects paths which match the type provided.
func (finder *Finder) Type(t fs.FileMode) *Finder {
	return finder.appendMatcher(Type(finder, t))
}
