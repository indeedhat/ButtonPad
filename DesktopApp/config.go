package main

import (
    "io/ioutil"
    "path/filepath"

    "github.com/indeedhat/vkb/layout"
    "./com"
    "./env"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Env    map[string]string `yaml:""`
    Layers []*Layer          `yaml:",flow"`
    Serial *com.SerialConfig `yaml:""`
}


// create a new blank config
func NewConfig() *Config {
    c := &Config{}
    c.Env = make(map[string]string)

    // setup environment defaults
    c.Env[env.LAYOUT] = env.DEFAULT_LAYOUT
    c.Env[env.LAYER]  = env.DEFAULT_LAYER

    // setup empty layers
    c.Layers = make([]*Layer, env.PAD_LAYERS)
    for i := 0; i < env.PAD_LAYERS; i++ {
        c.Layers[i], _ = NewLayer(env.PAD_BUTTONS)
    }

    // default serial config
    c.Serial = com.NewSerialConf()

    return c
}


// load the appropriate keyboard layout from indeedhat/vkb/layout
// this can be extended in the future to load a layout from anywhere
// it would be interesting to see about loading from a text file to allow people to easily create their own layouts
func (c *Config) LoadKeyboardLayout() *layout.Layout {
    // i realise just how pointless this seems but it is here for future use when i have added more layouts
    switch c.Env[env.LAYOUT] {
    default:
        l := layout.QwertyUs()
        return &l
    }
}


// attempt to load the config from file
func (a *App) loadConfig() (err error) {
    var path string
    var data []byte

    if path, err = filepath.Abs(env.CONFIG_PATH); nil != err {
        return
    }

    if data, err = ioutil.ReadFile(path); nil != err {
        return
    }

    c := &Config{}
    if err = yaml.Unmarshal(data, c); nil == err {
        a.config = c
    }

    return
}


// attempt to save the config to file
// this will overwrite any existing config should one already exist
func (a *App) saveConfig() (err error) {
    var data []byte
    var path string

    if path, err = filepath.Abs(env.CONFIG_PATH); nil != err {
        return
    }

    if data, err = yaml.Marshal(a.config); nil != err {
        return
    }

    err = ioutil.WriteFile(path, data, 0755)
    return
}