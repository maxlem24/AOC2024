package advent

import (
	"fmt"
	"strings"
)

type antennaCoords struct {
	row int
	col int
}

type antennaVect struct {
	deltaRow int
	deltaCol int
}

func (antenna antennaCoords) addVect(vect antennaVect, coef int) antennaCoords {
	return antennaCoords{row: antenna.row + coef*vect.deltaRow, col: antenna.col + coef*vect.deltaCol}
}

func Day8Part1(input string) {
	antennaMap := strings.Split(input, "\n")
	m, n := len(antennaMap), len(antennaMap[0])
	frequencies := make(map[byte][]antennaCoords)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if val := antennaMap[row][col]; val != '.' {
				frequencies[val] = append(frequencies[val], antennaCoords{row: row, col: col})
			}
		}
	}
	antinodeMap := make([][]bool, len(antennaMap))
	for row := 0; row < m; row++ {
		antinodeMap[row] = make([]bool, n)
	}
	count := 0
	for _, coords := range frequencies {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				vect := antennaVect{deltaRow: coords[i].row - coords[j].row, deltaCol: coords[i].col - coords[j].col}
				node1 := coords[i].addVect(vect, 1)
				if node1.row >= 0 && node1.row < m && node1.col >= 0 && node1.col < n {
					if !antinodeMap[node1.row][node1.col] {
						antinodeMap[node1.row][node1.col] = true
						count++
					}
				}
				node2 := coords[j].addVect(vect, -1)
				if node2.row >= 0 && node2.row < m && node2.col >= 0 && node2.col < n {
					if !antinodeMap[node2.row][node2.col] {
						antinodeMap[node2.row][node2.col] = true
						count++
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func Day8Part2(input string) {
	antennaMap := strings.Split(input, "\n")
	m, n := len(antennaMap), len(antennaMap[0])
	frequencies := make(map[byte][]antennaCoords)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if val := antennaMap[row][col]; val != '.' {
				frequencies[val] = append(frequencies[val], antennaCoords{row: row, col: col})
			}
		}
	}
	antinodeMap := make([][]bool, len(antennaMap))
	for row := 0; row < m; row++ {
		antinodeMap[row] = make([]bool, n)
	}
	count := 0
	for _, coords := range frequencies {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				vect := antennaVect{deltaRow: coords[i].row - coords[j].row, deltaCol: coords[i].col - coords[j].col}
				coef := 0
				for node := coords[i].addVect(vect, coef); node.row >= 0 && node.row < m && node.col >= 0 && node.col < n; node = coords[i].addVect(vect, coef) {
					if !antinodeMap[node.row][node.col] {
						antinodeMap[node.row][node.col] = true
						count++
					}
					coef++
				}
				coef = -1
				for node := coords[i].addVect(vect, coef); node.row >= 0 && node.row < m && node.col >= 0 && node.col < n; node = coords[i].addVect(vect, coef) {
					if !antinodeMap[node.row][node.col] {
						antinodeMap[node.row][node.col] = true
						count++
					}
					coef--
				}
			}
		}
	}
	fmt.Println(count)
}
