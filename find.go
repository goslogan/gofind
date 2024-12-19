// find provides builder functions to create a fs.WalkDirFunc that emulates `find`
package find

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Default capacity for matched file array
var DefaultCapacity int = 100

// Match function type - should return (true, nil) if matching should continue with the current path
// and false, nil if not. If an error occurs, it should return false, err.
// the internal error handler function (Finder.InternalErrorHandler) should be called by the matcher
type Matcher func(path string, info fs.DirEntry) (bool, error)

// ErrorHandler is the function called by a Finder when an error occurs whilst calling
// fs.WalkDir or when an internal error occurs. If it returns true, processing should continue, if false
// we return the error to WalkDir and processing stops
type ErrorHandler func(err error) bool

// Found function type. The Found function is called whenever all the matches pass. The caller can do
// anything in this function but the path passed to it will not be recorded as matched unless it returns
// true.
type FoundFn func(path string, info fs.DirEntry) bool

// FinderError wraps internal errors in processing
type FinderError struct {
	Err   error       // underlying error
	Entry fs.DirEntry // what was being processed
	Info  string      // operation specific information
	Path  string      // path that led to the error
}

type Finder struct {
	Found                FoundFn
	WalkErrorHandler     ErrorHandler // function called when filepath.WalkFunc is called with an error
	InternalErrorHandler ErrorHandler // function called when a matcher errors
	matchers             []Matcher    // internal array of matchers to be called
	root                 string
	Paths                []string // paths matched during processing
}

func NewFinder() *Finder {
	return &Finder{
		Paths:                make([]string, 0, DefaultCapacity),
		WalkErrorHandler:     DefaultWalkErrorHandler,
		InternalErrorHandler: DefaultInternalErrorHandler,
		Found:                DefaultFound,
	}
}

// Reset clears the found paths and the stored root from a Finder struct
func (finder *Finder) Reset() {
	finder.Paths = make([]string, DefaultCapacity)
	finder.root = ""
}

// Find searches the filesystem from `root` returning a slice of matching files.
// If no matchers have been added this behaves as is called with a Name matcher
// set to '*' - as does `find`
func (finder *Finder) Find(root string) ([]string, error) {
	finder.root = root
	return finder.Paths, filepath.WalkDir(root, finder.walkFn)
}

// Name appends a matcher which selects paths based on a glob of their name (where name is the final component in the path)
// `find . -name x`
func (finder *Finder) Name(glob string) *Finder {
	finder.matchers = append(finder.matchers, Name(finder, glob))
	return finder
}

// Dir appends a matcher which selects paths which are directories only
// `find . -type d`
func (finder *Finder) Dir() *Finder {
	finder.matchers = append(finder.matchers, Dir(finder))
	return finder
}

// File appends a matcher which selects paths which are regular files only
// `find . -type f`
func (finder *Finder) File() *Finder {
	finder.matchers = append(finder.matchers, File(finder))
	return finder
}

// Depth matches a path if the depth of the path relative to the starting point is `depth`
// `find . -depth n`
func (finder *Finder) Depth(depth int) *Finder {
	finder.matchers = append(finder.matchers, Depth(finder, depth))
	return finder
}

// MaxDepth matches a path if the depth of the path relative to the starting point is `depth`
// `find . -maxdepth n`
func (finder *Finder) MaxDepth(depth int) *Finder {
	finder.matchers = append(finder.matchers, MaxDepth(finder, depth))
	return finder
}

// MinDepth matches a path if the depth of the path relative to the starting point is `depth`
// `find . -mindepth n`
func (finder *Finder) MinDepth(depth int) *Finder {
	finder.matchers = append(finder.matchers, MinDepth(finder, depth))
	return finder
}

// Or appends a Matcher which returns true if any of the supplied Matchers returns true.
func (finder *Finder) Or(matchers ...Matcher) *Finder {
	finder.matchers = append(finder.matchers, Or(finder, matchers...))
	return finder
}

// Owner appends a Matcher which returns true if the owning user of the path is the specified
// string (testing name then uid)
func (finder *Finder) Owner(name string) *Finder {
	finder.matchers = append(finder.matchers, Owner(finder, name))
	return finder
}

// Group appends a Matcher which returns true if the owning group of the path is the specified
// string (testing name then uid)
func (finder *Finder) Group(name string) *Finder {
	finder.matchers = append(finder.matchers, Owner(finder, name))
	return finder
}

// walkFn is passed to filepath.WalkDir to operate the path walk
func (finder *Finder) walkFn(path string, info fs.DirEntry, err error) error {

	// only occurs if the stat on the root fails at which point we go no further
	if info == nil {
		finder.WalkErrorHandler(err)
	}

	for _, matcher := range finder.matchers {
		matched, err := finder.callMatcher(matcher, path, info, err)
		if err != nil || !matched {
			return err
		}
	}

	if finder.Found(path, info) {
		finder.Paths = append(finder.Paths, path)
	}

	return nil
}

// callMatcher calls the supplied matcher after handling any error that might have been
// passed to the WalkFunc
func (finder *Finder) callMatcher(currentMatcher Matcher, path string, info fs.DirEntry, err error) (bool, error) {
	if err != nil && !finder.WalkErrorHandler(err) {
		return false, err
	}
	return currentMatcher(path, info)
}

// pathDepth returns the depth of the passed path compared to the root (where 0 means it *is* the root)
func (finder *Finder) pathDepth(path string) int {
	tail := strings.TrimPrefix(path, finder.root)
	components := strings.Split(tail, string(os.PathSeparator))
	if len(components) > 0 && components[0] == "" {
		return len(components) - 1
	} else {
		return len(components)
	}
}

// DefaultWalkErrorHandler writes the error to stderr and returns true
// allowing the recursion to continue.
func DefaultWalkErrorHandler(err error) bool {
	log.Printf("%+v", err)
	return true
}

// DefaultInternalErrorHandler writes the error to stderr and returns false
// forcing the process to stop.
func DefaultInternalErrorHandler(err error) bool {
	log.Printf("%+v", err)
	return false
}

// DefaultFound simply returns true to simplify the call model
func DefaultFound(path string, info fs.DirEntry) bool {
	return true
}

// Error() returns a string representation of an internal error
func (e *FinderError) Error() string {
	return fmt.Sprintf("gofind: internal error processing '%s', underlying error - '%s", e.Path, e.Err.Error())
}
