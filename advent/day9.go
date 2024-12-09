package advent

import (
	"fmt"
	"strconv"
)

type space struct {
	startIndex   int
	currentIndex int
	len          int
}

func Day9Part1(input string) {
	if len(input)%2 == 0 {
		input = input[:len(input)-1]
	}
	checksum := 0
	firstIndex, lastIndex := 0, len(input)/2
	index := 0
	firstValue, err := strconv.Atoi(input[2*firstIndex : 2*firstIndex+1])
	if err != nil {
		panic(err)
	}
	lastValue, err := strconv.Atoi(input[2*lastIndex : 2*lastIndex+1])
	if err != nil {
		panic(err)
	}
	spaceAvailable := 0
	for firstIndex < lastIndex {
		if spaceAvailable == 0 {
			for firstValue > 0 {
				checksum += index * firstIndex
				index++
				firstValue--
			}
			firstIndex++
			spaceAvailable, err = strconv.Atoi(input[2*firstIndex-1 : 2*firstIndex])
			if err != nil {
				panic(err)
			}
			firstValue, err = strconv.Atoi(input[2*firstIndex : 2*firstIndex+1])
			if err != nil {
				panic(err)
			}
		} else {
			for spaceAvailable > 0 && lastValue > 0 {
				checksum += index * lastIndex
				index++
				lastValue--
				spaceAvailable--
			}
			if lastValue == 0 {
				lastIndex--
				lastValue, err = strconv.Atoi(input[2*lastIndex : 2*lastIndex+1])
				if err != nil {
					panic(err)
				}
			}
		}
	}
	for lastValue > 0 {
		checksum += index * lastIndex
		index++
		lastValue--
	}
	fmt.Println(checksum)
}

func Day9Part2(input string) {
	if len(input)%2 == 0 {
		input = input[:len(input)-1]
	}
	spaces := make([]space, len(input)/2)
	initVal, err := strconv.Atoi(input[0:1])
	if err != nil {
		panic(err)
	}
	checksum := 0
	index := initVal
	for i := 0; i < len(input)/2; i++ {
		val, err := strconv.Atoi(input[2*i+1 : 2*i+3])
		if err != nil {
			panic(err)
		}
		spaces[i] = space{len: val / 10, startIndex: index, currentIndex: index}
		index += val % 10
		index += val / 10
	}
	for lastIndex := len(input) / 2; lastIndex > 0; lastIndex-- {
		count, err := strconv.Atoi(input[2*lastIndex : 2*lastIndex+1])
		if err != nil {
			panic(err)
		}
		newPlace := false
		for i, space := range spaces {
			if space.len >= count {
				for cpt := 0; cpt < count; cpt++ {
					checksum += (space.currentIndex + cpt) * lastIndex
				}
				spaces[i].len -= count
				spaces[i].currentIndex += count
				newPlace = true
				break
			}
		}
		if !newPlace {
			for cpt := 0; cpt < count; cpt++ {
				index--
				checksum += index * lastIndex
			}
		}
		index = spaces[len(spaces)-1].startIndex
		spaces = spaces[:len(spaces)-1]
	}
	fmt.Println(checksum)
}
