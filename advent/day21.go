package advent

import (
	"fmt"
	"strconv"
	"strings"
)

func Day21Part1(input string) {
	count := 0
	for _, code := range strings.Split(input, "\n") {
		moves := numpadSequence(code)
		firstRobot := dirpadSequence(moves)
		secondRobot := dirpadSequence(firstRobot)
		val, err := strconv.Atoi(code[:3])
		if err != nil {
			panic(err)
		}
		count += len(secondRobot) * val
	}
	fmt.Println(count)
}

func Day21Part2(input string) {
	count := 0
	cache := make(map[string][]int)
	for _, code := range strings.Split(input, "\n") {
		moves := numpadSequence(code)
		length := countSequence(moves, 25, 1, cache)
		val, err := strconv.Atoi(code[:3])
		if err != nil {
			panic(err)
		}
		count += length * val
	}
	fmt.Println(count)
}

func countSequence(moves string, nbRobots int, currentRobot int, cache map[string][]int) int {
	if val, exist := cache[moves]; exist {
		if len(val) >= currentRobot && val[currentRobot-1] != 0 {
			return val[currentRobot-1]
		}
	} else {
		cache[moves] = make([]int, nbRobots)
	}
	nextMoves := dirpadSequence(moves)
	if currentRobot == nbRobots {
		return len(nextMoves)
	}
	count := 0

	for _, step := range strings.Split(strings.ReplaceAll(nextMoves, "A", "A "), " ") {
		count += countSequence(step, nbRobots, currentRobot+1, cache)
	}
	cache[moves][currentRobot-1] = count

	return count

}

func numpadSequence(code string) string {
	type coord struct {
		x int
		y int
	}
	pos := make(map[byte]coord)
	pos['7'] = coord{0, 0}
	pos['8'] = coord{1, 0}
	pos['9'] = coord{2, 0}
	pos['4'] = coord{0, 1}
	pos['5'] = coord{1, 1}
	pos['6'] = coord{2, 1}
	pos['1'] = coord{0, 2}
	pos['2'] = coord{1, 2}
	pos['3'] = coord{2, 2}
	pos['0'] = coord{1, 3}
	pos['A'] = coord{2, 3}
	result := ""
	var start byte = 'A'
	for _, char := range []byte(code) {
		dx := diffAbs(pos[start].x, pos[char].x)
		dy := diffAbs(pos[start].y, pos[char].y)
		xStr := ""
		yStr := ""
		if pos[start].x > pos[char].x {
			xStr = strings.Repeat("<", dx)
		} else {
			xStr = strings.Repeat(">", dx)
		}

		if pos[start].y > pos[char].y {
			yStr = strings.Repeat("^", dy)
		} else {
			yStr = strings.Repeat("v", dy)
		}

		if pos[start].y == 3 && pos[char].x == 0 {
			result += yStr + xStr
		} else if pos[start].x == 0 && pos[char].y == 3 {
			result += xStr + yStr
		} else if pos[start].x > pos[char].x {
			result += xStr + yStr
		} else {
			result += yStr + xStr
		}
		result += "A"
		start = char
	}
	return result
}

func dirpadSequence(code string) string {
	type coord struct {
		x int
		y int
	}
	pos := make(map[byte]coord)
	pos['^'] = coord{1, 0}
	pos['A'] = coord{2, 0}
	pos['<'] = coord{0, 1}
	pos['v'] = coord{1, 1}
	pos['>'] = coord{2, 1}
	instructions := strings.Split(strings.ReplaceAll(code, "A", "A "), " ")
	cache := make(map[string]string)
	result := ""
	for _, instruction := range instructions {
		if val, exist := cache[instruction]; exist {
			result += val
			continue
		}
		str := ""
		var start byte = 'A'
		for _, char := range []byte(instruction) {
			dx := diffAbs(pos[start].x, pos[char].x)
			dy := diffAbs(pos[start].y, pos[char].y)
			xStr := ""
			yStr := ""
			if pos[start].x > pos[char].x {
				xStr = strings.Repeat("<", dx)
			} else {
				xStr = strings.Repeat(">", dx)
			}

			if pos[start].y > pos[char].y {
				yStr = strings.Repeat("^", dy)
			} else {
				yStr = strings.Repeat("v", dy)
			}

			if pos[start].x == 0 && pos[char].y == 0 {
				str += xStr + yStr
			} else if pos[start].y == 0 && pos[char].x == 0 {
				str += yStr + xStr
			} else if pos[char].x < pos[start].x {
				str += xStr + yStr
			} else {
				str += yStr + xStr
			}
			start = char
			str += "A"
		}
		cache[instruction] = str
		result += str
	}
	return result
}
