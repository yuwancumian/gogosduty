package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//dir, _ := os.Getwd()
	fmt.Printf("1111")
	//fmt.Printf("2222", dir)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

}
