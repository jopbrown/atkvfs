package atkvfs

/*
#cgo linux CFLAGS: -I/usr/include/tcl
#cgo linux LDFLAGS: -ltcl -ltk -lX11 -lm -lz -ldl

#include <errno.h>
#include <string.h>
#include <sys/stat.h>
#include <tcl.h>

extern int _go_DriverCloseProc(ClientData instanceData, Tcl_Interp *interp);
static int _c_DriverCloseProc(ClientData instanceData, Tcl_Interp *interp) {
	return _go_DriverCloseProc(instanceData, interp);
}

extern int _go_DriverInputProc(ClientData instanceData, char *buf, int bufSize, int *errorCodePtr);
static int _c_DriverInputProc(ClientData instanceData, char *buf, int bufSize, int *errorCodePtr) {
	return _go_DriverInputProc(instanceData, buf, bufSize, errorCodePtr);
}

static int _c_DriverOutputProc(ClientData instanceData, const char *buf, int toWrite, int *errorCodePtr) {
	*errorCodePtr = EINVAL;
	return -1;
}

extern int _go_DriverSeekProc(ClientData instanceData, long offset, int seekMode, int *errorCodePtr);
static int _c_DriverSeekProc(ClientData instanceData, long offset, int seekMode, int *errorCodePtr) {
	return _go_DriverSeekProc(instanceData, offset, seekMode, errorCodePtr);
}

static void _c_DriverWatchProc(ClientData instanceData, int mask) {
	return;
}

static int _c_DriverGetHandleProc(ClientData instanceData, int direction, ClientData *handlePtr) {
	return TCL_ERROR;
}

Tcl_ChannelType _c_ChannelType = {
  "govfs",
  NULL, // Tcl_ChannelTypeVersion
  _c_DriverCloseProc,
  _c_DriverInputProc,
  _c_DriverOutputProc,
  _c_DriverSeekProc,
  NULL, // Tcl_DriverSetOptionProc
  NULL, // Tcl_DriverGetOptionProc
  _c_DriverWatchProc,
  _c_DriverGetHandleProc
};

extern int _go_FSPathInFilesystemProc(Tcl_Obj *pathPtr);
static int _c_FSPathInFilesystemProc(Tcl_Obj *pathPtr, ClientData *clientDataPtr) {
	return _go_FSPathInFilesystemProc(pathPtr);
}

static Tcl_Obj *_c_FSFilesystemPathTypeProc(Tcl_Obj *pathPtr) {
	return NULL;
}

extern int _go_FSStatProc(Tcl_Obj *pathPtr, Tcl_StatBuf *statPtr);
static int _c_FSStatProc(Tcl_Obj *pathPtr, Tcl_StatBuf *statPtr) {
  memset(statPtr, 0, sizeof(Tcl_StatBuf));
	return _go_FSStatProc(pathPtr, statPtr);
}

extern int _go_FSAccessProc(Tcl_Obj *pathPtr, int mode);
static int _c_FSAccessProc(Tcl_Obj *pathPtr, int mode) {
	return _go_FSAccessProc(pathPtr, mode);
}

extern Tcl_Channel _go_FSOpenFileChannelProc(Tcl_Interp *interp, Tcl_Obj *pathPtr, int mode, int permissions);
static Tcl_Channel _c_FSOpenFileChannelProc(Tcl_Interp *interp, Tcl_Obj *pathPtr, int mode, int permissions) {
	return _go_FSOpenFileChannelProc(interp, pathPtr, mode, permissions);
}

extern int _go_FSMatchInDirectoryProc(Tcl_Interp *interp, Tcl_Obj *resultPtr, Tcl_Obj *pathPtr, char *pattern, Tcl_GlobTypeData *types);
static int _c_FSMatchInDirectoryProc(Tcl_Interp *interp, Tcl_Obj *resultPtr, Tcl_Obj *pathPtr, const char *pattern, Tcl_GlobTypeData *types) {
	return _go_FSMatchInDirectoryProc(interp, resultPtr, pathPtr, (char *)pattern, types);
}

// extern Tcl_Obj *_go_FSListVolumesProc();
// static Tcl_Obj *_c_FSListVolumesProc(void) {
// 	return _go_FSListVolumesProc();
// }

// extern char *const *_go_FSFileAttrStringsProc(Tcl_Obj *pathPtr, Tcl_Obj **objPtrRef);
// static const char *const *_c_FSFileAttrStringsProc(Tcl_Obj *pathPtr, Tcl_Obj **objPtrRef) {
// 	return _go_FSListVolumesProc(pathPtr, objPtrRef);
// }

// extern int _go_FSFileAttrsGetProc(Tcl_Interp *interp, int index, Tcl_Obj *pathPtr, Tcl_Obj **objPtrRef);
// static int _c_FSFileAttrsGetProc(Tcl_Interp *interp, int index, Tcl_Obj *pathPtr, Tcl_Obj **objPtrRef) {
// 	return _go_FSFileAttrsGetProc(interp, index, pathPtr, objPtrRef);
// }

// extern int _go_FSFileAttrsSetProc(Tcl_Interp *interp, int index, Tcl_Obj *pathPtr, Tcl_Obj *objPtr);
// static int _c_FSFileAttrsSetProc(Tcl_Interp *interp, int index, Tcl_Obj *pathPtr, Tcl_Obj *objPtr) {
// 	return _go_FSFileAttrsSetProc(interp, index, pathPtr, objPtr);
// }

static int _c_FSChdirProc(Tcl_Obj *pathPtr) {
	return TCL_OK;
}

Tcl_Filesystem _c_Filesystem = {
    "govfs",
    sizeof(Tcl_Filesystem),
    TCL_FILESYSTEM_VERSION_1,
    _c_FSPathInFilesystemProc,
    0, // _c_FSDupInternalRepProc,
    0, // _c_FSFreeInternalRepProc,
    0, // _c_FSInternalToNormalizedProc,
    0, // _c_FSCreateInternalRepProc,
    0, // _c_FSNormalizePathProc,
    _c_FSFilesystemPathTypeProc,
    0, // _c_FSFilesystemSeparatorProc,
    _c_FSStatProc,
    _c_FSAccessProc,
    _c_FSOpenFileChannelProc,
    _c_FSMatchInDirectoryProc,
    0, // _c_FSUtimeProc,
    0, // _c_FSLinkProc,
    0, // _c_FSListVolumesProc,
    0, // _c_FSFileAttrStringsProc,
    0, // _c_FSFileAttrsGetProc,
    0, // _c_FSFileAttrsSetProc,
    0, // _c_FSCreateDirectoryProc,
    0, // _c_FSRemoveDirectoryProc,
    0, // _c_FSDeleteFileProc,
    0, // _c_FSCopyFileProc,
    0, // _c_FSRenameFileProc,
    0, // _c_FSCopyDirectoryProc,
    0, // _c_FSLoadFileProc,
    0, // _c_FSUnloadFileProc,
    0, // _c_FSGetCwdProc,
    _c_FSChdirProc
};

*/
import "C"

var (
	tclVFSCallbacks     = &C._c_Filesystem
	tclChannelCallbacks = &C._c_ChannelType
)

func init() {
	C.Tcl_FSRegister(nil, tclVFSCallbacks)
}

func objToString(interp *C.Tcl_Interp, obj *C.Tcl_Obj) string {
	var n C.int
	out := C.Tcl_GetStringFromObj(obj, &n)
	return C.GoStringN(out, n)
}
