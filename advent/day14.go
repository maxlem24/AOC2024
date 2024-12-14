package advent

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	x  int
	y  int
	vX int
	vY int
}

func Day14Part1(input string) {
	robots := strings.Split(input, "\n")
	re := regexp.MustCompile(`-?\d+`)
	height, wide := 103, 101
	quarters := make([]int, 4)
	for _, robot := range robots {
		values := re.FindAllStringSubmatch(robot, -1)
		x, _ := strconv.Atoi(values[0][0])
		y, _ := strconv.Atoi(values[1][0])
		vX, _ := strconv.Atoi(values[2][0])
		vY, _ := strconv.Atoi(values[3][0])
		finalX, finalY := ((x+100*vX)%wide+wide)%wide, ((y+100*vY)%height+height)%height
		if finalX < wide/2 {
			if finalY < height/2 {
				quarters[0]++
			} else if finalY > height/2 {
				quarters[1]++
			}
		} else if finalX > wide/2 {
			if finalY < height/2 {
				quarters[2]++
			} else if finalY > height/2 {
				quarters[3]++
			}
		}
	}
	fmt.Println(quarters[0] * quarters[1] * quarters[2] * quarters[3])
}

func Day14Part2(input string) {
	robots := strings.Split(input, "\n")
	robotsPos := make([]Robot, len(robots))
	re := regexp.MustCompile(`-?\d+`)
	height, wide := 103, 101

	printRobots := func() bool {
		grid := make([][]bool, height)
		for i := 0; i < height; i++ {
			grid[i] = make([]bool, wide)
		}
		for _, robot := range robotsPos {
			grid[robot.y][robot.x] = true
		}
		shouldPrint := false
		for i := 1; i < height-1 && !shouldPrint; i++ {
			for j := 1; j < wide-1; j++ {
				if  grid[i-1][j-1] && grid[i][j-1] && grid[i+1][j-1] && grid[i-1][j] && grid[i][j] && grid[i+1][j] && grid[i-1][j+1] && grid[i][j+1] && grid[i+1][j+1] {
					shouldPrint = true
					break
				}
			}
		}
		if !shouldPrint {
			return false
		}
		output := ""
		for i := 0; i < height; i++ {
			for j := 0; j < wide; j++ {
				if grid[i][j] {
					output += "X"
				} else {
					output += " "
				}
			}
			output += "\n"
		}
		err := os.WriteFile("../files/day14_output.txt", []byte(output), 0666)
		if err != nil {
			panic(err)
		}
		return true
	}
	startFrame := 0
	for i, robot := range robots {
		values := re.FindAllStringSubmatch(robot, -1)
		x, _ := strconv.Atoi(values[0][0])
		y, _ := strconv.Atoi(values[1][0])
		vX, _ := strconv.Atoi(values[2][0])
		vY, _ := strconv.Atoi(values[3][0])
		robotsPos[i] = Robot{x: ((x+startFrame*vX)%wide + wide) % wide, y: ((y+startFrame*vY)%height + height) % height, vX: vX, vY: vY}
	}
	for taped, err := fmt.Scanln(); err == nil && taped == 0; taped, err = fmt.Scanln() {
		for !printRobots() {
			for i, robot := range robotsPos {
				robotsPos[i].x = ((robot.x+robot.vX)%wide + wide) % wide
				robotsPos[i].y = ((robot.y+robot.vY)%height + height) % height
			}
			startFrame++
		}
	}
	fmt.Println(startFrame)
}
