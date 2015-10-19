package main

import(
  "fmt"
  "time"
  "math/rand"
)

func small(i int){
  amt := time.Duration(rand.Intn(2))
  time.Sleep(time.Millisecond * amt * 100)
  fmt.Println("Small #", i)
}

func medium(i int){
  amt := time.Duration(rand.Intn(10))
  time.Sleep(time.Millisecond * amt * 100)
  fmt.Println("Medium #", i)
}

func large(i int){
  amt := time.Duration(rand.Intn(50))
  time.Sleep(time.Millisecond * amt * 100)
  fmt.Println("Large #", i)
}

func processPhoto(i int) {
  fmt.Println("Queuing Photo #", i)

  go func() {
    large(i)
    medium(i)
    small(i)
  }()

  /*go large(i)*/
  /*go medium(i)*/
  /*go small(i)*/

}

func main() {

  for i := 1; i < 100; i++ {
    processPhoto(i)
    time.Sleep(time.Millisecond * 100)
  }

  var input string
  fmt.Scanln(&input)
}
