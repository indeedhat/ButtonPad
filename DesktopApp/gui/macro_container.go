package gui

import (
	"fmt"
	"github.com/go-errors/errors"
	"github.com/gotk3/gotk3/gtk"
	"github.com/indeedhat/ButtonPad/DesktopApp/env"
	"strconv"
)

var currentForm *gtk.Grid

func MacroUpdateContainer() (container *gtk.Grid, err error) {
	container, err = gtk.GridNew()
	if nil != err {
		return
	}

	container.SetOrientation(gtk.ORIENTATION_VERTICAL)

	macroTypes, err := gtk.ComboBoxTextNew()
	if nil != err {
		return
	}

	macroTypes.Append(strconv.Itoa(int(env.MACRO_TEXT)), "Text")
	macroTypes.Append(strconv.Itoa(int(env.MACRO_SEQUENCE)), "Sequence")
	macroTypes.Append(strconv.Itoa(int(env.MACRO_OPEN)), "Run Program")
	macroTypes.Append(strconv.Itoa(int(env.MACRO_CODE)), "Run Code")
	macroTypes.Connect("changed", func() {
		fmt.Println(macroTypes.GetActiveID())
		fmt.Println(strconv.Itoa(int(env.MACRO_TEXT)))
		fmt.Println(strconv.Itoa(int(env.MACRO_SEQUENCE)))
		fmt.Println(strconv.Itoa(int(env.MACRO_OPEN)))
		fmt.Println(strconv.Itoa(int(env.MACRO_CODE)))
		var newForm *gtk.Grid
		switch macroTypes.GetActiveID() {
		case strconv.Itoa(int(env.MACRO_TEXT)):
			newForm, err = MacroTextForm()
		case strconv.Itoa(int(env.MACRO_SEQUENCE)):
			newForm, err = MacroSequenceForm()
		case strconv.Itoa(int(env.MACRO_OPEN)):
			newForm, err = MacroAppForm()
		case strconv.Itoa(int(env.MACRO_CODE)):
			newForm, err = MacroCodeForm()
		}

		if nil != err {
			fmt.Printf("Error: %s\n", err)
		}

		if nil != currentForm {
			container.Remove(currentForm)
		}

		fmt.Println(newForm)
		if nil != newForm {
			fmt.Println("doing the thing")
			currentForm = newForm
			container.Add(currentForm)
			container.ShowAll()
		}
	})

	macroLabel, err := gtk.EntryNew()
	if nil != err {
		return
	}

	container.Add(macroTypes)
	container.Add(macroLabel)

	macroTypes.SetActiveID(strconv.Itoa(int(env.MACRO_TEXT)))

	return
}



func MacroTextForm() (form *gtk.Grid, err error) {
	form, err = gtk.GridNew()
	if nil != err {
		return
	}

	scrollWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if nil != err {
		return
	}
	scrollWindow.SetSizeRequest(300, 300)

	macroText, err := gtk.TextViewNew()
	if nil != err {
		return
	}
	macroText.SetSizeRequest(300, 300)

	buf, err := macroText.GetBuffer()
	if nil != err {
		return
	}

	buf.SetText("Example Text")

	scrollWindow.Add(macroText)
	form.Add(scrollWindow)
	return
}


func MacroCodeForm() (form *gtk.Grid, err error) {
	form, err = gtk.GridNew()
	if nil != err {
		return
	}

	scrollWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if nil != err {
		return
	}
	scrollWindow.SetSizeRequest(300, 300)

	macroText, err := gtk.TextViewNew()
	if nil != err {
		return
	}
	macroText.SetSizeRequest(300, 300)

	buf, err := macroText.GetBuffer()
	if nil != err {
		return
	}

	buf.SetText("func() {\n\n}")

	scrollWindow.Add(macroText)
	form.Add(scrollWindow)
	return
}


func MacroAppForm() (form *gtk.Grid, err error) {
	form, err = gtk.GridNew()
	if nil != err {
		return
	}
	form.SetOrientation(gtk.ORIENTATION_HORIZONTAL)

	macroFile, err := gtk.EntryNew()
	if nil != err {
		return
	}

	macroChooser, err := gtk.FileChooserButtonNew("Browse", gtk.FILE_CHOOSER_ACTION_OPEN)
	if nil != err {
		return
	}

	macroChooser.Connect("update_preview", func() {
		macroFile.SetText(macroChooser.GetPreviewFilename())
	})


	form.Add(macroFile)
	form.Add(macroChooser)
	return
}


func MacroSequenceForm() (form *gtk.Grid, err error) {
	// TODO: this needs implementing
	err = errors.New("Not yet implemented")
	return
}
