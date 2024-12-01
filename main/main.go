package main

import (
	"aoc2024/advent"
	"fmt"
	"os"
)

func openFile(path string) string {
	buffer, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(buffer)
}

func main() {
	var day int
	fmt.Println("Which day ?")
	_, err := fmt.Scanf("%d", &day)
	if err != nil {
		panic(err)
	}
	var file string
	if len(os.Args) != 1 && os.Args[1] == "example" {
		file = openFile(fmt.Sprintf("../files/day%d_example.txt", day))
	} else {
		file = openFile(fmt.Sprintf("../files/day%d.txt", day))
	}
	switch day {
	case 1:
		advent.Day1Part1(file)
		advent.Day1Part2(file)
	default:
		fmt.Println("Code not available")
	}
}
