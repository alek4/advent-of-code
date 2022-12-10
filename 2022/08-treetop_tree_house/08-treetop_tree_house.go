package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func parseInput() (mat [][]int) {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		row := sc.Text()
		rowArr := []int{}
		for _, n := range row {
			num, _ := strconv.Atoi(string(n))

			rowArr = append(rowArr, num)
		}

		mat = append(mat, rowArr)
	}

	return mat
}

func partOne(matrix [][]int) {
	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || i == len(matrix)-1 {
				count += len(matrix)
				break
			}

			if j == 0 || j == len(matrix[i])-1 {
				count++
			} else {
				// oriz dx
				visible := true
				for k := j + 1; k < len(matrix[i]) && visible; k++ {
					if matrix[i][j] <= matrix[i][k] {
						visible = false
					}
				}

				if visible {
					count++
					continue
				}

				visible = true
				// oriz sx
				for k := j - 1; k >= 0 && visible; k-- {
					if matrix[i][j] <= matrix[i][k] {
						visible = false
					}
				}

				if visible {
					count++
					continue
				}

				visible = true
				// vert down
				for x := i + 1; x < len(matrix) && visible; x++ {
					if matrix[i][j] <= matrix[x][j] {
						visible = false
					}
				}

				if visible {
					count++
					continue
				}

				visible = true
				// vert up
				for x := i - 1; x >= 0 && visible; x-- {
					if matrix[i][j] <= matrix[x][j] {
						visible = false
					}
				}

				if visible {
					count++
					continue
				}
			}

		}
	}

	fmt.Println(count)

}

func partTwo(matrix [][]int) {
  scenicArr := []int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
      scenicScore := 1
			if i == 0 || i == len(matrix)-1 {
				// count += len(matrix)
				break
			}

			if j == 0 || j == len(matrix[i])-1 {
				// count++
			} else {
        var k int
				// oriz dx
				visible := true
				for k = j + 1; k < len(matrix[i]) && visible; k++ {
					if matrix[i][j] <= matrix[i][k] {
						visible = false
					}
				}

        scenicScore *= k - j - 1

				visible = true
				// oriz sx
				for k = j - 1; k >= 0 && visible; k-- {
					if matrix[i][j] <= matrix[i][k] {
						visible = false
					}
				}

        scenicScore *= j - k - 1

        var x int
				visible = true
				// vert down
				for x = i + 1; x < len(matrix) && visible; x++ {
					if matrix[i][j] <= matrix[x][j] {
						visible = false
					}
				}

        scenicScore *= x - i - 1

				visible = true
				// vert up
				for x = i - 1; x >= 0 && visible; x-- {
					if matrix[i][j] <= matrix[x][j] {
						visible = false
					}
				}

			  scenicScore *= i - x - 1

        scenicArr = append(scenicArr, scenicScore)
      }
		}
	}

  sort.Ints(scenicArr)
  fmt.Println(scenicArr[len(scenicArr)-1])
}

func main() {
	matrix := parseInput()
	// fmt.Println(matrix)
	partTwo(matrix)
}
