package main

import(
  "fmt"
)

func main() {
  a, b := returnTwo()[0];
  fmt.Println(b)
}

func returnTwo() []int {
  return []int{1, 2}
}
