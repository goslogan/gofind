//go:build linux

package find

import (
	"io/fs"
	"os"
	"syscall"
	"testing"
	"testing/fstest"
	"time"

	"github.com/stretchr/testify/assert"
)

var testFS = fstest.MapFS{

	"testdata": &fstest.MapFile{
		ModTime: time.Unix(1735596822, 0),
		Mode:    fs.FileMode(2147484157),
		Sys: &syscall.Stat_t{
			Dev:     61,
			Ino:     26622346,
			Nlink:   4,
			Mode:    16893,
			Uid:     1000,
			Gid:     1000,
			Rdev:    0,
			Size:    4096,
			Blksize: 4096,
			Blocks:  8,
			Atim:    syscall.Timespec{Sec: 1735596823, Nsec: 116297741},
			Mtim:    syscall.Timespec{Sec: 1735596822, Nsec: 646307745},
			Ctim:    syscall.Timespec{Sec: 1735596822, Nsec: 646307745},
		},
	},

	"testdata/l1.txt": &fstest.MapFile{
		ModTime: time.Unix(1735568263, 0),
		Mode:    fs.FileMode(436),
		Sys: &syscall.Stat_t{
			Dev:     61,
			Ino:     26623511,
			Nlink:   1,
			Mode:    33204,
			Uid:     1000,
			Gid:     1000,
			Rdev:    0,
			Size:    13,
			Blksize: 4096,
			Blocks:  24,
			Atim:    syscall.Timespec{Sec: 1735568272, Nsec: 152496234},
			Mtim:    syscall.Timespec{Sec: 1735568263, Nsec: 221680703},
			Ctim:    syscall.Timespec{Sec: 1735568263, Nsec: 224680641},
		},
	},

	"testdata/l2": &fstest.MapFile{
		ModTime: time.Unix(1735568232, 0),
		Mode:    fs.FileMode(2147484157),
		Sys: &syscall.Stat_t{
			Dev:     61,
			Ino:     26623509,
			Nlink:   2,
			Mode:    16893,
			Uid:     1000,
			Gid:     1000,
			Rdev:    0,
			Size:    4096,
			Blksize: 4096,
			Blocks:  8,
			Atim:    syscall.Timespec{Sec: 1735597007, Nsec: 641428956},
			Mtim:    syscall.Timespec{Sec: 1735568232, Nsec: 799310933},
			Ctim:    syscall.Timespec{Sec: 1735596793, Nsec: 266934796},
		},
	},

	"testdata/other": &fstest.MapFile{
		ModTime: time.Unix(1735596869, 0),
		Mode:    fs.FileMode(2147484157),
		Sys: &syscall.Stat_t{
			Dev:     61,
			Ino:     26608192,
			Nlink:   2,
			Mode:    16893,
			Uid:     1000,
			Gid:     1000,
			Rdev:    0,
			Size:    4096,
			Blksize: 4096,
			Blocks:  8,
			Atim:    syscall.Timespec{Sec: 1735596874, Nsec: 915200294},
			Mtim:    syscall.Timespec{Sec: 1735596869, Nsec: 80323429},
			Ctim:    syscall.Timespec{Sec: 1735596869, Nsec: 80323429},
		},
	},

	"testdata/other/binary.dat": &fstest.MapFile{
		ModTime: time.Unix(1735568213, 0),
		Mode:    fs.FileMode(436),
		Sys: &syscall.Stat_t{
			Dev:     61,
			Ino:     26622527,
			Nlink:   1,
			Mode:    33204,
			Uid:     1000,
			Gid:     1000,
			Rdev:    0,
			Size:    2048,
			Blksize: 4096,
			Blocks:  24,
			Atim:    syscall.Timespec{Sec: 1735568220, Nsec: 448567633},
			Mtim:    syscall.Timespec{Sec: 1735568213, Nsec: 568710842},
			Ctim:    syscall.Timespec{Sec: 1735596811, Nsec: 244550688},
		},
	},

	"testdata/other/perms.txt": &fstest.MapFile{
		ModTime: time.Unix(1735568272, 0),
		Mode:    fs.FileMode(448),
		Sys: &syscall.Stat_t{
			Dev:     61,
			Ino:     26622486,
			Nlink:   1,
			Mode:    33216,
			Uid:     0,
			Gid:     1000,
			Rdev:    0,
			Size:    13,
			Blksize: 4096,
			Blocks:  24,
			Atim:    syscall.Timespec{Sec: 1735568272, Nsec: 151496254},
			Mtim:    syscall.Timespec{Sec: 1735568272, Nsec: 152496234},
			Ctim:    syscall.Timespec{Sec: 1735596888, Nsec: 873906199},
		},
	},
}

// Test the walk process - matching all files without error
func TestWalk(t *testing.T) {
	finder := NewFinder()
	finder.Name("*")
	matches, err := finder.FindFS("testdata", testFS)
	assert.Nil(t, err)
	assert.Len(t, matches, len(testFS))
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
