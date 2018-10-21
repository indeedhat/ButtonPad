package com


type SerialConfig struct {
    Port            string
    BaudRate        uint   `yaml:"baud"`
    DataBits        uint   `yaml:"data-bits"`
    StopBits        uint   `yaml:"stop-bits"`
    MinimumReadSize uint   `yaml:"min-read"`
}


// sets the default vals for serial config
func NewSerialConf() *SerialConfig {
    // TODO: this will need different vals for different platforms
    // TODO: run a test for this
    return &SerialConfig{
        "/dev/ttyACM0",
        19200,
        8,
        1,
        4,
    }
}