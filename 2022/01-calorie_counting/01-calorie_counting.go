package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(arr []int) (int, int) {
  min := arr[0]
  var i int
  for idx, n := range arr {
    if n < min {
      min = n
      i = idx
    }
  }

  return i, min
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	elves := []int{0}
	i := 0
	for sc.Scan() {
		if (sc.Text() == "") {
			i++
			elves = append(elves, 0)
			continue
		}

		n, err := strconv.Atoi(sc.Text())
		if err == nil {
			elves[i] += n
		}
	}

	max := 0
	for _, n := range elves {
		if n > max {
			max = n
		}
	}

	fmt.Println(max)

  topThree := make([]int, 3)
  for _, elf := range elves {
    idx, min := min(topThree)
    if elf > min {
      topThree[idx] = elf
    }
  }
	 

  fmt.Println(topThree[0] + topThree[1] + topThree[2])
}
