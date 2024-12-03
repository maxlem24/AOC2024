package advent

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day3Part1(input string) {
	fmt.Println(sumMul(input))
}

func Day3Part2(input string) {
	dont := regexp.MustCompile(`don't\(\)`)
	do := regexp.MustCompile(`do\(\)`)
	nextDo := []int{0}
	nextDont := dont.FindStringIndex(input)
	total := 0
	for nextDo != nil && nextDont != nil {
		if nextDont[0] > nextDo[0] {
			total += sumMul(input[nextDo[0]:nextDont[0]])
		}
		input = input[nextDont[1]:]
		nextDo = do.FindStringIndex(input)
		nextDont = dont.FindStringIndex(input)
	}
	if nextDont == nil && nextDo != nil {
		total += sumMul(input[nextDo[0]:])
	}
	fmt.Println(total)
}

func sumMul(input string) int {
	mul := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	numbers := regexp.MustCompile(`[0-9]{1,3}`)
	valid := mul.FindAllStringSubmatch(input, -1)
	total := 0
	for _, match := range valid {
		var err error
		values := numbers.FindAllStringSubmatch(match[0], -1)
		val1, err := strconv.Atoi(values[0][0])
		if err != nil {
			panic(err)
		}
		val2, err := strconv.Atoi(values[1][0])
		if err != nil {
			panic(err)
		}
		total += val1 * val2
	}
	return total
}
