package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func partOne() {
  sc := bufio.NewScanner(os.Stdin)

  scoreShape := map[string]int {"X": 1, "Y": 2, "Z": 3}
  mapToElfMove := map[string]string{"X": "A", "Y": "B","Z": "C"}
  totalScore := 0
  for sc.Scan() {
    row := sc.Text()
    rowSplit := strings.Split(row, " ")
    oppPlay := rowSplit[0]
    myPlay := rowSplit[1]
    
    totalScore += scoreShape[myPlay]
    if oppPlay == mapToElfMove[myPlay] {
      totalScore += 3
    } else {
      switch oppPlay {
        case "A":
          if myPlay == "Y" {
            totalScore += 6
          }
        
        case "B":
          if myPlay == "Z" {
            totalScore += 6
          }

        case "C":
          if myPlay == "X" {
            totalScore += 6
          }
      }
    }
    
  }

  fmt.Println(totalScore)
}

func main() {
  sc := bufio.NewScanner(os.Stdin)

  scoreMove := map[string]int {"A": 1, "B": 2, "C": 3}
  mapLoseMove := map[string]string{"A": "C", "B": "A", "C": "B"}
  mapWinMove := map[string]string{"A": "B", "B": "C", "C": "A"}
  totalScore := 0
  for sc.Scan() {
    row := sc.Text()
    rowSplit := strings.Split(row, " ")
    oppPlay := rowSplit[0]
    roundRes := rowSplit[1]
       
    if roundRes == "X" { // lose
      myMove := mapLoseMove[oppPlay]
      totalScore += scoreMove[myMove]
    } else if roundRes == "Y" { // draw
      totalScore += 3 + scoreMove[oppPlay]
    } else {
      myMove := mapWinMove[oppPlay]
      totalScore += 6 + scoreMove[myMove]
    }
  }

  fmt.Println(totalScore)

}
