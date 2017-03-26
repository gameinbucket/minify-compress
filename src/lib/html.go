package lib

import (
    "bytes"
)

func MiniHTML(src []byte) ([]byte) {
    size := len(src)
    dest := make([]byte, size)
    
    j := 0
    for i := 0; i < size; i++ {
        if src[i] == '\n' {
            
        } else {
            dest[j] = src[i]
            j++
        }
    }
    
    return bytes.TrimRight(dest, "\x00")
}