package atkvfs

import "io/fs"

func Mount(prefix string, fs fs.FS) {
	globalMountedFSManager.Mount(prefix, fs)
}

func Unmount(prefix string) {
	globalMountedFSManager.Unmount(prefix)
}
