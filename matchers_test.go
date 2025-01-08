package find

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the walk process - matching all files without error
func TestWalk(t *testing.T) {
	finder := NewFinder()
	finder.Name("*")
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, len(testFS))
}

func TestNameMatcher(t *testing.T) {
	finder := NewFinder()
	matcher := Name(finder, "*.txt")

	entries, err := fs.ReadDir(testFS, "test")
	assert.Nil(t, err)

	for _, entry := range entries {
		if entry.Name() == "l1.txt" {
			matched, err := matcher("test/l1.txt", entry)
			assert.Nil(t, err)
			assert.True(t, matched)
		}
	}
}

func TestReset(t *testing.T) {
	finder := NewFinder()
	finder.Name("*")
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, len(testFS))
	assert.NotEmpty(t, finder.root)

	finder.Reset()
	assert.Zero(t, len(finder.Paths))
	assert.Empty(t, finder.root)
}

func TestDirMatcher(t *testing.T) {
	finder := NewFinder()
	matcher := Dir(finder)

	entries, err := fs.ReadDir(testFS, "test")
	assert.Nil(t, err)

	for _, entry := range entries {
		if entry.Name() == "l2" {
			matched, err := matcher("test/l2", entry)
			assert.Nil(t, err)
			assert.True(t, matched)
		}
	}
}

// Find the directories only
func TestDirFind(t *testing.T) {
	finder := NewFinder()
	finder.Dir()
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 3)
	assert.Equal(t, []string{"test", "test/l2", "test/other"}, matches)
}

// Find the directories only
func TestFileFind(t *testing.T) {
	finder := NewFinder()
	finder.File()
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 4)
	assert.Equal(t, []string{"test/l1.txt", "test/l2/l2.txt", "test/other/binary.dat", "test/other/perms.txt"}, matches)
}

func TestNameFind(t *testing.T) {
	finder := NewFinder()
	finder.Name("*.txt")

	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 3)
	assert.Equal(t, []string{"test/l1.txt", "test/l2/l2.txt", "test/other/perms.txt"}, matches)
}

// Find at a specific depth
func TestExactDepthFInd(t *testing.T) {
	finder := NewFinder()
	finder.Depth(1)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 3)
	assert.ElementsMatch(t, []string{"test/other", "test/l2", "test/l1.txt"}, matches)
}

// Find at a specific depth
func TestMaxDepthFind(t *testing.T) {
	finder := NewFinder()
	finder.MaxDepth(1)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 4)
	assert.ElementsMatch(t, []string{"test", "test/l2", "test/l1.txt", "test/other"}, matches)
}

// Find at a minimum depth
func TestMinDepthFind(t *testing.T) {
	finder := NewFinder()
	finder.MinDepth(2)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 4)
	assert.ElementsMatch(t, []string{"test/l2/l2.txt", "test/other/perms.txt", "test/other/binary.dat", "test/other/link.dat"}, matches)
}

// Find at an exact depth
func TestExactDepthFind(t *testing.T) {
	finder := NewFinder()
	finder.Depth(0)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test"}, matches)
}

// Find by type.
func TestFindType(t *testing.T) {
	finder := NewFinder()
	finder.Type(fs.ModeSymlink)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/link.dat"}, matches)
}

// Invert match test.
func TestNot(t *testing.T) {
	finder := NewFinder()
	finder.Not(File(finder))
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 4)
	assert.ElementsMatch(t, []string{"test", "test/other", "test/l2", "test/other/link.dat"}, matches)
}

// Simple prune test - more useful in complex use cases, this test
// is trivial
func TestPrune(t *testing.T) {
	finder := NewFinder().Prune()
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 0)
}
