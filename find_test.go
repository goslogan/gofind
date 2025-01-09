package find

import (
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

// Find at a specific depth
func TestExactDepthFInd(t *testing.T) {
	finder := NewFinder()
	finder.Depth(1)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 5)
	assert.ElementsMatch(t, []string{"test/empty", "test/other", "test/l2", "test/l1.txt", "test/data.csv"}, matches)
}

// Find at a specific depth
func TestMaxDepthFind(t *testing.T) {
	finder := NewFinder()
	finder.MaxDepth(1)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 6)
	assert.ElementsMatch(t, []string{"test/empty", "test", "test/l2", "test/l1.txt", "test/other", "test/data.csv"}, matches)
}

// Find at a minimum depth
func TestMinDepthFind(t *testing.T) {
	finder := NewFinder()
	finder.MinDepth(2)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 6)
	assert.ElementsMatch(t, []string{"test/other/zero.dat", "test/l2/l2.txt", "test/other/perms.txt", "test/other/binary.dat", "test/other/link.dat", "test/other/DATA.csv"}, matches)
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

// Invert match test.
func TestNot(t *testing.T) {
	finder := NewFinder()
	finder.Not(File(finder))
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 5)
	assert.ElementsMatch(t, []string{"test/empty", "test", "test/other", "test/l2", "test/other/link.dat"}, matches)
}

// Simple prune test - more useful in complex use cases, this test
// is trivial
func TestPrune(t *testing.T) {
	finder := NewFinder().Prune()
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 0)
}
