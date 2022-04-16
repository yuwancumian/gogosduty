package main

import "fmt"

func testArr1() {
	var arr1 [2]string
	var m1 = map[string]string{}
	m1["a"] = "aaaa"
	m1["b"] = "bbbb"

	key := "good"
	value := m1[key]
	fmt.Printf("%v", arr1)
	fmt.Printf("%q means %#v\n", key, value)
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
