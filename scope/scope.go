package main

import "fmt"

var x string = "Hello World"
const y string = "Can't change this"

func main() {
  fmt.Println(x)
  f()
// The following line won't work
//  y = "New Value"
}

func f() {
  fmt.Println(x)
}
