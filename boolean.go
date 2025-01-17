package find

import "io/fs"

// Or generates a Matcher which returns true if any of the Matchers provided returns true.
func Or(finder *Finder, matchers ...Matcher) Matcher {

	return func(path string, info fs.FileInfo) (bool, error) {
		for _, matcher := range matchers {
			// internal error handler has already been called and err returned if err is not nil
			matched, err := matcher(path, info)
			if err != nil {
				return false, err
			} else if matched {
				return true, nil
			}
		}
		return false, nil
	}
}

// Not generates a Matcher which returns true if the provided Matcher returns false.
func Not(finder *Finder, matcher Matcher) Matcher {
	return func(path string, info fs.FileInfo) (bool, error) {
		matched, err := matcher(path, info)
		if err != nil {
			return false, err
		} else {
			return !matched, nil
		}
	}
}

// And generates a Matcher which returns true if all of the Matchers provided return true.
// Default behaviour is to AND results but this allows for complex nested matching behaviour.
func And(finder *Finder, matchers ...Matcher) Matcher {
	return func(path string, info fs.FileInfo) (bool, error) {
		for _, matcher := range matchers {
			// internal error handler has already been called and err returned if err is not nil
			matched, err := matcher(path, info)
			if err != nil {
				return false, err
			} else if !matched {
				return false, nil
			}
		}
		return true, nil
	}
}

// Or appends a Matcher which returns true if any of the supplied Matchers returns true.
func (finder *Finder) Or(matchers ...Matcher) *Finder {
	return finder.appendMatcher(Or(finder, matchers...))
}

// Not appends a Matcher which returns true if the provided Matcher retursn false.
func (finder *Finder) Not(matcher Matcher) *Finder {
	return finder.appendMatcher(Not(finder, matcher))
}

// And appends a Matcher which returns true if all of the supplied Matchers return true.
func (finder *Finder) And(matchers ...Matcher) *Finder {
	return finder.appendMatcher(And(finder, matchers...))
}
