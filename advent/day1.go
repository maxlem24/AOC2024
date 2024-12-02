package advent

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day1Part1(input string) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	n := len(lines)
	list1 := make([]int, n)
	list2 := make([]int, n)
	for i, line := range lines {
		numbers := strings.Split(line, "   ")
		var err error
		list1[i], err = strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		list2[i], err = strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
	}
	slices.Sort(list1)
	slices.Sort(list2)
	sum := 0
	for i := 0; i < n; i++ {
		sum += diffAbs(list1[i], list2[i])
	}
	fmt.Println(sum)
}

func Day1Part2(input string) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	n := len(lines)
	list1 := make([]int, n)
	countOcc := make(map[int]int, n)
	for i, line := range lines {
		numbers := strings.Split(line, "   ")
		var err error
		list1[i], err = strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		val, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		if _, exist := countOcc[val]; exist {
			countOcc[val]++
		} else {
			countOcc[val] = 1
		}
	}
	total := 0
	for i := 0; i < n; i++ {
		total += list1[i] * countOcc[list1[i]]
	}
	fmt.Println(total)
}

func diffAbs(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}
