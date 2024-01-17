package main

import (
  "fmt"
  "strings"
  "os"
)

func main() {
  file, err := os.ReadFile("./input.txt")
  check(err)
  input := string(file)
  rounds := strings.Split(strings.TrimSpace(input), "\n")
  score := 0;
  for i := range rounds {
    initScore := score;
    round := string(int(rounds[i][0])+23)+string(rounds[i][2])
    score += (int(round[1])-88)*3
    coeff := 2
    if(round[1] == 'Y') {
      coeff = 0
    } else if(round[1] == 'Z'){
      coeff = 1
    }
    calc := int(round[0]-87)+coeff
    if(calc > 3)
      score += calc-3
    } else {
      score += calc
    }
    fmt.Println("round: ", string(round[0]), string(round[1]), ", score: +", score-initScore);
  }

  fmt.Println(score);
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}
