package main

import "fmt"

const test3 string = "123456"

func testArray2() {
	var arr1 [3]int
	var arr2 = [2]string{"aa", "bb"}

	fmt.Println(arr1)
	fmt.Printf("books: #{arr2}")
	fmt.Printf("books  : %q\n", arr2)
}
