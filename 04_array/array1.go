package main

import "fmt"

func testArray1() {
	//var names = []int8{1, 2, 3}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
