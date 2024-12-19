package advent

import (
	"fmt"
	"strings"
)

func Day19Part1(input string) {
	parts := strings.Split(input, "\n\n")
	towels := make(map[byte][]string)
	var validPattern func(string) bool
	validPattern = func(pattern string) bool {
		if pattern == "" {
			return true
		}
		for _, towel := range towels[pattern[0]] {
			if len(towel) <= len(pattern) && towel == pattern[:len(towel)] {
				if validPattern(pattern[len(towel):]) {
					return true
				}
			}
		}
		return false
	}

	for _, towel := range strings.Split(parts[0], ", ") {
		towels[towel[0]] = append(towels[towel[0]], towel)
	}
	count := 0
	for _, pattern := range strings.Split(parts[1], "\n") {
		if validPattern(pattern) {
			count++
		}
	}
	fmt.Println(count)
}

func Day19Part2(input string) {
	parts := strings.Split(input, "\n\n")
	towels := make(map[byte][]string)
	memory := make(map[string]int)
	var validPattern func(string) int
	validPattern = func(pattern string) int {
		if pattern == "" {
			return 1
		}
		if val, exist := memory[pattern]; exist {
			return val
		}
		count := 0
		for _, towel := range towels[pattern[0]] {
			if len(towel) <= len(pattern) && towel == pattern[:len(towel)] {
				count += validPattern(pattern[len(towel):])
			}
		}
		memory[pattern] = count
		return count
	}

	for _, towel := range strings.Split(parts[0], ", ") {
		towels[towel[0]] = append(towels[towel[0]], towel)
	}
	count := 0
	for _, pattern := range strings.Split(parts[1], "\n") {
		count += validPattern(pattern)
	}
	fmt.Println(count)
}
