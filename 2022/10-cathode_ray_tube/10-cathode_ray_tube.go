package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	sc := bufio.NewScanner(os.Stdin)

	cycles := 1
	x := 1
	signal_strength := 20
	res := 0
	for sc.Scan() {
		op := sc.Text()
		instr := strings.Split(op, " ")[0]

		if cycles == signal_strength {
			fmt.Println(signal_strength, x)
			res += signal_strength * x
			signal_strength += 40
		}

		if instr == "addx" {
			cycles++

			if cycles == signal_strength {
				fmt.Println(signal_strength, x)
				res += signal_strength * x
				signal_strength += 40
			}

			value, _ := strconv.Atoi(strings.Split(op, " ")[1])
			x += value
		}

		cycles++
	}

	fmt.Println(res)

}

func drawCRT(x int, crtPos *int) {
	// fmt.Println(x, cycles)
	if x-1 == *crtPos || x == *crtPos || x+1 == *crtPos {

		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	*crtPos++
}

func partTwo() {
	sc := bufio.NewScanner(os.Stdin)

	cycles := 1
	x := 1
	signal_strength := 40
	crtPos := 0
	for sc.Scan() {
		op := sc.Text()
		instr := strings.Split(op, " ")[0]

		if cycles == signal_strength+1 {
			// fmt.Println(signal_strength, x)
			signal_strength += 40
			fmt.Println()
			crtPos = 0
		}

		if instr == "addx" {
			drawCRT(x, &crtPos)

			cycles++

			if cycles == signal_strength+1 {
				// fmt.Println(signal_strength, x)
				signal_strength += 40
				fmt.Println()
				crtPos = 0
			}

			drawCRT(x, &crtPos)
			value, _ := strconv.Atoi(strings.Split(op, " ")[1])
			x += value
			cycles++
		} else {

			drawCRT(x, &crtPos)
			cycles++
		}
	}

}

func main() {
	partTwo()
}
