package main

import (
  "fmt"
  "strings"
  "strconv"
  "os"
)

func main() {
  file, err := os.ReadFile("./input.txt")
  check(err)
  // input := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
  input := string(file)
  res := 0;
  pairs := strings.Split(strings.TrimSpace(input), "\n")
  for i := range pairs {
    sections := strings.Split(pairs[i], ",")
    section1Start,_ := strconv.Atoi(strings.Split(sections[0], "-")[0])
    section1End,_ := strconv.Atoi(strings.Split(sections[0], "-")[1])
    section2Start,_ := strconv.Atoi(strings.Split(sections[1], "-")[0])
    section2End,_ := strconv.Atoi(strings.Split(sections[1], "-")[1])
    if(
      ((section1Start <= section2End) && (section1Start >= section2Start)) ||
      ((section1End >= section2Start) && (section1End <= section2End)) ||
      ((section2Start >= section1Start) && (section2Start <= section1End)) ||
      ((section2End >= section1Start) && (section2End <= section1End))) {
      res++;
    }
  }
  fmt.Println(res);
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
