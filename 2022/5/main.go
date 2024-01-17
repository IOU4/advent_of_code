package main

import (
  "fmt"
  "os"
)

func main() {
  file, err := os.ReadFile("./input.txt")
  check(err)
  // input := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
  input := string(file)
  fmt.Println(input)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

