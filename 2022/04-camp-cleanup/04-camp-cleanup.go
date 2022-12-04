package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringSliceToInt(s []string) []int {
  n := []int{}
  for _, v := range s {
    x, err := strconv.Atoi(v)
    if err != nil { os.Exit(-1) }
    n = append(n, x)
  }

  return n
}

func main() {
  sc := bufio.NewScanner(os.Stdin)

  res := 0
  for sc.Scan() {
    row := sc.Text()
    pairs := strings.Split(row, ",")
    
    pairOneStr := strings.Split(pairs[0], "-")
    pairTwoStr := strings.Split(pairs[1], "-")

    pairOne := stringSliceToInt(pairOneStr) 
    pairTwo := stringSliceToInt(pairTwoStr)

    if pairOne[0] >= pairTwo[0] && pairOne[0] <= pairTwo[1] || pairOne[1] >= pairTwo[0] && pairOne[1] <= pairTwo[1] {
      res += 1
    } else{ 
      if pairTwo[0] >= pairOne[0] && pairTwo[0] <= pairOne[0] || pairTwo[1] >= pairOne[0] && pairTwo[1] <= pairOne[1] {
        res += 1
      }
    }
  }
  fmt.Println(res)
}
