//go:build darwin

package find

import (
	"io/fs"
	"syscall"
	"testing"
	"testing/fstest"
	"time"

	"github.com/stretchr/testify/assert"
)

var testFS = fstest.MapFS{

	"test": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 6, 12, 19, 40, 0, time.UTC),
		Mode:    fs.FileMode(fs.ModeDir | 0o755),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8538125,
			Nlink:         5,
			Mode:          0x41ED,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          160,
			Blksize:       4096,
			Blocks:        0,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 19, 40, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 19, 40, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 19, 40, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
		},
	},

	"test/l1.txt": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC),
		Mode:    fs.FileMode(0o644),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8538126,
			Nlink:         1,
			Mode:          0x81A4,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          13,
			Blksize:       4096,
			Blocks:        8,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 13, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
		},
	},

	"test/l2": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC),
		Mode:    fs.FileMode(fs.ModeDir | 0o755),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8557940,
			Nlink:         3,
			Mode:          0x41ED,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          96,
			Blksize:       4096,
			Blocks:        0,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 19, 40, 0, time.UTC).Unix()},
		},
	},

	"test/l2/l2.txt": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC),
		Mode:    fs.FileMode(0o644),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8557962,
			Nlink:         1,
			Mode:          0x81A4,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          8,
			Blksize:       4096,
			Blocks:        8,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 16, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
		},
	},

	"test/other": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC),
		Mode:    fs.FileMode(fs.ModeDir | 0o755),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8538127,
			Nlink:         4,
			Mode:          0x41ED,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          128,
			Blksize:       4096,
			Blocks:        0,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 14, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
		},
	},

	"test/other/binary.dat": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC),
		Mode:    fs.FileMode(0o644),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8538128,
			Nlink:         1,
			Mode:          0x81A4,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          2048,
			Blksize:       4096,
			Blocks:        8,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
		},
	},

	"test/other/perms.txt": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC),
		Mode:    fs.FileMode(0o755),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8538129,
			Nlink:         1,
			Mode:          0x81ED,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          13,
			Blksize:       4096,
			Blocks:        8,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 13, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
		},
	},
}

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
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, []string{"test/l2", "test/l1.txt"}, matches)
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
	assert.Len(t, matches, 3)
	assert.ElementsMatch(t, []string{"test/l2/l2.txt", "test/other/perms.txt", "test/other/binary.dat"}, matches)
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

// Find by owner -
