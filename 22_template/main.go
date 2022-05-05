package main

import (
  "fmt"
  "text/template"
  "os"
)

type Person struct {
  Name string
  Age int8
}

func main() {
  royi := Person{"royi",20}
  tmpl := `
    My name is {{.Name}},
    my age is {{.Age}}
  `
  t, err := template.New("tmpl.txt").Parse(tmpl)
  if err != nil {
    fmt.Println(err)
  }
  err = t.Execute(os.Stdout, royi)
  if err != nil { panic(err) }
}


