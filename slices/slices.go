package main

import "fmt"

func main() {
  arr := [5]float64{1,2,3,4,5}
  fmt.Println(arr[0:5])
  fmt.Println(arr[1:4])
  fmt.Println(arr[3:])
  fmt.Println(arr[:2])

  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice2)

  slice3 := []int{1,2,3}
  slice4 := make([]int, 2)
  copy(slice4, slice3)
  fmt.Println(slice3, slice4)

}
