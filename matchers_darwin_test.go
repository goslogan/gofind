//go:build darwin

package find

import (
	"io/fs"
	"syscall"
	"testing/fstest"
	"time"
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
}
