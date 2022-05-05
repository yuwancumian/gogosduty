package main

import (
  "fmt"
  "strconv"
)

func Cast1() {
  str := "12"
  a1, err := strconv.Atoi(str)
  a2 := int8(a1)
  fmt.Println(a1)
  fmt.Println(a2)
  fmt.Println(err)

}
