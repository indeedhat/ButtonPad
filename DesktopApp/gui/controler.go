package main

import (
    "github.com/fyne-io/fyne"
    "github.com/fyne-io/fyne/desktop"
)

// keep track of windo id's
var id = 0

type UI struct {
    OnConfirm     func(confirm bool)
    OnChangeMacro func (layer, macro int, typ byte, label string, payload []byte)

    app           fyne.App
    windows       map[int]*fyne.Window
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

    app := desktop.NewApp()
    ui.app     = app
    ui.windows = make(map[int]*fyne.Window)

    return ui
}


// create a new window and setup its default behaviour such as what to do when closed
func (ui *UI) newWindow(title string, content fyne.CanvasObject) {
    // need to create an index so the window can be removed later
    winid := id
    id++

    window := ui.app.NewWindow(title)
    ui.windows[winid] = &window

    // remove the window from the ui when it is closed
    window.SetOnClosed(func() {
        delete(ui.windows, winid)
    })

    // show the window with its content
    window.SetContent(content)
    window.Show()
}