package main

import (
  "fmt"
  "time"
  "math/rand"
)

func f(n int) {
  for i :=0; i < 10; i++ {
    fmt.Println("Routine #", n, ":", i)
    delay := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * delay)
  }
}

// Create ten goroutines that count to 10
func main() {
  for i :=0; i < 10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
