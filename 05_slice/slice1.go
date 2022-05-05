package main
import (
  "fmt"
)

func TestSlice1() {
  sl := []int8{
    1,2,3,
  }  
  for i, num := range sl {
    fmt.Println(i, num)
  }
}
