package advent

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func Day11Part1(input string) {
	re := regexp.MustCompile(`\d+`)
	values := re.FindAllStringSubmatch(input, -1)
	currentVals := make([]int, len(values))
	for i, value := range values {
		intVal, err := strconv.Atoi(value[0])
		if err != nil {
			panic(err)
		}
		currentVals[i] = intVal
	}
	for i := 0; i < 25; i++ {
		nextVals := make([]int, 0)
		for _, val := range currentVals {
			if val == 0 {
				nextVals = append(nextVals, 1)
			} else if valLength := len(strconv.Itoa(val)); valLength%2 == 0 {
				power := int(math.Pow10(valLength / 2))
				nextVals = append(nextVals, val/power)
				nextVals = append(nextVals, val%power)
			} else {
				nextVals = append(nextVals, val*2024)
			}
		}
		currentVals = nextVals

	}
	fmt.Println(len(currentVals))
}

func Day11Part2(input string) {
	re := regexp.MustCompile(`\d+`)
	values := re.FindAllStringSubmatch(input, -1)
	currentVals := make(map[int]int)
	for _, value := range values {
		intVal, err := strconv.Atoi(value[0])
		if err != nil {
			panic(err)
		}
		currentVals[intVal] = 1
	}
	for i := 0; i < 75; i++ {
		nextVals := make(map[int]int, 0)
		for key, val := range currentVals {
			if key == 0 {
				nextVals[1] += val
			} else if keyLength := len(strconv.Itoa(key)); keyLength%2 == 0 {
				power := int(math.Pow10(keyLength / 2))
				nextVals[key%power] += val
				nextVals[key/power] += val
			} else {
				nextVals[key*2024] += val
			}
		}
		currentVals = nextVals

	}
	count := 0
	for _,v := range currentVals{
		count+=v
	}
	fmt.Println(count)
}
