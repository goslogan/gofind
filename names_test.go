package find

import (
	"errors"
	"io/fs"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameMatcher(t *testing.T) {
	finder := NewFinder()
	matcher := Name(finder, "*.txt")

	entries, err := fs.ReadDir(testFS, "test")
	assert.Nil(t, err)

	for _, entry := range entries {
		info, err := entry.Info()
		assert.Nil(t, err)
		if entry.Name() == "l1.txt" {
			matched, err := matcher("test/l1.txt", info)
			assert.Nil(t, err)
			assert.True(t, matched)
		}
	}
}

func TestBadGlobFind(t *testing.T) {
	finder := NewFinder()
	finder.Name("[")
	matches, err := finder.FindFS("test", testFS)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, filepath.ErrBadPattern))
	assert.Empty(t, matches)
}

func TestBadIGlobFind(t *testing.T) {
	finder := NewFinder()
	finder.IName("[")
	matches, err := finder.FindFS("test", testFS)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, filepath.ErrBadPattern))
	assert.Empty(t, matches)
}

func TestNameFind(t *testing.T) {
	finder := NewFinder()
	finder.Name("*.txt")

	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 3)
	assert.ElementsMatch(t, []string{"test/l1.txt", "test/l2/l2.txt", "test/other/perms.txt"}, matches)
}

func TestINameFind(t *testing.T) {
	finder := NewFinder()
	finder.IName("data.*")

	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, []string{"test/data.csv", "test/other/DATA.csv"}, matches)
}

func TestPathFind(t *testing.T) {
	finder := NewFinder()
	finder.Path("test/*/*.txt")

	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, []string{"test/other/perms.txt", "test/l2/l2.txt"}, matches)
}

func TestPathBadGlobFind(t *testing.T) {
	finder := NewFinder()
	finder.Path("test/[.csv")

	matches, err := finder.FindFS("test", testFS)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, filepath.ErrBadPattern))
	assert.Empty(t, matches)
}

func TestIPathFind(t *testing.T) {
	finder := NewFinder()
	finder.IPath("test/*/*.txt")

	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, matches, []string{"test/other/perms.txt", "test/l2/l2.txt"})
}

func TestIPathBadGlobFind(t *testing.T) {
	finder := NewFinder()
	finder.IPath("test/[.csv")

	matches, err := finder.FindFS("test", testFS)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, filepath.ErrBadPattern))
	assert.Empty(t, matches)
}

func TestRegexFind(t *testing.T) {
	re := regexp.MustCompile(`(?i).+data.+`)
	finder := NewFinder()
	finder.Regex(re)

	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, matches, []string{"test/data.csv", "test/other/DATA.csv"})
}
