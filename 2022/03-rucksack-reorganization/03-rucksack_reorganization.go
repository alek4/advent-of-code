package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getPriority(c rune) int {
  if unicode.IsLower(c) {
    // fmt.Println(c)
    return int(c - 96)
  } else {
    return int(c - 38)
  }
}

func parOne() {
  sc := bufio.NewScanner(os.Stdin)

  res := 0
  for sc.Scan() {
    row := sc.Text()
    firstHalf := row[:len(row)/2]
    secondHalf := row[len(row)/2:]

    // fmt.Println(firstHalf, secondHalf)

    var lastFound rune
    for _, c := range firstHalf {
      for _, c2 := range secondHalf {
        if c == c2 && c != lastFound {
          res += getPriority(c)
          lastFound = c
          // fmt.Println(getPriority(c))
        }
      }
    }
  }

  fmt.Println(res)
}

func main() {
  sc := bufio.NewScanner(os.Stdin)

  res := 0
  group := []string{}
  for sc.Scan() {
    row := sc.Text()
    if len(group) < 3 {
      group = append(group, row)
    }

    if len(group) == 3 {
      var lastFound rune
      for _, c := range group[0] {
        if c != lastFound && strings.ContainsRune(group[1], c) && strings.ContainsRune(group[2], c) {
          res += getPriority(c)
          lastFound = c
        }
      }

      group = []string{}
    }
  }

  fmt.Println(res)
}
