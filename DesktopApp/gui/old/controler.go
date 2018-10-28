package old

import (
    "fmt"
    "github.com/gotk3/gotk3/gtk"
)


const (
    APP_TITLE = "Button Pad"
)


// keep track of windo id's
var id = 0

type UI struct {
    OnConfirm     func(confirm bool)
    OnChangeMacro func (layer, macro int, typ byte, label string, payload []byte)

    windows       map[int]*gtk.Window
}


// create a new app ui instance
func NewUiController(
    title   string,
    confirm func(confirm bool),
    change  func (layer, macro int, typ byte, label string, payload []byte),
) *UI {
    ui := &UI{
        OnConfirm:     confirm,
        OnChangeMacro: change,
    }

    gtk.Init(nil)

    win, err := ui.createWindow(APP_TITLE, gtk.WINDOW_TOPLEVEL, func() {
        gtk.MainQuit()
    })
    if nil != err {
        return nil
    }

    win.ShowAll()
    gtk.Main()

    return ui
}


// create a new window and setup its default behaviour such as what to do when closed
func (ui *UI) NewWindow(title string, content gtk.IWidget, width, height int) (win *gtk.Window, err error) {
    if win, err = ui.createWindow(title, gtk.WINDOW_POPUP, func() {fmt.Printf("closed window")}); nil != err {
        return
    }

    win.SetDefaultSize(width, height)
    win.Add(content)
    win.ShowAll()

    return
}


func (ui *UI) createWindow(title string, windowType gtk.WindowType, onClose func()) (win *gtk.Window, err error) {

    if win, err = gtk.WindowNew(windowType); nil != err {
        return
    }

    win.SetTitle(title)
    win.Connect("destroy", onClose)

    // set the default size (this will probably need to be set outside of this method but meh)
    win.SetDefaultSize(400, 600)

    // add to the controller
    winid := id
    id++
    ui.windows[winid] = win
    return
}