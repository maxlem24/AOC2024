package advent

import (
	"fmt"
	"strings"
)

func Day4Part1(input string) {
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	count := 0

	dirs := [][2]int{{1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}}

	countXMAS := func(i int, j int) int {
		subCount := 0
		for _, dir := range dirs {
			maxI, maxJ := i+3*dir[0], j+3*dir[1]
			if maxI >= 0 && maxI < m && maxJ >= 0 && maxJ < n {
				if lines[i+dir[0]][j+dir[1]] == 'M' && lines[i+2*dir[0]][j+2*dir[1]] == 'A' && lines[maxI][maxJ] == 'S' {
					subCount++
				}
			}
		}
		return subCount
	}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'X' {
				count += countXMAS(i, j)
			}
		}
	}
	fmt.Println(count)
}

func Day4Part2(input string) {
	lines := strings.Split(input, "\n")
	m, n := len(lines), len(lines[0])
	count := 0

	isX_MAS := func(i int, j int) bool {
		if i-1 >= 0 && i+1 < m && j-1 >= 0 && j+1 < n {
			if (lines[i+1][j+1] == 'M' && lines[i-1][j-1] == 'S') || (lines[i+1][j+1] == 'S' && lines[i-1][j-1] == 'M') {
				if (lines[i+1][j-1] == 'M' && lines[i-1][j+1] == 'S') || (lines[i+1][j-1] == 'S' && lines[i-1][j+1] == 'M') {
					return true
				}
			}
		}
		return false
	}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'A' && isX_MAS(i, j) {
				count++
			}
		}
	}
	fmt.Println(count)
}
