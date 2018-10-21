package util

import (
    "bytes"
)


// converts an []int to a []byte
func IntArrayToPayload(data []int) []byte {
    var buf bytes.Buffer

    for _, val := range data {
        buf.WriteByte(byte(val))
    }

    return buf.Bytes()
}


// converts a []byte to an []int
// this is designed specifically for this application and is sufficient for the virtual keycodes
// used by this package
// it will not work for any situation where each int cannot be represented by a single byte
func PayloadToIntArray(payload []byte) []int {
    out := make([]int, len(payload))
    for i, val := range payload {
        out[i] = int(val)
    }

    return out
}