package main

type MacroUpdate struct {
    Type    int
    Layer   byte
    MacroId byte
    Label   string
    Payload []byte
}