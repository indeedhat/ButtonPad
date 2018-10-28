package gui

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
)


func LayerTemplate(li int, callback func(button *gtk.Button)) (grid *gtk.Grid, err error) {

	if grid, err = gtk.GridNew(); nil != err {
		return
	}

	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
	grid.SetHAlign(gtk.ALIGN_CENTER)


	for i := 1; i <= 4; i++ {
		var row *gtk.Grid
		if row, err = gtk.GridNew(); nil != err {
			return
		}

		row.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
		row.SetHAlign(gtk.ALIGN_CENTER)

		for j := 1; j <= 4; j++ {
			var btn *gtk.Button
			btn, err = gtk.ButtonNewWithLabel(fmt.Sprintf("L %v M %v",li, (i -1) *4 + j))
			if nil != err {
				return
			}

			btn.Connect("clicked", callback)
			btn.SetMarginTop(10)
			btn.SetMarginStart(5)
			btn.SetMarginEnd(5)



			row.Add(btn)
		}

		grid.Add(row)
	}

	return

	//return widget.NewGroup("Macros", []fyne.CanvasObject{
	//	fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
	//		&widget.Button{Text: "Macro1", OnTapped: func() { fmt.Println("Macro 1") }},
	//		&widget.Button{Text: "Macro2", OnTapped: func() { fmt.Println("Macro 2") }},
	//		&widget.Button{Text: "Macro3", OnTapped: func() { fmt.Println("Macro 3") }},
	//		&widget.Button{Text: "Macro4", OnTapped: func() { fmt.Println("Macro 4") }},
	//	}...),
	//	fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
	//		&widget.Button{Text: "Macro5", OnTapped: func() { fmt.Println("Macro 5") }},
	//		&widget.Button{Text: "Macro6", OnTapped: func() { fmt.Println("Macro 6") }},
	//		&widget.Button{Text: "Macro7", OnTapped: func() { fmt.Println("Macro 7") }},
	//		&widget.Button{Text: "Macro8", OnTapped: func() { fmt.Println("Macro 8") }},
	//	}...),
	//	fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
	//		&widget.Button{Text: "Macro9", OnTapped: func() { fmt.Println("Macro 9") }},
	//		&widget.Button{Text: "Macro10", OnTapped: func() { fmt.Println("Macro 10") }},
	//		&widget.Button{Text: "Macro11", OnTapped: func() { fmt.Println("Macro 11") }},
	//		&widget.Button{Text: "Macro12", OnTapped: func() { fmt.Println("Macro 12") }},
	//	}...),
	//	fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
	//		&widget.Button{Text: "Macro13", OnTapped: func() { fmt.Println("Macro 13") }},
	//		&widget.Button{Text: "Macro14", OnTapped: func() { fmt.Println("Macro 14") }},
	//		&widget.Button{Text: "Macro15", OnTapped: func() { fmt.Println("Macro 15") }},
	//		&widget.Button{Text: "Macro16", OnTapped: func() { fmt.Println("Macro 16") }},
	//	}...),
	//}...)
}


func LayerSelector(callback func(button *gtk.Button)) (grid *gtk.Grid, err error) {
	if grid, err = gtk.GridNew(); nil != err {
		return
	}

	grid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)

	for i := 1; i <= 4; i++ {
		var btn *gtk.Button
		btn, err = gtk.ButtonNewWithLabel(fmt.Sprintf("Layer %d", i))
		if nil != err {
			return
		}

		btn.Connect("clicked", callback)

		grid.Add(btn)
	}

	return


	//return widget.NewGroup("Layers", []fyne.CanvasObject{
	//	fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
	//		&widget.Button{
	//			Text: "Layer1", OnTapped: func() { fmt.Println("Layer 1") }, Style: widget.PrimaryButton,
	//		},
	//		&widget.Button{
	//			Text: "Layer2", OnTapped: func() { fmt.Println("Layer 2") }, Style: widget.PrimaryButton,
	//		},
	//		&widget.Button{
	//			Text: "Layer3", OnTapped: func() { fmt.Println("Layer 3") }, Style: widget.PrimaryButton,
	//		},
	//		&widget.Button{
	//			Text: "Layer4", OnTapped: func() { fmt.Println("Layer 4") }, Style: widget.PrimaryButton,
	//		},
	//	}...),
	//}...)
}