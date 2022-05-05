package main

import (
	"fmt"
	"strings"
)

func TestString1() {
	var a = "ahaha"
	fmt.Printf("aa %s \n", a)
	fmt.Printf("%s \n", strings.Title(a))
	var b = "bbb"
	fmt.Printf("%s is some", b)
}
