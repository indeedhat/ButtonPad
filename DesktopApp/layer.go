package main

import (
    "errors"
)

type Layer struct {
    Macros []*Macro
}


// create a new layer with the desired number macros
func NewLayer(size int) (*Layer, error) {
    if 1 > size {
        return nil, errors.New("layer must have at least one macro")
    }

    layer := &Layer{}
    layer.Macros = make([]*Macro, size)


    return layer, nil
}


// get a specific macro from the layer
func (l *Layer) GetMacro(i int) (*Macro, bool) {
    if 0 > i || len(l.Macros) < i + 1 {
        return nil, false
    }

    return l.Macros[i], true
}


// replace the macro on the layer by its index
func (l *Layer) ReplaceMacro(macro *Macro, i int) error {
    if 0 > i || len(l.Macros) < i + 1 {
        return errors.New("macro index out of range")
    }

    l.Macros[i] = macro
    return nil
}