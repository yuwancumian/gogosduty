package main
import (
  "fmt"
  "flag"
)

func main() {
  fmt.Println("hello")
  method := flag.String("i", "123", "方法")
  fmt.Println("method is ",*method)
}
