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

}