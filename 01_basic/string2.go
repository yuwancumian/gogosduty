package main
import (
  "fmt"
)

func TestString2() {
  var str string = "hello world"
  var b = []byte(str)
  //fmt.Printf("%0x",b)
  fmt.Println(string(b))
  for _, v := range b {
    fmt.Printf("%s\n", string(v))
  }
}
