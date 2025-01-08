package find

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.ElementsMatch(t, []string{"test", "test/l2", "test/other"}, matches)
}

// Find the regular files only
func TestFileFind(t *testing.T) {
	finder := NewFinder()
	finder.File()
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 6)
	assert.ElementsMatch(t, []string{"test/l1.txt", "test/data.csv", "test/l2/l2.txt", "test/other/binary.dat", "test/other/perms.txt", "test/other/DATA.csv"}, matches)
}

// Find by type
func TestFindType(t *testing.T) {
	finder := NewFinder()
	finder.Type(fs.ModeSymlink)
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/link.dat"}, matches)
}
