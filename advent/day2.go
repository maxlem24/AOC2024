package advent

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2Part1(input string) {
	reports := strings.Split(input, "\n")
	count := 0
	for _, report := range reports[:len(reports)-1] {
		levels := strings.Split(report, " ")
		levelValues := make([]int, len(levels))
		var err error
		for i, val := range levels {
			levelValues[i], err = strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
		}
		if isReportValid(levelValues) {
			count++
		}
	}
	fmt.Println(count)
}

func Day2Part2(input string) {
	reports := strings.Split(input, "\n")
	count := 0
	for _, report := range reports[:len(reports)-1] {
		levels := strings.Split(report, " ")
		levelValues := make([]int, len(levels))
		var err error
		for i, val := range levels {
			levelValues[i], err = strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
		}
		if isReportValidWithDampener(levelValues) {
			count++
		}
	}
	fmt.Println(count)
}

func isReportValid(levels []int) bool {
	if len(levels) == 1 {
		return true
	}
	increasing := true
	if levels[1] < levels[0] {
		increasing = false
	}
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if !increasing {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isReportValidWithDampener(reportNum []int) bool {

	for i := 0; i < len(reportNum); i++ {
		isSafe := isReportSafeWithDeletion(reportNum, i)
		if isSafe {
			return true
		}
	}

	return false
}

func isReportSafeWithDeletion(report []int, deleteIndex int) bool {
	copyReport := make([]int, len(report))
	copy(copyReport, report)

	if deleteIndex == len(copyReport)-1 {
		copyReport = copyReport[:deleteIndex]
	} else {
		copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)
	}
	return isReportValid(copyReport)
}


