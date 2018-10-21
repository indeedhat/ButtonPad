package com

import (
    "bufio"
    "io"

    "github.com/jacobsa/go-serial/serial"
)

type Adapter struct {
    readBuffer  []byte
    writeBuffer []byte
    client      io.ReadWriteCloser
    reader      *bufio.Reader
}


// create a new adapter instance
func NewAdapter(conf *SerialConfig) (adapter *Adapter, err error) {
    opts := serial.OpenOptions{
        PortName:        conf.Port,
        BaudRate:        conf.BaudRate,
        DataBits:        conf.DataBits,
        StopBits:        conf.StopBits,
        MinimumReadSize: conf.MinimumReadSize,
    }

    adapter = &Adapter{}
    if adapter.client, err = serial.Open(opts); nil != err {
        adapter = nil
        return
    }

    adapter.reader = bufio.NewReader(adapter.client)
    return
}


// send a frame to the arduino
func (a *Adapter) WriteFrame(frame Frame) (err error) {
    a.clearWriteBuffer()
    frame.Marshal(&a.writeBuffer)

    _, err = a.client.Write(a.writeBuffer)
    return
}

// read a frame from the serial connection
func (a *Adapter) ReadFrame() (frame *Frame, err error) {
    a.clearReadBuffer()
    var data []byte
    if a.readBuffer, err = a.reader.ReadBytes(byte(nil)); nil != err {
        return
    }

    frame = &Frame{}
    if err = frame.Unmarshal(data); nil != err {
        frame = nil
    }

    return
}


// clear the read buffer ready for a new read cycle
func (a *Adapter) clearReadBuffer() {
    a.readBuffer = []byte{}
}


// clear the write buffer ready for a new write cycle
func (a *Adapter) clearWriteBuffer() {
    a.writeBuffer = []byte{}
}