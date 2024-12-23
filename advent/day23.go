package advent

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func Day23Part1(input string) {
	graph := make(map[string][]string)
	for _, link := range strings.Split(input, "\n") {
		computer1 := link[:2]
		computer2 := link[3:]
		graph[computer1] = append(graph[computer1], computer2)
		graph[computer2] = append(graph[computer2], computer1)
	}
	sets := make(map[string]bool)
	count := 0
	for name, connected := range graph {
		if name[0] == 't' {
			for i := 0; i < len(connected); i++ {
				for j := i + 1; j < len(connected); j++ {
					for _, neighbourgh := range graph[connected[i]] {
						if neighbourgh == connected[j] {
							lan := []string{name, neighbourgh, connected[i]}
							sort.Strings(lan)
							trio := strings.Join(lan, "-")
							if _, exist := sets[trio]; !exist {
								sets[trio] = true
								count++
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func Day23Part2(input string) {
	graph := make(map[string][]string)
	for _, link := range strings.Split(input, "\n") {
		computer1 := link[:2]
		computer2 := link[3:]
		graph[computer1] = append(graph[computer1], computer2)
		graph[computer2] = append(graph[computer2], computer1)
	}
	P := make([]string, 0)
	for k := range graph {
		P = append(P, k)
	}
	completeGraph := BronKerbosch(graph, make([]string, 0), P, make([]string, 0))
	sort.Strings(completeGraph)
	lan := strings.Join(completeGraph, ",")
	fmt.Println(lan)
}

func BronKerbosch(G map[string][]string, R []string, P []string, X []string) []string {
	if len(P) == 0 && len(X) == 0 {
		return R
	}
	var pivot string
	if len(P) != 0 {
		pivot = P[0]
	} else {
		return []string{}
	}
	bestComplete := make([]string, 0)
	for _, vertex := range P {
		if slices.Contains(G[pivot], vertex) {
			continue
		}
		newR := make([]string, len(R)+1)
		copy(newR, R)
		newR[len(R)] = vertex
		newP := make([]string, 0)
		newX := make([]string, 0)
		for _, elem := range G[vertex] {
			if slices.Contains(P, elem) {
				newP = append(newP, elem)
			}
			if slices.Contains(X, elem) {
				newP = append(newX, elem)
			}
		}
		candidate := BronKerbosch(G, newR, newP, newX)
		if len(candidate) > len(bestComplete) {
			bestComplete = candidate
		}
		P = P[1:]
		X = append(X, vertex)
	}
	return bestComplete
}
