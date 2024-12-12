package advent

import (
	"fmt"
	"strings"
)

func Day12Part1(input string) {
	garden := strings.Split(input, "\n")
	m, n := len(garden), len(garden[0])
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	explorePlots := func(startRow int, startCol int) int {
		plant := garden[startRow][startCol]
		visited[startRow][startCol] = true
		area := 1
		perimeter := 0
		queue := [][2]int{{startRow, startCol}}
		for len(queue) != 0 {
			currentRow, currentCol := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, dir := range dirs {
				nextRow, nextCol := currentRow+dir[0], currentCol+dir[1]
				if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && garden[nextRow][nextCol] == plant {
					if !visited[nextRow][nextCol] {
						area++
						queue = append(queue, [2]int{nextRow, nextCol})
						visited[nextRow][nextCol] = true
					}
				} else {
					perimeter++
				}
			}
		}
		return perimeter * area
	}
	price := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if !visited[row][col] {
				price += explorePlots(row, col)
			}
		}
	}
	fmt.Println(price)
}

func Day12Part2(input string) {
	garden := strings.Split(input, "\n")
	m, n := len(garden), len(garden[0])
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	areCoordValid := func(row int, col int) bool {
		return row >= 0 && row < m && col >= 0 && col < n
	}

	explorePlots := func(startRow int, startCol int) int {
		plant := garden[startRow][startCol]
		visited[startRow][startCol] = true
		area := 1
		sides := 0
		queue := [][2]int{{startRow, startCol}}
		for len(queue) != 0 {
			currentRow, currentCol := queue[0][0], queue[0][1]
			queue = queue[1:]
			for indexDir, dir := range dirs {
				nextRow, nextCol := currentRow+dir[0], currentCol+dir[1]
				if areCoordValid(nextRow, nextCol) && garden[nextRow][nextCol] == plant {
					if !visited[nextRow][nextCol] {
						area++
						queue = append(queue, [2]int{nextRow, nextCol})
						visited[nextRow][nextCol] = true
					}
				} else {
					innerRow, innerCol := nextRow+dirs[(indexDir+1)%4][0], nextCol+dirs[(indexDir+1)%4][1]
					outerRow, outerCol := currentRow+dirs[(indexDir+1)%4][0], currentCol+dirs[(indexDir+1)%4][1]
					if (areCoordValid(innerRow, innerCol) && garden[innerRow][innerCol] == plant) || (!areCoordValid(outerRow, outerCol)|| garden[outerRow][outerCol] != plant) {
						
						sides++
					}
				}
			}
		}
		return sides * area
	}
	price := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if !visited[row][col] {
				price += explorePlots(row, col)
			}
		}
	}
	fmt.Println(price)
}
