package advent

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type labyrinthTile struct {
	row      int
	col      int
	dirIndex int
	score    int
	path     map[labyrinthCoord]int
}

type labyrinthCoord struct {
	row int
	col int
}

type TileHeap []labyrinthTile

func (h TileHeap) Len() int           { return len(h) }
func (h TileHeap) Less(i, j int) bool { return h[i].score < h[j].score }
func (h TileHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TileHeap) Push(x any) {
	*h = append(*h, x.(labyrinthTile))
}

func (h *TileHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Day16Part1(input string) {
	score, _ := dijktras(input)
	fmt.Println(score)
}

func Day16Part2(input string) {
	_, path := dijktras(input)
	fmt.Println(len(path))
}

func dijktras(input string) (int, map[labyrinthCoord]int) {
	path := make(map[labyrinthCoord]int)
	score := math.MaxInt32
	labyrinth := strings.Split(input, "\n")
	m, n := len(labyrinth), len(labyrinth[0])
	visited := make([][][4]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([][4]bool, n)
	}
	queue := &TileHeap{}
	heap.Init(queue)
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	heap.Push(queue, labyrinthTile{row: m - 2, col: 1, dirIndex: 0, score: 0, path: make(map[labyrinthCoord]int)})
	for queue.Len() > 0 {
		current := heap.Pop(queue).(labyrinthTile)
		current.path[labyrinthCoord{current.row, current.col}] = current.score
		if current.score > score {
			continue
		}
		if current.row == 1 && current.col == n-2 {
			if current.score < score {
				score = current.score
				path = copyMap(current.path)
			} else if current.score == score {
				for coord, dist := range current.path {
					path[coord] = dist
				}
			}
		}
		for i, dir := range dirs {
			if (i+2)%4 == current.dirIndex {
				continue
			}
			score := 1
			if (i+1)%4 == current.dirIndex || (i+3)%4 == current.dirIndex {
				score += 1000
			}
			nextPos := labyrinthTile{row: current.row + dir[0], col: current.col + dir[1], dirIndex: i, score: current.score + score, path: copyMap(current.path)}
			if !visited[nextPos.row][nextPos.col][nextPos.dirIndex] && labyrinth[nextPos.row][nextPos.col] != '#' {
				heap.Push(queue, nextPos)
			}
		}
		visited[current.row][current.col][current.dirIndex] = true
	}
	return score, path
}

func copyMap(path map[labyrinthCoord]int) map[labyrinthCoord]int {
	new := make(map[labyrinthCoord]int, len(path))
	for key, value := range path {
		new[key] = value
	}
	return new
}
