package main

import (
    "fmt"

    "github.com/fyne-io/fyne"
    "github.com/fyne-io/fyne/layout"
    "github.com/fyne-io/fyne/widget"
)

func LayerTemplate(callback func(macro int)) *widget.Group {
   return widget.NewGroup("Macros", []fyne.CanvasObject{
       fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
           &widget.Button{Text: "Macro1", OnTapped: func() { fmt.Println("Macro 1") }},
           &widget.Button{Text: "Macro2", OnTapped: func() { fmt.Println("Macro 2") }},
           &widget.Button{Text: "Macro3", OnTapped: func() { fmt.Println("Macro 3") }},
           &widget.Button{Text: "Macro4", OnTapped: func() { fmt.Println("Macro 4") }},
       }...),
       fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
           &widget.Button{Text: "Macro5", OnTapped: func() { fmt.Println("Macro 5") }},
           &widget.Button{Text: "Macro6", OnTapped: func() { fmt.Println("Macro 6") }},
           &widget.Button{Text: "Macro7", OnTapped: func() { fmt.Println("Macro 7") }},
           &widget.Button{Text: "Macro8", OnTapped: func() { fmt.Println("Macro 8") }},
       }...),
       fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
           &widget.Button{Text: "Macro9", OnTapped: func() { fmt.Println("Macro 9") }},
           &widget.Button{Text: "Macro10", OnTapped: func() { fmt.Println("Macro 10") }},
           &widget.Button{Text: "Macro11", OnTapped: func() { fmt.Println("Macro 11") }},
           &widget.Button{Text: "Macro12", OnTapped: func() { fmt.Println("Macro 12") }},
       }...),
       fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
           &widget.Button{Text: "Macro13", OnTapped: func() { fmt.Println("Macro 13") }},
           &widget.Button{Text: "Macro14", OnTapped: func() { fmt.Println("Macro 14") }},
           &widget.Button{Text: "Macro15", OnTapped: func() { fmt.Println("Macro 15") }},
           &widget.Button{Text: "Macro16", OnTapped: func() { fmt.Println("Macro 16") }},
       }...),
   }...)
}


func LayerSelector(callback func(layer int)) *widget.Group {
    return widget.NewGroup("Layers", []fyne.CanvasObject{
        fyne.NewContainerWithLayout(layout.NewGridLayout(4), []fyne.CanvasObject{
            &widget.Button{
                Text: "Layer1", OnTapped: func() { fmt.Println("Layer 1") }, Style: widget.PrimaryButton,
            },
            &widget.Button{
                Text: "Layer2", OnTapped: func() { fmt.Println("Layer 2") }, Style: widget.PrimaryButton,
            },
            &widget.Button{
                Text: "Layer3", OnTapped: func() { fmt.Println("Layer 3") }, Style: widget.PrimaryButton,
            },
            &widget.Button{
                Text: "Layer4", OnTapped: func() { fmt.Println("Layer 4") }, Style: widget.PrimaryButton,
            },
        }...),
    }...)
}