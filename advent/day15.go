package advent

import (
	"fmt"
	"strings"
)

func Day15Part1(input string) {
	parts := strings.Split(input, "\n\n")
	grid := strings.Split(parts[0], "\n")
	m, n := len(grid), len(grid[0])
	warehouse := make([][]int, m)
	for i := range warehouse {
		warehouse[i] = make([]int, n)
	}
	robotRow, robotCol := -1, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'O' {
				warehouse[i][j] = 1
			} else if grid[i][j] == '#' {
				warehouse[i][j] = -1
			} else if grid[i][j] == '@' {
				robotRow, robotCol = i, j
			}
		}
	}
	dirBlocked := ' '
	dirs := map[rune][2]int{'^': {-1, 0}, '>': {0, 1}, 'v': {1, 0}, '<': {0, -1}}
	for _, instruction := range parts[1] {
		if instruction == dirBlocked {
			continue
		} else {
			if dir, exist := dirs[instruction]; exist {
				blocked := true
				row, col := robotRow+dir[0], robotCol+dir[1]
				for blocked && row >= 0 && row < m && col >= 0 && col < n {
					if warehouse[row][col] == 0 {
						blocked = false
					} else if warehouse[row][col] == -1 {
						break
					} else {
						row += dir[0]
						col += dir[1]
					}
				}
				if blocked {
					dirBlocked = instruction
					continue
				} else {
					dirBlocked = ' '
				}
				for row != robotRow+dir[0] || col != robotCol+dir[1] {
					prevRow, prevCol := row-dir[0], col-dir[1]
					warehouse[row][col], warehouse[prevRow][prevCol] = warehouse[prevRow][prevCol], warehouse[row][col]
					row, col = prevRow, prevCol
				}
				robotRow, robotCol = row, col
			}
		}
	}
	score := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if warehouse[i][j] == 1 {
				score += 100*i + j
			}
		}
	}
	fmt.Println(score)
}

func Day15Part2(input string) {
	parts := strings.Split(input, "\n\n")
	newWarehouse := strings.ReplaceAll(parts[0], "#", "##")
	newWarehouse = strings.ReplaceAll(newWarehouse, "O", "[]")
	newWarehouse = strings.ReplaceAll(newWarehouse, ".", "..")
	newWarehouse = strings.ReplaceAll(newWarehouse, "@", "@.")
	grid := strings.Split(newWarehouse, "\n")
	m, n := len(grid), len(grid[0])
	warehouse := make([][]int, m)
	for i := range warehouse {
		warehouse[i] = make([]int, n)
	}
	robotRow, robotCol := -1, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '[' {
				warehouse[i][j] = 1
			} else if grid[i][j] == ']' {
				warehouse[i][j] = 2
			} else if grid[i][j] == '#' {
				warehouse[i][j] = -1
			} else if grid[i][j] == '@' {
				robotRow, robotCol = i, j
			}
		}
	}
	moveVertical := func(row int, col int, dir [2]int) bool {
		queue := [][2]int{{row, col}}
		if warehouse[row][col] == 2 {
			queue = append(queue, [2]int{row, col - 1})
		} else {
			queue = append(queue, [2]int{row, col + 1})
		}
		visited := make(map[[2]int]bool)
		visitOrder := make([][2]int, 0)

		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]

			if _, ok := visited[current]; ok {
				continue
			}

			visited[current] = true
			visitOrder = append(visitOrder, current)

			newRow, newCol := current[0]+dir[0], current[1]+dir[1]
			switch warehouse[newRow][newCol] {
			case 0:
				continue
			case -1:
				return false
			case 2:
				queue = append(queue, [2]int{newRow, newCol})
				queue = append(queue, [2]int{newRow, newCol - 1})
			case 1:
				queue = append(queue, [2]int{newRow, newCol})
				queue = append(queue, [2]int{newRow, newCol + 1})
			}
		}
		for i := len(visitOrder) - 1; i >= 0; i-- {
			newRow, newCol := visitOrder[i][0]+dir[0], visitOrder[i][1]+dir[1]
			warehouse[newRow][newCol] = warehouse[visitOrder[i][0]][visitOrder[i][1]]
			warehouse[visitOrder[i][0]][visitOrder[i][1]] = 0
		}

		return true
	}
	dirBlocked := ' '
	dirs := map[rune][2]int{'^': {-1, 0}, '>': {0, 1}, 'v': {1, 0}, '<': {0, -1}}
	for _, instruction := range parts[1] {
		if instruction == dirBlocked {
			continue
		} else {
			if dir, exist := dirs[instruction]; exist {
				row, col := robotRow+dir[0], robotCol+dir[1]
				if dir[0] != 0 {
					if warehouse[row][col] == -1 {
						dirBlocked = instruction
					} else if warehouse[row][col] == 1 || warehouse[row][col] == 2 {
						if moveVertical(row, col, dir) {
							dirBlocked = ' '
							robotRow, robotCol = row, col
						} else {
							dirBlocked = instruction
						}
					} else if warehouse[row][col] == 0 {
						dirBlocked = ' '
						robotRow, robotCol = row, col
					}
					continue
				}
				blocked := true
				for blocked && row >= 0 && row < m && col >= 0 && col < n {
					if warehouse[row][col] == 0 {
						blocked = false
					} else if warehouse[row][col] == -1 {
						break
					} else {
						row += dir[0]
						col += dir[1]
					}
				}
				if blocked {
					dirBlocked = instruction
					continue
				} else {
					dirBlocked = ' '
				}
				for row != robotRow+dir[0] || col != robotCol+dir[1] {
					prevRow, prevCol := row-dir[0], col-dir[1]
					warehouse[row][col], warehouse[prevRow][prevCol] = warehouse[prevRow][prevCol], warehouse[row][col]
					row, col = prevRow, prevCol
				}
				robotRow, robotCol = row, col
			}
		}
	}
	score := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if warehouse[i][j] == 1 {
				score += 100*i + j
			}
		}
	}
	fmt.Println(score)
}
