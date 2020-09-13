package atkvfs

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

var (
	globalMountedFSManager  = newMountedFSManager()
	globalOpenedFileManager = newOpenedFileManager()
)

type mountedFSManager struct {
	fsMap map[string]http.FileSystem
}

func newMountedFSManager() *mountedFSManager {
	return &mountedFSManager{make(map[string]http.FileSystem)}
}

func (m *mountedFSManager) GetFS(fname string) http.FileSystem {
	for prefix, fs := range m.fsMap {
		if strings.HasPrefix(path.Clean(fname), prefix) {
			return fs
		}
	}

	return nil
}

func (m *mountedFSManager) Open(fname string) (http.File, error) {
	fs := m.GetFS(fname)
	if nil == fs {
		return nil, fmt.Errorf("file not register in vfs: %s", fname)
	}

	return fs.Open(fname)
}

func (m *mountedFSManager) Mount(prefix string, fs http.FileSystem) {
	m.fsMap[prefix] = fs
}

func (m *mountedFSManager) Unmount(prefix string) {
	delete(m.fsMap, prefix)
}

type openedFileManager struct {
	fileMap map[uintptr]http.File
	id      uintptr
}

func newOpenedFileManager() *openedFileManager {
	return &openedFileManager{make(map[uintptr]http.File), 1}
}

func (m *openedFileManager) Register(f http.File) uintptr {
	m.id = m.id + 1
	m.fileMap[m.id] = f
	return m.id
}

func (m *openedFileManager) UnRegister(id uintptr) {
	delete(m.fileMap, id)
}

func (m *openedFileManager) GetFile(id uintptr) http.File {
	return m.fileMap[id]
}

func vfsScanFiles(p string) (paths []string, err error) {
	f, err := globalMountedFSManager.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	paths = []string{p}
	if !stat.IsDir() {
		return paths, nil
	}

	subs, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, sub := range subs {
		subPaths, err := vfsScanFiles(path.Join(p, sub.Name()))
		if err != nil {
			return nil, err
		}

		paths = append(paths, subPaths...)
	}

	return paths, nil
}
