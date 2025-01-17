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
		ModTime: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC),
		Mode:    fs.FileMode(fs.ModeDir | 0o755),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8538125,
			Nlink:         7,
			Mode:          0x41ED,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          224,
			Blksize:       4096,
			Blocks:        0,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
		},
	},

	"test/data.csv": &fstest.MapFile{
		Data: []byte{0x4e, 0x61, 0x6d, 0x65, 0x2c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0xa, 0x4e, 0x69, 0x63, 0x2c,
			0x31, 0x32, 0xa, 0x45, 0x6c, 0x69, 0x73, 0x61, 0x62, 0x65, 0x74, 0x68, 0x2c, 0x32, 0x37, 0xa, 0x52,
			0x61, 0x68, 0x75, 0x6c, 0x2c, 0x39, 0xa},
		ModTime: time.Date(2025, time.January, 8, 15, 5, 27, 0, time.UTC),
		Mode:    fs.FileMode(0o644),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8885237,
			Nlink:         1,
			Mode:          0x81A4,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          39,
			Blksize:       4096,
			Blocks:        8,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 18, 33, 50, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 15, 5, 27, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 15, 5, 27, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 8, 15, 5, 27, 0, time.UTC).Unix()},
		},
	},

	"test/empty": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC),
		Mode:    fs.FileMode(fs.ModeDir | 0o755),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8944200,
			Nlink:         2,
			Mode:          0x41ED,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          64,
			Blksize:       4096,
			Blocks:        0,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 9, 7, 47, 24, 0, time.UTC).Unix()},
		},
	},

	"test/l1.txt": &fstest.MapFile{
		Data:    []byte{0x6c, 0x65, 0x76, 0x65, 0x6c, 0x20, 0x31, 0x20, 0x66, 0x69, 0x6c, 0x65, 0xa},
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
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 13, 4, 0, 0, time.UTC).Unix()},
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
		Data:    []byte{0x6c, 0x32, 0x20, 0x66, 0x69, 0x6c, 0x65, 0xa},
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
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 18, 33, 50, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 12, 20, 14, 0, time.UTC).Unix()},
		},
	},

	"test/other": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 9, 17, 45, 20, 0, time.UTC),
		Mode:    fs.FileMode(fs.ModeDir | 0o755),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8538127,
			Nlink:         7,
			Mode:          0x41ED,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          224,
			Blksize:       4096,
			Blocks:        0,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 17, 45, 20, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 17, 45, 20, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 17, 45, 20, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
		},
	},

	"test/other/DATA.csv": &fstest.MapFile{
		Data: []byte{0x4e, 0x61, 0x6d, 0x65, 0x2c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0xa, 0x4e, 0x69, 0x63, 0x2c, 0x31,
			0x32, 0xa, 0x45, 0x6c, 0x69, 0x73, 0x61, 0x62, 0x65, 0x74, 0x68, 0x2c, 0x32, 0x37, 0xa, 0x52, 0x61, 0x68,
			0x75, 0x6c, 0x2c, 0x39, 0xa},
		ModTime: time.Date(2025, time.January, 8, 15, 5, 37, 0, time.UTC),
		Mode:    fs.FileMode(0o644),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8885246,
			Nlink:         1,
			Mode:          0x81A4,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          39,
			Blksize:       4096,
			Blocks:        8,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 18, 33, 50, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 15, 5, 37, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 15, 5, 37, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 8, 15, 5, 37, 0, time.UTC).Unix()},
		},
	},

	"test/other/binary.dat": &fstest.MapFile{
		Data: []byte{0x4b, 0xf7, 0x80, 0xc7, 0x68, 0x7a, 0x88, 0x56, 0xd8, 0xaf, 0xc4, 0xd6, 0xc3, 0xcd, 0xcf, 0xbd,
			0x8a, 0x9, 0xd9, 0x7f, 0xa8, 0x55, 0xee, 0x47, 0x12, 0x77, 0x5e, 0x5a, 0xb6, 0x5, 0x63, 0xdc, 0xd9, 0xa8,
			0x41, 0x92, 0x2b, 0x87, 0x97, 0xcb, 0x88, 0x3e, 0x76, 0x2, 0xbd, 0x2a, 0x8e, 0x58, 0xeb, 0xd7, 0x94, 0x8e,
			0xb8, 0xf8, 0x21, 0xb6, 0x90, 0xe7, 0x6e, 0x2f, 0x9e, 0x43, 0xc8, 0x48, 0x1c, 0x60, 0xb0, 0x51, 0x31, 0x71,
			0x5b, 0x98, 0xa9, 0x8a, 0x37, 0xfe, 0x7c, 0xf2, 0xda, 0x62, 0xc5, 0xe0, 0xbd, 0xe2, 0x8c, 0xcb, 0xa4, 0x7a,
			0x21, 0x4c, 0xdf, 0xdb, 0xfe, 0xdc, 0x5e, 0x35, 0x79, 0xbb, 0x85, 0x15, 0x82, 0x45, 0x45, 0x9b, 0x5c, 0xc7,
			0x58, 0xce, 0xce, 0xac, 0xc4, 0xee, 0x42, 0x79, 0x97, 0xad, 0x90, 0x6f, 0xf9, 0xd9, 0xfa, 0xf9, 0xc7, 0x2c,
			0x5e, 0xfe, 0xa1, 0xb1},
		ModTime: time.Date(2025, time.January, 9, 17, 45, 2, 0, time.UTC),
		Mode:    fs.FileMode(0o644),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8538128,
			Nlink:         1,
			Mode:          0x81A4,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          128,
			Blksize:       4096,
			Blocks:        8,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 7, 18, 33, 50, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 7, 15, 5, 37, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 7, 15, 5, 37, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 7, 15, 5, 37, 0, time.UTC).Unix()},
		},
	},

	"test/other/link.dat": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 7, 7, 33, 49, 0, time.UTC),
		Mode:    fs.FileMode(fs.ModeSymlink | 0o755),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8654636,
			Nlink:         1,
			Mode:          0xA1ED,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          10,
			Blksize:       4096,
			Blocks:        0,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 7, 7, 33, 49, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 7, 7, 33, 49, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 7, 7, 33, 49, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 7, 7, 33, 49, 0, time.UTC).Unix()},
		},
	},

	"test/other/perms.txt": &fstest.MapFile{
		Data:    []byte{0x6c, 0x65, 0x76, 0x65, 0x6c, 0x20, 0x31, 0x20, 0x66, 0x69, 0x6c, 0x65, 0xa},
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
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 13, 4, 0, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 6, 11, 9, 12, 0, time.UTC).Unix()},
		},
	},

	"test/other/zero.dat": &fstest.MapFile{
		ModTime: time.Date(2025, time.January, 8, 23, 15, 1, 0, time.UTC),
		Mode:    fs.FileMode(0o644),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8937723,
			Nlink:         1,
			Mode:          0x81A4,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          0,
			Blksize:       4096,
			Blocks:        0,
			Atimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 23, 17, 1, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 8, 23, 15, 1, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2025, time.January, 9, 7, 21, 33, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2025, time.January, 8, 23, 15, 1, 0, time.UTC).Unix()},
		},
	},

	"test/other/sparsefile.dat": &fstest.MapFile{
		ModTime: time.Date(2023, time.December, 1, 11, 12, 22, 0, time.UTC),
		Mode:    fs.FileMode(0o644),
		Sys: &syscall.Stat_t{
			Dev:           16777232,
			Ino:           8937723,
			Nlink:         1,
			Mode:          0x81A4,
			Uid:           502,
			Gid:           20,
			Rdev:          0,
			Size:          65536,
			Blksize:       4096,
			Blocks:        8,
			Atimespec:     syscall.Timespec{Sec: time.Date(2023, time.December, 2, 8, 43, 19, 0, time.UTC).Unix()},
			Mtimespec:     syscall.Timespec{Sec: time.Date(2023, time.December, 1, 11, 12, 22, 0, time.UTC).Unix()},
			Ctimespec:     syscall.Timespec{Sec: time.Date(2023, time.December, 1, 11, 12, 22, 0, time.UTC).Unix()},
			Birthtimespec: syscall.Timespec{Sec: time.Date(2023, time.December, 1, 10, 9, 1, 0, time.UTC).Unix()},
		},
	},
}

var newFile = fstest.MapFile{
	ModTime: time.Now(),
	Mode:    fs.FileMode(0o644),
	Sys: &syscall.Stat_t{
		Dev:           16777232,
		Ino:           8947723,
		Nlink:         1,
		Mode:          0x81A4,
		Uid:           502,
		Gid:           20,
		Rdev:          0,
		Size:          0,
		Blksize:       4096,
		Blocks:        0,
		Atimespec:     syscall.Timespec{Sec: time.Now().Unix()},
		Mtimespec:     syscall.Timespec{Sec: time.Now().Unix()},
		Ctimespec:     syscall.Timespec{Sec: time.Now().Unix()},
		Birthtimespec: syscall.Timespec{Sec: time.Now().Unix()},
	},
}

func TestBMinLess(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Birthtimespec = syscall.Timespec{Sec: fTime.Add(time.Minute).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})

	finder := NewFinder()
	finder.Bmin(2*time.Minute, LessThan)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/newfile.dat"}, matches)
}

func TestBMinGreater(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Birthtimespec = syscall.Timespec{Sec: fTime.Add(-5000 * time.Hour).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})
	finder := NewFinder()
	finder.Bmin(4500*time.Hour, GreaterThan)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, []string{"test/other/sparsefile.dat", "test/other/newfile.dat"}, matches)
}

func TestBMinEqual(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Birthtimespec = syscall.Timespec{Sec: fTime.Add(500 * time.Hour).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})
	finder := NewFinder()
	finder.Bmin(500*time.Hour, Equal)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/newfile.dat"}, matches)
}

func TestAminLess(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Atimespec = syscall.Timespec{Sec: fTime.Add(time.Minute).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})

	finder := NewFinder()
	finder.Amin(2*time.Minute, LessThan)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/newfile.dat"}, matches)
}

func TestAminGreater(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Atimespec = syscall.Timespec{Sec: fTime.Add(-5000 * time.Hour).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})
	finder := NewFinder()
	finder.Amin(4500*time.Hour, GreaterThan)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, []string{"test/other/sparsefile.dat", "test/other/newfile.dat"}, matches)
}

func TestAminEqual(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Atimespec = syscall.Timespec{Sec: fTime.Add(500 * time.Hour).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})
	finder := NewFinder()
	finder.Amin(500*time.Hour, Equal)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/newfile.dat"}, matches)
}

func TestCminLess(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Ctimespec = syscall.Timespec{Sec: fTime.Add(time.Minute).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})

	finder := NewFinder()
	finder.Cmin(2*time.Minute, LessThan)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/newfile.dat"}, matches)
}

func TestCminGreater(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Ctimespec = syscall.Timespec{Sec: fTime.Add(-5000 * time.Hour).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})
	finder := NewFinder()
	finder.Cmin(4500*time.Hour, GreaterThan)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, []string{"test/other/sparsefile.dat", "test/other/newfile.dat"}, matches)
}

func TestCminEqual(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Ctimespec = syscall.Timespec{Sec: fTime.Add(500 * time.Hour).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})
	finder := NewFinder()
	finder.Cmin(500*time.Hour, Equal)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/newfile.dat"}, matches)
}

func TestMminLess(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Mtimespec = syscall.Timespec{Sec: fTime.Add(time.Minute).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})

	finder := NewFinder()
	finder.Mmin(2*time.Minute, LessThan)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/newfile.dat"}, matches)
}

func TestMminGreater(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Mtimespec = syscall.Timespec{Sec: fTime.Add(-5000 * time.Hour).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})
	finder := NewFinder()
	finder.Mmin(4500*time.Hour, GreaterThan)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 2)
	assert.ElementsMatch(t, []string{"test/other/sparsefile.dat", "test/other/newfile.dat"}, matches)
}

func TestMminEqual(t *testing.T) {

	fTime := time.Now()
	newFile.Sys.(*syscall.Stat_t).Mtimespec = syscall.Timespec{Sec: fTime.Add(500 * time.Hour).Unix()}
	cpy := copyFSAndAdd(fstest.MapFS{"test/other/newfile.dat": &newFile})
	finder := NewFinder()
	finder.Mmin(500*time.Hour, Equal)
	matches, err := finder.FindFS("test", cpy)
	assert.Nil(t, err)
	assert.Len(t, matches, 1)
	assert.ElementsMatch(t, []string{"test/other/newfile.dat"}, matches)
}

func copyFSAndAdd(additional fstest.MapFS) fstest.MapFS {

	cpy := fstest.MapFS{}
	for k, v := range testFS {
		cpy[k] = v
	}
	for k, v := range additional {
		cpy[k] = v
	}

	return cpy
}
