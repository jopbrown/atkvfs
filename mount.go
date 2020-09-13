package atkvfs

import "net/http"

func Mount(prefix string, fs http.FileSystem) {
	globalMountedFSManager.Mount(prefix, fs)
}

func Unmount(prefix string) {
	globalMountedFSManager.Unmount(prefix)
}
