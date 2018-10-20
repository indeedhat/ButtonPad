package main

import (
    "errors"

    "github.com/indeedhat/vkb"
    "github.com/indeedhat/vkb/layout"
)

type App struct {
    Layers    []*Layer
    Active    *Layer
    VirtualKB *vkb.VirtKB
}


// initialize a new application
func NewApp(layers, macros int) (*App, error) {
    if 1 > layers || 1 > macros {
        return nil, errors.New("must have at least one layer and macro")
    }

    // init app
    var err error
    app := &App{}

    // setup layers
    app.Layers = make([]*Layer, layers)
    for i := 0; i < layers; i++ {
        if app.Layers[i], err = NewLayer(macros); nil != err {
            return nil, err
        }
    }

    // initialize virtual keyboard
    if app.VirtualKB, err = vkb.NewVKB(); nil != err {
        return nil, err
    }

    // select the keyboard layout
    // TODO: this will need to be able to be set dynamically but for now ill just use my layout
    qus := layout.QwertyUs()
    app.VirtualKB.AssignLayout(&qus)

    // select the active layer
    return app, app.SetActiveLayer(0)
}


// set a layer as active
// returns an error if the layer doesn't exist
func (app *App) SetActiveLayer(i int) error {
    if i < 0 || i > len(app.Layers) {
        return errors.New("layer out of range")
    }

    app.Active = app.Layers[i]

    return nil
}