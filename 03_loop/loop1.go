package main
import (
  "fmt"
)

func Loop1() {
    i := 5
    for {
      i = i -1
      fmt.Println(i)
      if i< 0 {
        break
      }
    }   
}


