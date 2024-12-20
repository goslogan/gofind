package find

import (
	"os"
	"syscall"
	"testing"
	"testing/fstest"
	"time"

	"github.com/stretchr/testify/assert"
)

var baseTime, _ = time.Parse(time.RFC3339, "2024-01-01T00:00:00Z")

var testFS = fstest.MapFS{
	"l1.txt": &fstest.MapFile{
		ModTime: baseTime.Add(400 * time.Hour),
		Mode:    0644,
		Sys:     &syscall.Stat_t{},
	},
}

// Test the walk process - matching all files without error
func TestWalk(t *testing.T) {
	finder := NewFinder()
	finder.Name("*")
	matches, err := finder.FindFS("/", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 4)
}

func TestNameMatcher(t *testing.T) {
	finder := NewFinder()
	matcher := Name(finder, "*.txt")

	entries, err := os.ReadDir("testdata")
	assert.Nil(t, err)

	for _, entry := range entries {
		if entry.Name() == "l1.txt" {
			matched, err := matcher("testdata/l1.txt", entry)
			assert.Nil(t, err)
			assert.True(t, matched)
		}
	}
}

func TestDirMatcher(t *testing.T) {
	finder := NewFinder()
	matcher := Dir(finder)

	entries, err := os.ReadDir("testdata")
	assert.Nil(t, err)

	for _, entry := range entries {
		if entry.Name() == "l2" {
			matched, err := matcher("testdata/l2", entry)
			assert.Nil(t, err)
			assert.True(t, matched)
		}
	}
}

// Find the directories only
func TestDirFind(t *testing.T) {
	finder := NewFinder()
	finder.Dir()
	matches, err := finder.Find("testdata")
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.Equal(t, []string{"testdata", "testdata/l2"}, matches)
}

// Find the directories only
func TestFileFind(t *testing.T) {
	finder := NewFinder()
	finder.File()
	matches, err := finder.Find("testdata")
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.Equal(t, []string{"testdata/l1.txt", "testdata/l2/l2.txt"}, matches)
}

func TestNameFind(t *testing.T) {
	finder := NewFinder()
	finder.Name("*.txt")

	paths, err := finder.Find("testdata")
	assert.Nil(t, err)
	assert.Len(t, paths, 2)
}

// Find at a specific depth
func TestExactDepthFInd(t *testing.T) {
	finder := NewFinder()
	finder.Depth(1)
	matches, err := finder.Find("testdata")
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, []string{"testdata/l2", "testdata/l1.txt"}, matches)
}

// Find at a specific depth
func TestMaxDepthFInd(t *testing.T) {
	finder := NewFinder()
	finder.MaxDepth(1)
	matches, err := finder.Find("testdata")
	assert.Nil(t, err)
	assert.Len(t, matches, 3)
	assert.ElementsMatch(t, []string{"testdata", "testdata/l2", "testdata/l1.txt"}, matches)
}

// Find at a specific depth
func TestMinDepthFInd(t *testing.T) {
	finder := NewFinder()
	finder.MinDepth(2)
	matches, err := finder.Find("testdata")
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"testdata/l2/l2.txt"}, matches)
}

// Find by owner -
