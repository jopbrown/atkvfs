package atkvfs

/*
#include <stdlib.h>
#include <errno.h>
#include <string.h>
#include <sys/stat.h>
#include <tcl.h>

static void stat_set_time(Tcl_StatBuf *stat, time_t t) {
	stat->st_atime = t;
	stat->st_mtime = t;
	stat->st_ctime = t;
}
*/
import "C"
import (
	"unsafe"
)

//export _go_FSPathInFilesystemProc
func _go_FSPathInFilesystemProc(pathPtr *C.Tcl_Obj) C.int {
	p := objToString(nil, pathPtr)
	if globalMountedFSManager.GetFS(p) != nil {
		return 0
	}

	return -1
}

//export _go_FSStatProc
func _go_FSStatProc(pathPtr *C.Tcl_Obj, statPtr *C.Tcl_StatBuf) C.int {
	p := objToString(nil, pathPtr)
	f, err := globalMountedFSManager.Open(p)
	if err != nil {
		return -1
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return -1
	}

	statPtr.st_mode = C.uint(stat.Mode())
	statPtr.st_size = C.long(stat.Size())
	modTime := stat.ModTime().Unix()
	C.stat_set_time(statPtr, C.long(modTime))

	return 0
}

//export _go_FSAccessProc
func _go_FSAccessProc(pathPtr *C.Tcl_Obj, mode C.int) C.int {
	if (mode & 3) != 0 {
		return -1
	}

	p := objToString(nil, pathPtr)
	f, err := globalMountedFSManager.Open(p)
	if err != nil {
		return -1
	}
	defer f.Close()
	return 0
}

//export _go_FSOpenFileChannelProc
func _go_FSOpenFileChannelProc(interp *C.Tcl_Interp, pathPtr *C.Tcl_Obj, mode, permissions C.int) C.Tcl_Channel {
	var n C.int
	cpath := C.Tcl_GetStringFromObj(pathPtr, &n)
	gpath := C.GoStringN(cpath, n)
	f, err := globalMountedFSManager.Open(gpath)
	if err != nil {
		return nil
	}

	id := globalOpenedFileManager.Register(f)

	return C.Tcl_CreateChannel((tclChannelCallbacks), cpath, C.ClientData(unsafe.Pointer(id)), C.TCL_READABLE)
}

//export _go_FSMatchInDirectoryProc
func _go_FSMatchInDirectoryProc(interp *C.Tcl_Interp, resultPtr *C.Tcl_Obj, pathPtr *C.Tcl_Obj, pattern *C.char, types *C.Tcl_GlobTypeData) C.int {
	gpath := objToString(nil, pathPtr)
	paths, err := vfsScanFiles(gpath)
	if err != nil {
		// cs := C.CString(err.String())
		// C.Tcl_AppendResult(interp, "vfsGlob err: ", cs, 0)
		// C.free(unsafe.Pointer(cs))
		return -1
	}

	for _, p := range paths {
		cs := C.CString(p)

		if C.Tcl_StringCaseMatch(cs, pattern, 0) != 0 {
			obj := C.Tcl_NewStringObj(cs, C.int(len(p)))
			C.Tcl_ListObjAppendElement(interp, resultPtr, obj)
		}

		C.free(unsafe.Pointer(cs))
	}

	return 0
}

//export _go_DriverCloseProc
func _go_DriverCloseProc(clientData unsafe.Pointer, interp *C.Tcl_Interp) C.int {
	id := uintptr(clientData)
	f := globalOpenedFileManager.GetFile(id)
	if f == nil {
		return 0
	}

	f.Close()
	globalOpenedFileManager.UnRegister(id)
	return 0
}

//export _go_DriverInputProc
func _go_DriverInputProc(clientData unsafe.Pointer, cbuf *C.char, bufSize C.int, errorCodePtr *C.int) C.int {
	id := uintptr(clientData)
	f := globalOpenedFileManager.GetFile(id)
	if f == nil {
		*errorCodePtr = C.EINVAL
		return -1
	}

	gbuf := make([]byte, int(bufSize))

	*errorCodePtr = 0

	n, err := f.Read(gbuf)
	C.memcpy(unsafe.Pointer(cbuf), unsafe.Pointer(&gbuf[0]), C.size_t(n))

	if err != nil {
		*errorCodePtr = C.EIO
	}

	return C.int(n)
}

//export _go_DriverSeekProc
func _go_DriverSeekProc(clientData unsafe.Pointer, offset C.long, seekMode C.int, errorCodePtr *C.int) C.int {
	id := uintptr(clientData)
	f := globalOpenedFileManager.GetFile(id)
	if f == nil {
		*errorCodePtr = C.EINVAL
		return -1
	}

	n, err := f.Seek(int64(offset), int(seekMode))
	if err != nil {
		*errorCodePtr = C.EIO
	}
	return C.int(n)
}
