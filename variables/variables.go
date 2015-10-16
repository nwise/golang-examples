package main

import "fmt"

func main() {
  var x string
  x = "This seems like too much work"
  fmt.Println(x)

  var y string = "This is a little better"
  fmt.Println(y)

  z := "Let Golang do the boring work"
  fmt.Println(z)


  var(
    a = 5
    b = 7
    c = 9
  )
  fmt.Println(a + b + c)
}
