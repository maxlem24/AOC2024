package advent

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day7Part1(input string) {
	equations := strings.Split(input, "\n")
	re := regexp.MustCompile(`\d+`)
	sum := 0
	for _, equation := range equations {
		var err error
		nums := re.FindAllStringSubmatch(equation, -1)
		goal, err := strconv.Atoi(nums[0][0])
		if err != nil {
			panic(err)
		}
		values := make([]int, len(nums)-1)
		for i := 1; i < len(nums); i++ {
			val, err := strconv.Atoi(nums[i][0])
			if err != nil {
				panic(err)
			}
			values[i-1] = val
		}
		if isSolvable(values, goal, []string{"*", "+"}) {
			sum += goal
		}
	}
	fmt.Println(sum)
}

func Day7Part2(input string) {
	equations := strings.Split(input, "\n")
	re := regexp.MustCompile(`\d+`)
	sum := 0
	for _, equation := range equations {
		var err error
		nums := re.FindAllStringSubmatch(equation, -1)
		goal, err := strconv.Atoi(nums[0][0])
		if err != nil {
			panic(err)
		}
		values := make([]int, len(nums)-1)
		for i := 1; i < len(nums); i++ {
			val, err := strconv.Atoi(nums[i][0])
			if err != nil {
				panic(err)
			}
			values[i-1] = val
		}
		if isSolvable(values, goal, []string{"*", "+","||"}) {
			sum += goal
		}
	}
	fmt.Println(sum)
}

func isSolvable(values []int, goal int, ops []string) bool {
	if len(values) == 1 {
		return values[0] == goal
	}
	for _, op := range ops {
		var result int
		switch op {
		case "*":
			result = values[0] * values[1]
		case "+":
			result = values[0] + values[1]
		case "||":
			var err error
			str0, str1 := strconv.Itoa(values[0]), strconv.Itoa(values[1])
			result, err = strconv.Atoi(str0 + str1)
			if err != nil {
				panic(err)
			}
		}
		if result <= goal && isSolvable(append([]int{result}, values[2:]...), goal, ops) {
			return true
		}
	}
	return false
}
