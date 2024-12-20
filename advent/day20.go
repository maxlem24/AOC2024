package advent

import (
	"fmt"
	"math"
	"strings"
)

func Day20Part1(input string) {
	type cheatCoords struct {
		row int
		col int
	}

	labyrinth := strings.Split(input, "\n")
	queue := make([]cheatCoords, 0)
	path := make([]cheatCoords, 0)
	m, n := len(labyrinth), len(labyrinth[0])
	dist := make([][]int, m)
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var start cheatCoords

	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
			if labyrinth[i][j] == 'E' {
				start = cheatCoords{i, j}
			}
		}
	}

	queue = append(queue, start)
	dist[start.row][start.col] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentDist := dist[current.row][current.col]
		for _, dir := range dirs {
			nextRow, nextCol := current.row+dir[0], current.col+dir[1]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && labyrinth[nextRow][nextCol] != '#' && dist[nextRow][nextCol] > currentDist+1 {
				dist[nextRow][nextCol] = currentDist + 1
				queue = append(queue, cheatCoords{nextRow, nextCol})
			}
		}
		path = append(path, current)
	}

	count := 0
	cheatDirs := [][2]int{{0, 2}, {1, 1}, {2, 0}, {1, -1}, {0, -2}, {-1, -1}, {-2, 0}, {-1, 1}}
	for len(path) > 0 {
		current := path[0]
		path = path[1:]
		currentDist := dist[current.row][current.col]
		for _, dir := range cheatDirs {
			nextRow, nextCol := current.row+dir[0], current.col+dir[1]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && labyrinth[nextRow][nextCol] != '#' && dist[nextRow][nextCol]-currentDist > 2 {
				if dist[nextRow][nextCol]-currentDist-2 >= 100 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func Day20Part2(input string) {
	type cheatCoords struct {
		row int
		col int
	}

	labyrinth := strings.Split(input, "\n")
	queue := make([]cheatCoords, 0)
	path := make([]cheatCoords, 0)
	m, n := len(labyrinth), len(labyrinth[0])
	dist := make([][]int, m)
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var start cheatCoords

	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
			if labyrinth[i][j] == 'E' {
				start = cheatCoords{i, j}
			}
		}
	}

	queue = append(queue, start)
	dist[start.row][start.col] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentDist := dist[current.row][current.col]
		for _, dir := range dirs {
			nextRow, nextCol := current.row+dir[0], current.col+dir[1]
			if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && labyrinth[nextRow][nextCol] != '#' && dist[nextRow][nextCol] > currentDist+1 {
				dist[nextRow][nextCol] = currentDist + 1
				queue = append(queue, cheatCoords{nextRow, nextCol})
			}
		}
		path = append(path, current)
	}

	count := 0

	for cheatSize := 2; cheatSize <= 20; cheatSize++ {
		cheatDirs := make([][2]int, 0)
		for row, col := 1, cheatSize-1; row < cheatSize && col > 0; row, col = row+1, col-1 {
			cheatDirs = append(cheatDirs, [2]int{row, col})
			cheatDirs = append(cheatDirs, [2]int{-row, col})
			cheatDirs = append(cheatDirs, [2]int{row, -col})
			cheatDirs = append(cheatDirs, [2]int{-row, -col})
		}
		cheatDirs = append(cheatDirs, [2]int{0, cheatSize})
		cheatDirs = append(cheatDirs, [2]int{cheatSize, 0})
		cheatDirs = append(cheatDirs, [2]int{0, -cheatSize})
		cheatDirs = append(cheatDirs, [2]int{-cheatSize, 0})

		for _, current := range path {
			currentDist := dist[current.row][current.col]
			for _, dir := range cheatDirs {
				nextRow, nextCol := current.row+dir[0], current.col+dir[1]
				if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && labyrinth[nextRow][nextCol] != '#' && dist[nextRow][nextCol]-currentDist > cheatSize {
					if dist[nextRow][nextCol]-currentDist-cheatSize >= 100 {
						count++
					}
				}
			}
		}
	}
	fmt.Println(count)
}
