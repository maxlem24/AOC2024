package advent

import (
	"fmt"
	"strings"
)

func Day6Part1(input string) {
	lines := strings.Split(input, "\n")
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	count := 0
	currentDir := 0
	m, n := len(lines), len(lines[0])
	posRow, posCol := -1, -1
	for i := 0; posRow == -1 && i < m; i++ {
		for j := 0; posCol == -1 && j < n; j++ {
			switch lines[i][j] {
			case '^':
				posRow, posCol = i, j
				currentDir = 0
			case '>':
				posRow, posCol = i, j
				currentDir = 1
			case 'v':
				posRow, posCol = i, j
				currentDir = 2
			case '<':
				posRow, posCol = i, j
				currentDir = 3
			}
		}
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for posCol >= 0 && posCol < n && posRow >= 0 && posRow < m {
		newRow, newCol := posRow+dirs[currentDir][0], posCol+dirs[currentDir][1]
		if newCol >= 0 && newCol < n && newRow >= 0 && newRow < m && lines[newRow][newCol] == '#' {
			currentDir = (currentDir + 1) % 4
		} else {
			if !visited[posRow][posCol] {
				visited[posRow][posCol] = true
				count++
			}
			posRow, posCol = newRow, newCol
		}
	}
	fmt.Println(count)
}

func Day6Part2(input string) {
	lines := strings.Split(input, "\n")
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	currentDir := 0
	m, n := len(lines), len(lines[0])
	posRow, posCol := -1, -1
	for i := 0; posRow == -1 && i < m; i++ {
		for j := 0; posCol == -1 && j < n; j++ {
			switch lines[i][j] {
			case '^':
				posRow, posCol = i, j
				currentDir = 0
			case '>':
				posRow, posCol = i, j
				currentDir = 1
			case 'v':
				posRow, posCol = i, j
				currentDir = 2
			case '<':
				posRow, posCol = i, j
				currentDir = 3
			}
		}
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	obstruction := 0
	for posCol >= 0 && posCol < n && posRow >= 0 && posRow < m {
		newRow, newCol := posRow+dirs[currentDir][0], posCol+dirs[currentDir][1]
		if newCol >= 0 && newCol < n && newRow >= 0 && newRow < m && lines[newRow][newCol] == '#' {
			currentDir = (currentDir + 1) % 4
		} else {
			if newCol >= 0 && newCol < n && newRow >= 0 && newRow < m && !visited[newRow][newCol] && checkCycle(posRow, posCol, (currentDir+1)%4, newRow, newCol, lines) {
				obstruction++
			}
			if !visited[posRow][posCol] {
				visited[posRow][posCol] = true
			}
			posRow, posCol = newRow, newCol
		}
	}
	fmt.Println(obstruction)
}

func checkCycle(posRow int, posCol int, currentDir int, obsRow int, obsCol int, grid []string) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]int, m)
	for i := range visited {
		visited[i] = make([]int, n)
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for posCol >= 0 && posCol < n && posRow >= 0 && posRow < m {
		newRow, newCol := posRow+dirs[currentDir][0], posCol+dirs[currentDir][1]
		if newCol >= 0 && newCol < n && newRow >= 0 && newRow < m && (grid[newRow][newCol] == '#' || (newRow == obsRow && newCol == obsCol)) {
			currentDir = (currentDir + 1) % 4
		} else {
			if visited[posRow][posCol] == -1 {
				visited[posRow][posCol] = currentDir
			}else if visited[posRow][posCol] == currentDir{
				return true
			}
			posRow, posCol = newRow, newCol
		}
	}
	return false
}
