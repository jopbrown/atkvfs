

This is a addon package for [atk](https://github.com/visualfc/atk) that wrapping [TCL FileSystem API](https://www.tcl-lang.org/man/tcl8.6/TclLib/FileSystem.htm) to mount `http.FileSystem` as virtual filesystem.

## Features:
* Embed TCL/TK stdlib to build a standalone GUI app without installing `libtcl/libtk`.
* Mount/Unmount `http.FileSystem` as virtual filesystem.

## Status
* `Linux` is available via `cgo`.
    * To build a fully standalone app, you need to [install TCL/TK from source](https://www.tcl.tk/doc/howto/compile.html) staticlly.
        ```bash
        # build libtcl staticlly
        cd tcl8.6.10/unix && \
        ./configure --disable-shared && \
        make && \
        make install

        # build libtk staticlly
        cd tk8.6.10/unix && \
        ./configure --disable-shared && \
        make && \
        make install

        # soft link lib path
        cd /usr/local/lib
        ln -s libtcl8.6.a libtcl.a
        ln -s libtk8.6.a libtk.a
        ```

* `Windows` is still planning.
* `Mac OS` has no plan because it can package libs into an app.

## Usage

### Embed TCL/TK stdlib
import `github.com/jopbrown/atkvfs/stdlib/v8.6.10/embed` as the blank identifier.

Sample:

```go
package main

import (
	_ "github.com/jopbrown/atkvfs/stdlib/v8.6.10/embed"
	"github.com/visualfc/atk/tk"
)

type Window struct {
	*tk.Window
}

func NewWindow() *Window {
	mw := &Window{tk.RootWindow()}
	lbl := tk.NewLabel(mw, "Hello World!")
	btn := tk.NewButton(mw, "Quit")
	btn.OnCommand(func() {
		tk.Quit()
	})
	tk.NewVPackLayout(mw).AddWidgets(lbl, tk.NewLayoutSpacer(mw, 0, true), btn)
	mw.ResizeN(300, 200)
	return mw
}

func main() {
	tk.MainLoop(func() {
		mw := NewWindow()
		mw.SetTitle("Hello World")
		mw.Center()
		mw.ShowNormal()
	})
}
```

### Mount/Unmount virtual filesystem.
```go
var fs http.FileSystem = assert("/myvfs")
atkvfs.Mount("/myvfs", fs)
```
