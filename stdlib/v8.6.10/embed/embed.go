package embed

import (
	"github.com/jopbrown/atkvfs"
	"github.com/visualfc/atk/tk"
)

const (
	_VFS_PREFIX           = "govfs/stdlib"
	EMBED_TCL_LIBARY_PATH = _VFS_PREFIX + "/tcl8.6"
	EMBED_TK_LIBARY_PATH  = _VFS_PREFIX + "/tk8.6"
)

func init() {
	atkvfs.Mount(_VFS_PREFIX, AssetFile())
	tk.InitEx(true, EMBED_TCL_LIBARY_PATH, EMBED_TK_LIBARY_PATH)
}
