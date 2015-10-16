package main

import "fmt"

func main() {
  firstSet := []float64{ 98, 93, 77, 82, 83 }
  secondSet := []float64{ 78, 73, 57, 62, 53 }

  fmt.Println(average(firstSet))
  fmt.Println(average(secondSet))

  fmt.Println(returnTwoThings(4, 5))

  fmt.Println("add(3,5) -> ", add(3,5))
  fmt.Println("add(3,4,5) -> ", add(3,4,5))

  x := 0
  increment := func() (int) {
    x++
    return x
  }

  fmt.Println(increment())
  fmt.Println(increment())
  fmt.Println(increment())

  nextEven := makeEvenGenerator()
  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())

}

func average(xs []float64) (float64) {
  total := 0.0

  for _, value := range xs {
    total += value
  }
  return total / float64(len(xs))
}

func returnTwoThings(x int, y int) (sum int, product int){
  sum = x + y
  product = x * y
  return
}

func add(args ...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func makeEvenGenerator() func() uint {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
