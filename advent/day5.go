package advent

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day5Part1(input string) {
	parts := strings.Split(input, "\n\n")
	rules := strings.Split(parts[0], "\n")
	prints := strings.Split(parts[1], "\n")
	numbers := regexp.MustCompile(`\d+`)
	afters := make(map[int][]int)
	for _, rule := range rules {
		var err error
		values := numbers.FindAllStringSubmatch(rule, -1)
		prec, err := strconv.Atoi(values[0][0])
		if err != nil {
			panic(err)
		}
		next, err := strconv.Atoi(values[1][0])
		if err != nil {
			panic(err)
		}
		if _, exist := afters[prec]; !exist {
			afters[prec] = []int{next}
		} else {
			afters[prec] = append(afters[prec], next)
		}
		if _, exist := afters[next]; !exist {
			afters[next] = make([]int, 0)
		}
	}
	manuals := make([][]int, 0)
	for _, print := range prints {
		pages := make([]int, 0)
		index := strings.Split(print, ",")
		for _, num := range index {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			pages = append(pages, val)
		}
		manuals = append(manuals, pages)
	}
	points := 0
	for _, manual := range manuals {
		isValid := true
		for curr := 0; isValid && curr < len(manual); curr++ {
			for prev := curr; isValid && prev >= 0; prev-- {
				if slices.Contains(afters[manual[curr]], manual[prev]) {
					isValid = false
				}
			}
		}
		if isValid {
			points += manual[len(manual)/2]
		}
	}
	fmt.Println(points)
}

func Day5Part2(input string) {
	parts := strings.Split(input, "\n\n")
	rules := strings.Split(parts[0], "\n")
	prints := strings.Split(parts[1], "\n")
	numbers := regexp.MustCompile(`\d+`)
	afters := make(map[int][]int)
	for _, rule := range rules {
		var err error
		values := numbers.FindAllStringSubmatch(rule, -1)
		prec, err := strconv.Atoi(values[0][0])
		if err != nil {
			panic(err)
		}
		next, err := strconv.Atoi(values[1][0])
		if err != nil {
			panic(err)
		}
		if _, exist := afters[prec]; !exist {
			afters[prec] = []int{next}
		} else {
			afters[prec] = append(afters[prec], next)
		}
		if _, exist := afters[next]; !exist {
			afters[next] = make([]int, 0)
		}
	}
	manuals := make([][]int, 0)
	for _, print := range prints {
		pages := make([]int, 0)
		index := strings.Split(print, ",")
		for _, num := range index {
			val, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			pages = append(pages, val)
		}
		manuals = append(manuals, pages)
	}
	points := 0
	for _, manual := range manuals {
		isValid := false
		isFixed := false
		for !isValid{
			isValid = true
			for curr := 0; curr < len(manual); curr++ {
				for prev := curr;  prev >= 0; prev-- {
					if slices.Contains(afters[manual[curr]], manual[prev]) {
						isValid = false
						manual[curr],manual[prev] = manual[prev], manual[curr]
						isFixed = true
					}
				}
			}
		}
		if isFixed {
			points += manual[len(manual)/2]
		}
	}
	fmt.Println(points)
}
