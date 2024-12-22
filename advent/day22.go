package advent

import (
	"fmt"
	"strconv"
	"strings"
)

func Day22Part1(input string) {
	cache := make(map[int]int)
	sum := 0
	for _, monkey := range strings.Split(input, "\n") {
		secret, err := strconv.Atoi(monkey)
		if err != nil {
			panic(err)
		}
		for i := 0; i < 2000; i++ {
			if val, exist := cache[secret]; exist {
				secret = val
				continue
			}
			step1 := ((secret * 64) ^ secret) % 16777216
			step2 := ((step1 / 32) ^ step1) % 16777216
			step3 := ((step2 * 2048) ^ step2) % 16777216
			cache[secret] = step3
			secret = step3
		}
		sum += secret
	}
	fmt.Println(sum)
}

func Day22Part2(input string) {

	cache := make(map[int]int)
	buyers := strings.Split(input, "\n")
	sequences := make([]map[string]int, len(buyers))
	for index, monkey := range buyers {
		sequences[index] = make(map[string]int)
		secret, err := strconv.Atoi(monkey)
		if err != nil {
			panic(err)
		}
		sequence := "0,0,0,0"
		for i := 0; i < 2000; i++ {
			if val, exist := cache[secret]; exist {
				sequence += "," + strconv.Itoa(val%10-secret%10)
				if sequence[0] == '-' {
					sequence = sequence[3:]
				} else {
					sequence = sequence[2:]
				}
				if _, exist := sequences[index][sequence]; !exist && i >= 3 {
					sequences[index][sequence] = val % 10
				}
				secret = val
				continue
			}
			step1 := ((secret * 64) ^ secret) % 16777216
			step2 := ((step1 / 32) ^ step1) % 16777216
			step3 := ((step2 * 2048) ^ step2) % 16777216
			cache[secret] = step3
			sequence += "," + strconv.Itoa(step3%10-secret%10)
			if sequence[0] == '-' {
				sequence = sequence[3:]
			} else {
				sequence = sequence[2:]
			}
			if _, exist := sequences[index][sequence]; !exist && i >= 3 {
				sequences[index][sequence] = step3 % 10
			}
			secret = step3
		}
	}
	bestPrice := 0
	prices := make(map[string]int)
	for i := 0; i < len(buyers); i++ {
		for sequence := range sequences[i] {
			if _, exist := prices[sequence]; exist {
				continue
			}
			for j := 0; j < len(buyers); j++ {
				prices[sequence] += sequences[j][sequence]
			}
			if prices[sequence] > bestPrice {
				bestPrice = prices[sequence]
			}
		}
	}
	fmt.Println(bestPrice)
}
