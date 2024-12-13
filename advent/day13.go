package advent

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day13Part1(input string) {
	clawMachines := strings.Split(input, "\n\n")
	total := 0
	re := regexp.MustCompile(`\d+`)
	for _, machine := range clawMachines {
		values := re.FindAllStringSubmatch(machine, -1)
		xA, _ := strconv.Atoi(values[0][0])
		yA, _ := strconv.Atoi(values[1][0])
		xB, _ := strconv.Atoi(values[2][0])
		yB, _ := strconv.Atoi(values[3][0])
		xP, _ := strconv.Atoi(values[4][0])
		yP, _ := strconv.Atoi(values[5][0])
		minTokens := 500
		for tA := 0; tA <= 100 && 3*tA < minTokens; tA++ {
			remX, remY := xP-tA*xA, yP-tA*yA
			if remX%xB == 0 && remY%yB == 0 {
				coefX, coefY := remX/xB, remY/yB
				if coefX == coefY && coefX <= 100 {
					minTokens = min(minTokens, 3*tA+coefX)
				}
			}
		}
		if minTokens != 500 {
			total += minTokens
		}
	}
	fmt.Println(total)
}

func Day13Part2(input string) {
	clawMachines := strings.Split(input, "\n\n")
	total := 0
	convError := 10000000000000
	re := regexp.MustCompile(`\d+`)
	for _, machine := range clawMachines {
		values := re.FindAllStringSubmatch(machine, -1)
		xA, _ := strconv.Atoi(values[0][0])
		yA, _ := strconv.Atoi(values[1][0])
		xB, _ := strconv.Atoi(values[2][0])
		yB, _ := strconv.Atoi(values[3][0])
		xP, _ := strconv.Atoi(values[4][0])
		yP, _ := strconv.Atoi(values[5][0])
		xP, yP = xP+convError, yP+convError
		rem := (xA*yP - xP*yA) % (yB*xA - xB*yA)
		if rem != 0 {
			continue
		}
		tB := (xA*yP - xP*yA) / (yB*xA - xB*yA)
		tA := (xP - tB*xB) / xA
		total += 3*tA + tB
	}
	fmt.Println(total)
}
