package main

import (
	"fmt"
	"strings"
)

func hasRepeatedChar(s string) bool {
	for _, c := range s {
		if strings.Count(s, string(c)) != 1 {
			return true
		}
	}

	return false
}

func main() {
	var in string
	fmt.Scanf("%s", &in)

	for idx := range in {
		if !hasRepeatedChar(in[idx:idx+14]) {
			fmt.Println(idx+14)
			break
		}
	}
}