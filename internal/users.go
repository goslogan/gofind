package internal

import (
	"fmt"
	"io/fs"
	"os/user"
	"syscall"
)

// FileOwnerUser tries to get the user owning a file starting with an fs.DirInfo value
func FileOwnerUser(path string, entry fs.FileInfo) (*user.User, error) {
	if sys := entry.Sys(); sys == nil {
		return nil, fmt.Errorf("gofind: unable to get system stat info for %s", path)
	} else if stat, ok := sys.(*syscall.Stat_t); ok {
		return user.LookupId(fmt.Sprintf("%d", stat.Uid))
	} else {
		return nil, fmt.Errorf("gofind: unable to cast system stat info for %s", path)
	}
}

// FileOwnerGroup tries to get the name the group owning a file starting with an fs.DirInfo value
func FileOwnerGroup(path string, entry fs.FileInfo) (*user.Group, error) {
	if sys := entry.Sys(); sys == nil {
		return nil, fmt.Errorf("gofind: unable to get system stat info for %s", path)
	} else if stat, ok := sys.(*syscall.Stat_t); ok {
		return user.LookupGroupId(fmt.Sprintf("%d", stat.Gid))
	} else {
		return nil, fmt.Errorf("gofind: unable to cast system stat info for %s", path)
	}
}
