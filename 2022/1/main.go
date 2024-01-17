package main

import (
  "fmt"
  "strings"
  "strconv"
  "os"
)

func main() {
  file, err := os.ReadFile("./input.txt")
  if err != nil {
    panic(err)
  }
  input := string(file);
  res := strings.Split(input, "\n\n")
  max1 := 0;
  max2 := 0;
  max3 := 0;
  for i:= 0; i<len(res); i++ {
    elfCarry := strings.Split(strings.TrimSpace(res[i]), "\n");
    calories := 0;
    for j := 0; j<len(elfCarry); j++ {
      parsedInt, err := strconv.Atoi(elfCarry[j])
      if err != nil {
        panic(err)
      }
      calories += parsedInt;
    }
    if (max1 < calories) {
      max3 = max2;
      max2 = max1;
      max1 = calories;
    } else if (max2 < calories) {
      max3 = max2;
      max2 = calories
    } else if ( max3 < calories) {
      max3 = calories;
    }
  }
  fmt.Println("max: ", max1+max2+max3)
}
