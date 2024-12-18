package advent

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day18Part1(input string) {
	size := 71
	lines := strings.Split(input, "\n")
	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}
	simulate := 1024
	for i := 0; i < simulate; i++ {
		coord := strings.Split(lines[i], ",")
		col, err := strconv.Atoi(coord[0])
		if err != nil {
			panic(err)
		}
		row, err := strconv.Atoi(coord[1])
		if err != nil {
			panic(err)
		}
		grid[row][col] = true
	}
	fmt.Println(memoryBFS(grid))
}

func Day18Part2(input string) {
	size := 71
	lines := strings.Split(input, "\n")
	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}
	for i := 0; i < len(lines); i++ {
		coord := strings.Split(lines[i], ",")
		col, err := strconv.Atoi(coord[0])
		if err != nil {
			panic(err)
		}
		row, err := strconv.Atoi(coord[1])
		if err != nil {
			panic(err)
		}
		grid[row][col] = true
		if memoryBFS(grid) == -1 {
			fmt.Println(lines[i])
			break
		}
	}
}

func memoryBFS(grid [][]bool) int {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	type coords struct {
		row  int
		col  int
		dist int
	}
	queue := make([]coords, 0)
	queue = append(queue, coords{0, 0, 0})
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current.row][current.col] {
			continue
		}

		if current.row == n-1 && current.col == n-1 {
			return current.dist
		}
		for _, dir := range dirs {
			next := coords{current.row + dir[0], current.col + dir[1], current.dist + 1}
			if next.row >= 0 && next.row < n && next.col >= 0 && next.col < n && !grid[next.row][next.col] && !visited[next.row][next.col] {
				queue = append(queue, next)
			}
		}
		slices.SortStableFunc(queue, func(a, b coords) int {
			return a.dist - b.dist
		})
		visited[current.row][current.col] = true
	}
	return -1
}
