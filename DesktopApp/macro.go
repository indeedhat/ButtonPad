package main

import (
    "errors"
    "github.com/indeedhat/vkb"
    "./util"
    "./env"
)




type Macro struct {
    Label   string    `yaml:""`
    Type    env.MacroType `yaml:""`

    // default payload
    Payload []byte    `yaml:""`

    // platform specific payloads
    Windows []byte    `yaml:""`
    Linux   []byte    `yaml:""`
    Darwin  []byte    `yaml:""`
}


// initialize a new macro
// setting platform specific payloads should be done outside of this constructor
func NewMacro(label string, payload []byte, macroType env.MacroType) Macro {
    return Macro{
        Label:   label,
        Payload: payload,
        Type:    macroType,
    }
}


// check if the macro has a payload and therefore can be triggered
func (mac *Macro) IsSet() bool {
    return 0 < len(mac.Payload)
}

// Run the macro
func (mac *Macro) Execute(kb vkb.VirtKB) (err error) {
    // clear the virtual keyboards stroke queue before running macro
    kb.Reset()

    switch mac.Type {
    case env.MACRO_TEXT:
        err = mac.executeText(kb)
    case env.MACRO_OPEN:
        err = mac.executeOpen(kb)
    case env.MACRO_SEQUENCE:
        err = mac.executeSequence(kb)
    case env.MACRO_CODE:
        err = mac.executeCode(kb)
    }

    return
}


// execute the macro based on the payload being a string of text
func (mac *Macro) executeText(kb vkb.VirtKB) error {
    // convert payload to a string for parsing
    text := string(mac.Payload)

    // parse text and execute keystrokes
    if err := kb.ParseWithAssignedLayout(text, false); nil != err {
        return err
    }

    return kb.Type()
}


// execute the macro and open the application specified in the payload
func (mac *Macro) executeOpen(kb vkb.VirtKB) error {
    application := string(mac.Payload)

    return util.OpenExecutable(application)
}


// execute the macro based on the payload being an array of keystrokes
func (mac *Macro) executeSequence(kb vkb.VirtKB) error {
    keystrokes := util.PayloadToIntArray(mac.Payload)

    // add payload to virtual keyboard
    kb.AddStrokes(keystrokes)

    return kb.Type()
}


// really not sure if im going to bother with this one or not
// would be nice to have but probably not worth the effort
func (mac *Macro) executeCode(kb vkb.VirtKB) error {
    return errors.New("not yet implemented")
}