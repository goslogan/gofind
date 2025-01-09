package find

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewerCreate(t *testing.T) {
	finder := NewFinder()
	finder.Newer(Created, "test/data.csv")
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 3)
	assert.ElementsMatch(t, []string{"test/empty", "test/other/zero.dat", "test/other/DATA.csv"}, matches)
}

// Test find newer files whilst cacheing the timestamps for the comparison
// file
func TestNewerCreateWithCacheing(t *testing.T) {
	finder := NewFinder()
	finder.CacheCmpFile = true
	finder.Newer(Created, "test/data.csv")
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 3)
	assert.ElementsMatch(t, []string{"test/empty", "test/other/zero.dat", "test/other/DATA.csv"}, matches)
}

func TestNewerAccessed(t *testing.T) {
	finder := NewFinder()
	finder.Newer(Accessed, "test/data.csv")
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 4)
	assert.ElementsMatch(t, []string{"test", "test/empty", "test/other", "test/other/zero.dat"}, matches)
}

func TestNewerAccessedNotExisting(t *testing.T) {
	finder := NewFinder()
	finder.Newer(Accessed, "test/badfile..csv")
	matches, err := finder.FindFS("test", testFS)
	assert.Error(t, err)
	assert.Len(t, matches, 0)
}

func TestNewChangedTime(t *testing.T) {
	finder := NewFinder()
	finder.Newer(Changed, "test/other/DATA.csv")
	matches, err := finder.FindFS("test", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, 4)
	assert.ElementsMatch(t, []string{"test", "test/empty", "test/other", "test/other/zero.dat"}, matches)
}
