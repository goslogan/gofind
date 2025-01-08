// Package find provides a wrapper around fs.WalkDir that uses matching functions to
package find

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/djherbis/times"
)

type FileTimeType rune

const (
	Created  FileTimeType = 'B'
	Accessed FileTimeType = 'a'
	Modified FileTimeType = 'm'
	Changed  FileTimeType = 'c'
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
// nil.
// The current directory can be skipped by returning fs.SkipDir. The recursion can be terminated by
// returning fs.SkipAll. Returning find.SkipThis will skip the current file only.
type FoundFn func(path string, info fs.DirEntry) error

// FinderError wraps internal errors in processing
type FinderError struct {
	Matcher string      // name of the matcher that raised the error
	Err     error       // underlying error
	Entry   fs.DirEntry // what was being processed
	Info    string      // operation specific information
	Path    string      // path that led to the error
}

type Finder struct {
	Found                FoundFn        // Set the Found function to control if the path matched is recorded.
	WalkErrorHandler     ErrorHandler   // function called when filepath.WalkFunc is called with an error
	InternalErrorHandler ErrorHandler   // function called when a matcher errors
	matchers             []Matcher      // internal array of matchers to be called
	Paths                []string       // paths matched during processing
	CacheCmpFile         bool           // if false the comparison file for NewerXY will not be cached but calculated on each call
	cmpFileTime          times.Timespec // Cache time data for comparison file
	root                 string         // keep track of the root during processing to get relative paths
	started              time.Time      // keep track of when Find started
}

// SkipThis can be returned to tell the walkder to skip this file (don't add to the paths returned)
var SkipThis = errors.New("skip this file")

// NewFinder creates and initialises a new Finder struct.
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
	finder.Paths = make([]string, 0, DefaultCapacity)
	finder.root = ""
	finder.started = time.Time{}
}

// Find searches the filesystem from `root` returning a slice of matching files.
// If no matchers have been added this behaves as is called with a Name matcher
// set to '*' - as does `find`
func (finder *Finder) Find(root string) ([]string, error) {
	return finder.FindFS(root, os.DirFS(root))
}

// FindFS searches the filesystem provided returning a slice of matching files.
// If no matchers have been added this behaves as is called with a Name matcher
// set to '*' - as does `find`
func (finder *Finder) FindFS(root string, rootFS fs.FS) ([]string, error) {
	finder.root = root
	finder.started = time.Now()
	return finder.Paths, fs.WalkDir(rootFS, root, finder.walkFn)
}

// Depth matches a path if the depth of the path relative to the starting point is `depth`
// `find . -depth n`
func (finder *Finder) Depth(depth int) *Finder {
	return finder.appendMatcher(Depth(finder, depth))
}

// MaxDepth matches a path if the depth of the path relative to the starting point is `depth`
// `find . -maxdepth n`
func (finder *Finder) MaxDepth(depth int) *Finder {
	return finder.appendMatcher(MaxDepth(finder, depth))
}

// MinDepth matches a path if the depth of the path relative to the starting point is `depth`
// `find . -mindepth n`
func (finder *Finder) MinDepth(depth int) *Finder {
	return finder.appendMatcher(MinDepth(finder, depth))
}

// Owner appends a Matcher which returns true if the owning user of the path is the specified
// string (testing name then uid)
func (finder *Finder) Owner(name string) *Finder {
	return finder.appendMatcher(Owner(finder, name))
}

// Group appends a Matcher which returns true if the owning group of the path is the specified
// string (testing name then uid)
func (finder *Finder) Group(name string) *Finder {
	return finder.appendMatcher(Owner(finder, name))
}

// Prune returns a Matcher which cause the current directory to be skipped
func (finder *Finder) Prune() *Finder {
	return finder.appendMatcher(Prune(finder))
}

// CallInternalErrorHandler wraps the internal error handler property call in order
// to handle skip type errors (which we have to always return).
func (finder *Finder) CallInternalErrorHandler(err error) error {
	if err == fs.SkipDir || err == fs.SkipAll || err == SkipThis {
		return err
	}
	if finder.InternalErrorHandler(err) {
		return nil
	}
	return err
}

// walkFn is passed to filepath.WalkDir to operate the path walk
func (finder *Finder) walkFn(path string, info fs.DirEntry, walkErr error) error {

	// only occurs if the stat on the root fails at which point we go no further
	if info == nil {
		finder.WalkErrorHandler(walkErr)
		return walkErr
	}

	if walkErr != nil && finder.WalkErrorHandler(walkErr) {
		return walkErr
	}

	for _, matcher := range finder.matchers {
		matched, err := matcher(path, info)
		if !matched || (err != nil && finder.WalkErrorHandler(err)) {
			return err
		}
	}

	err := finder.Found(path, info)
	if err == nil {
		finder.Paths = append(finder.Paths, path)
	}
	return err
}

// appendMatcher adds the provided matcher to the finder and returns the finder.
func (finder *Finder) appendMatcher(matcher Matcher) *Finder {
	finder.matchers = append(finder.matchers, matcher)
	return finder
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
func DefaultFound(path string, info fs.DirEntry) error {
	return nil
}

// Error() returns a string representation of an internal error
func (e *FinderError) Error() string {
	return fmt.Sprintf("gofind: internal error processing '%s', underlying error - '%s", e.Path, e.Err.Error())
}
