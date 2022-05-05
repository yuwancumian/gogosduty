package main

import (
	"bytes"
	"fmt"
)

func Bytestest() {
	//var b []byte
	var b bytes.Buffer
	(&b).Write([]byte("hello world"))
	fmt.Println(b.String())
}
