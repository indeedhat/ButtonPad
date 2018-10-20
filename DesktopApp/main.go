package main

import (
    "errors"

    "github.com/indeedhat/vkb"
)

type App struct {
    config    *Config
    Active    *Layer
    VirtualKB *vkb.VirtKB
}


// initialize a new application
func NewApp() (*App, error) {
    // init app
    var err error
    app := &App{}

    // load/create config
    if nil != app.loadConfig(){
        app.config = NewConfig()
        app.saveConfig()
    }

    // initialize virtual keyboard
    if app.VirtualKB, err = vkb.NewVKB(); nil != err {
        return nil, err
    }

    // select the keyboard layout
    // TODO: this will need to be able to be set dynamically but for now ill just use my layout
    qus := app.config.LoadKeyboardLayout()
    app.VirtualKB.AssignLayout(qus)

    // select the active layer
    return app, app.SetActiveLayer(0)
}


// set a layer as active
// returns an error if the layer doesn't exist
func (app *App) SetActiveLayer(i int) error {
    if i < 0 || i > len(app.config.Layers) {
        return errors.New("layer out of range")
    }

    app.Active = app.config.Layers[i]

    return nil
}