package main
import (
  "fmt"
)

func Struct1 () {
  type Person struct  {
    Name string
    Age int8
  }
  dapang := Person{"dapang", 20}
  dapang.Age = 21
  fmt.Println(dapang)
}
