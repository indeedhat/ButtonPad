package com

import (
    "errors"
    "strconv"
    "unsafe"
)

type FrameType byte

const (
    FRAME_UPDATE FrameType = iota
    FRAME_TRIGGER
)


type Frame struct {
    Len     uint16
    Type    FrameType
    Payload []byte
}


// create new comunication frame
func NewFrame(typ FrameType, payload []byte) *Frame {
    return &Frame{
        uint16(len(payload)),
        typ,
        payload,
    }
}


// unmarshal the byte stream into a frame
func (f *Frame) Unmarshal(buffer []byte) (err error) {
    var l int
    if l, err = strconv.Atoi(string(buffer[0:2])); nil != err {
        return
    }

    if byte(nil) != buffer[len(buffer) -1] {
        err = errors.New("payload was not terminated, discarding frame")
    }

    f.Len     = uint16(l)
    f.Type    = FrameType(buffer[2])
    f.Payload = buffer[3:len(buffer) -5]

    return
}


// marshal a frame into a byte buffer
func (f *Frame) Marshal(buffer *[]byte) {
    var tmp []byte

    // add length to payload
    l := (*[2]byte)(unsafe.Pointer(&f.Len))[:]
    tmp = append(tmp, l...)

    // add type to payload
    tmp = append(tmp, byte(f.Type))

    // add the payload
    tmp = append(tmp, f.Payload...)

    // terminate
    tmp = append(tmp, byte(nil))

    buffer = &tmp
}