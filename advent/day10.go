package advent

import (
	"fmt"
	"strings"
)

type topographicCoords struct {
	row int
	col int
}

func Day10Part1(input string) {
	topographicMap := strings.Split(input, "\n")
	m, n := len(topographicMap), len(topographicMap[0])
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	findTrailheads := func(start topographicCoords) (count int) {
		queue := []topographicCoords{start}
		visited := make(map[topographicCoords]bool)
		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]
			currentVal := topographicMap[current.row][current.col]
			if currentVal == '9' {
				if !visited[current] {
					count++
					visited[current] = true
				}
				continue
			}
			for _, dir := range dirs {
				nextTile := topographicCoords{row: current.row + dir[0], col: current.col + dir[1]}
				if nextTile.row >= 0 && nextTile.row < m && nextTile.col >= 0 && nextTile.col < n {
					if topographicMap[nextTile.row][nextTile.col] == currentVal+1 {
						queue = append(queue, nextTile)
					}
				}
			}
		}
		return
	}
	total := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if topographicMap[row][col] == '0' {
				total += findTrailheads(topographicCoords{row: row, col: col})
			}
		}
	}
	fmt.Println(total)
}

func Day10Part2(input string) {
	topographicMap := strings.Split(input, "\n")
	m, n := len(topographicMap), len(topographicMap[0])
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	findTrailheads := func(start topographicCoords) (count int) {
		queue := []topographicCoords{start}
		for len(queue) != 0 {
			current := queue[0]
			queue = queue[1:]
			currentVal := topographicMap[current.row][current.col]
			if currentVal == '9' {
				count++
				continue
			}
			for _, dir := range dirs {
				nextTile := topographicCoords{row: current.row + dir[0], col: current.col + dir[1]}
				if nextTile.row >= 0 && nextTile.row < m && nextTile.col >= 0 && nextTile.col < n {
					if topographicMap[nextTile.row][nextTile.col] == currentVal+1 {
						queue = append(queue, nextTile)
					}
				}
			}
		}
		return
	}
	total := 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if topographicMap[row][col] == '0' {
				total += findTrailheads(topographicCoords{row: row, col: col})
			}
		}
	}
	fmt.Println(total)
}
