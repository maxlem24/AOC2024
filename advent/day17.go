package advent

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day17Part1(input string) {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllStringSubmatch(input, -1)
	regA, err := strconv.Atoi(numbers[0][0])
	if err != nil {
		panic(err)
	}
	regB, err := strconv.Atoi(numbers[1][0])
	if err != nil {
		panic(err)
	}
	regC, err := strconv.Atoi(numbers[2][0])
	if err != nil {
		panic(err)
	}
	instructions := make([]int, len(numbers)-3)
	for i := 0; i < len(instructions); i++ {
		instructions[i] = int(numbers[3+i][0][0] - '0')
	}
	fmt.Println(runProgram(instructions, regA, regB, regC))
}

func Day17Part2(input string) {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllStringSubmatch(input, -1)
	regB, err := strconv.Atoi(numbers[1][0])
	if err != nil {
		panic(err)
	}
	regC, err := strconv.Atoi(numbers[2][0])
	if err != nil {
		panic(err)
	}
	instructions := make([]int, len(numbers)-3)
	for i := 0; i < len(instructions); i++ {
		instructions[i] = int(numbers[3+i][0][0] - '0')
	}
	program := strings.Split(input, "\n")[4][9:]
	n := len(program)
	start := 1
	for len(runProgram(instructions, start, regB, regC)) < n {
		start *= 2
	}
	step := start
	for i := 1; i <= len(program); i += 2 {
		for runProgram(instructions, start, regB, regC)[n-i:] != program[n-i:] {
			start += step
		}
		step /= 8
	}
	fmt.Println(start)
}

func runProgram(instructions []int, regA int, regB int, regC int) string {
	output := ""
	ptr := 0
	for ptr < len(instructions) {
		var val int
		switch instructions[ptr+1] {
		case 0, 1, 2, 3:
			val = instructions[ptr+1]
		case 4:
			val = regA
		case 5:
			val = regB
		case 6:
			val = regC
		}
		switch instructions[ptr] {
		case 0:
			regA = regA / pow2(val)
			ptr += 2
		case 1:
			regB = regB ^ instructions[ptr+1]
			ptr += 2
		case 2:
			regB = val % 8
			ptr += 2
		case 3:
			if regA == 0 {
				ptr += 2
			} else {
				ptr = val
			}
		case 4:
			regB = regB ^ regC
			ptr += 2
		case 5:
			if len(output) == 0 {
				output = strconv.Itoa(val % 8)
			} else {
				output += "," + strconv.Itoa(val%8)
			}
			ptr += 2
		case 6:
			regB = regA / pow2(val)
			ptr += 2
		case 7:
			regC = regA / pow2(val)
			ptr += 2
		}
	}
	return output
}

func pow2(pow int) int {
	val := 1
	for i := 1; i <= pow; i++ {
		val *= 2
	}
	return val
}
