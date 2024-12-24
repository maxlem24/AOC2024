package advent

import (
	"fmt"
	"sort"
	"strings"
)

type logicGate struct {
	a   string
	op  string
	b   string
	out string
}

func Day24Part1(input string) {
	parts := strings.Split(input, "\n\n")
	wires := make(map[string]bool)
	for _, initialValue := range strings.Split(parts[0], "\n") {
		key := initialValue[:3]
		if initialValue[5] == '1' {
			wires[key] = true
		} else {
			wires[key] = false
		}
	}
	gates := make([]string, 0)

	gates = append(gates, strings.Split(parts[1], "\n")...)

	for len(gates) > 0 {
		curr := gates[0]
		gates = gates[1:]
		elem := strings.Split(curr, " ")
		if _, exist := wires[elem[0]]; !exist {
			gates = append(gates, curr)
			continue
		}
		if _, exist := wires[elem[2]]; !exist {
			gates = append(gates, curr)
			continue
		}
		switch elem[1] {
		case "AND":
			wires[elem[4]] = wires[elem[0]] && wires[elem[2]]
		case "OR":
			wires[elem[4]] = wires[elem[0]] || wires[elem[2]]
		case "XOR":
			wires[elem[4]] = wires[elem[0]] != wires[elem[2]]
		}
	}
	total := 0
	zCount := 0
	power2 := 1
	for val, exist := wires[fmt.Sprintf("z%02d", zCount)]; exist; val, exist = wires[fmt.Sprintf("z%02d", zCount)] {
		if val {
			total += power2
		}
		power2 *= 2
		zCount++
	}
	fmt.Println(total)
}

func Day24Part2(input string) {

	parts := strings.Split(input, "\n\n")
	wires := make(map[string]bool)
	for _, initialValue := range strings.Split(parts[0], "\n") {
		key := initialValue[:3]
		if initialValue[5] == '1' {
			wires[key] = true
		} else {
			wires[key] = false
		}
	}
	inputBits := len(wires) / 2
	gates := make([]logicGate, 0)
	for _, gate := range strings.Split(parts[1], "\n") {
		elem := strings.Split(gate, " ")
		gates = append(gates, logicGate{a: elem[0], op: elem[1], b: elem[2], out: elem[4]})
	}

	/**
	 * FULL ADDER
	 * (first bits aren't a full adder)
	 *
	 * A    XOR B    -> VAL0
	 * A    AND B    -> VAL1
	 * VAL0 AND CIN  -> VAL2
	 * VAL0 XOR CIN  -> SUM
	 * VAL1 OR  VAL2 -> COUT
	 */

	cIn := ""
	swapped := make([]string, 0)
	for i := 0; i < inputBits; i++ {
		index := fmt.Sprintf("%02d", i)
		val0 := findGates("x"+index, "y"+index, "XOR", gates)
		val1 := findGates("x"+index, "y"+index, "AND", gates)
		if cIn == "" {
			cIn = val1
			continue
		}
		val2 := findGates(cIn, val0, "AND", gates)
		if val2 == "" { // Error Gate
			val0, val1 = val1, val0
			swapped = append(swapped, val0, val1)
			val2 = findGates(cIn, val0, "AND", gates)
		}

		sum := findGates(val0, cIn, "XOR", gates)

		if len(val0) > 0 && val0[0] == 'z' {
			sum, val0 = val0, sum
			swapped = append(swapped, val0, sum)
		}

		if len(val1) > 0 && val1[0] == 'z' {
			sum, val1 = val1, sum
			swapped = append(swapped, val1, sum)
		}

		if len(val2) > 0 && val2[0] == 'z' {
			sum, val2 = val2, sum
			swapped = append(swapped, val2, sum)
		}

		cOut := findGates(val1, val2, "OR", gates)

		if len(cOut) > 0 && cOut[0] == 'z' && cOut != fmt.Sprintf("z%02d", inputBits) {
			sum, cOut = cOut, sum
			swapped = append(swapped, cOut, sum)
		}

		cIn = cOut
	}
	sort.Strings(swapped)
	fmt.Println(strings.Join(swapped,","))
}

func findGates(a, b, op string, gates []logicGate) string {
	for _, gate := range gates {
		if gate.a == a && gate.b == b && gate.op == op {
			return gate.out
		}
		if gate.b == a && gate.a == b && gate.op == op {
			return gate.out
		}
	}
	return ""
}
