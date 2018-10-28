package gui


import (
    "fmt"
    "github.com/gotk3/gotk3/gtk"
    "strconv"
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
    confirm func(confirm bool),
    change  func (layer, macro int, typ byte, label string, payload []byte),
) *UI {
    ui := &UI{
        OnConfirm:     confirm,
        OnChangeMacro: change,
    }
    ui.windows = make(map[int]*gtk.Window)

    gtk.Init(nil)

    win, err := ui.createWindow(APP_TITLE, gtk.WINDOW_TOPLEVEL, func() {
        gtk.MainQuit()
    })
    if nil != err {
        return nil
    }

    layout, err := gtk.GridNew()
    if nil != err {
        panic(err.Error())
    }
    layout.SetOrientation(gtk.ORIENTATION_VERTICAL)

    macros := make(map[int]*gtk.Grid)
    activeLayer := 1
    for i := 1; i <= 4; i++ {
        macros[i], _ = LayerTemplate(i, func(button *gtk.Button) {fmt.Println(button.GetLabel())})

    }

    // add the default content
    layer, err := LayerSelector(func (button *gtk.Button) {
        layout.Remove(macros[activeLayer])
        lab, _ := button.GetLabel()
        activeLayer, _ = strconv.Atoi(string(lab[len(lab)-1]))


        fmt.Println("adding macros", activeLayer, macros[activeLayer])
        layout.Add(macros[activeLayer])
        layout.ShowAll()
    })
    if nil != err {
        panic(err.Error())
    }
    layout.Add(layer)

    // add macros
    layout.Add(macros[activeLayer])

    win.Add(layout)

    // display the window
    win.ShowAll()
    gtk.Main()

    return ui
}


// start the application
func (ui *UI) Start() {
    ui.windows[0].ShowAll()
    gtk.Main()
}


// expose createWindow to the rest of the app in a more useful format
func (ui *UI) NewWindow(title string, content gtk.IWidget, width, height int) (win *gtk.Window, err error) {
    if win, err = ui.createWindow(title, gtk.WINDOW_POPUP, func() {fmt.Printf("closed window")}); nil != err {
        return
    }

    win.SetDefaultSize(width, height)
    win.Add(content)
    win.ShowAll()

    return
}


// background method for doing the actual dirty work of making the window
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