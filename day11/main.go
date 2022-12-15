package main

import (
	"AoC2022"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items            []int
	operation        string
	divisor          int
	trueTarget       int
	falseTarget      int
	countInspections int
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey
	for i := 0; i < (len(lines)+1)/7; i++ {
		line2parts := strings.Split(lines[i*7+1], ": ")
		line4parts := strings.Split(lines[i*7+3], " ")
		line5parts := strings.Split(lines[i*7+4], " ")
		line6parts := strings.Split(lines[i*7+5], " ")
		items, _ := AoC2022.StringSliceToIntSlice(strings.Split(line2parts[1], ", "))
		operation := strings.Split(lines[i*7+2], " = ")[1]
		divisor, _ := strconv.Atoi(line4parts[len(line4parts)-1])
		trueTarget, _ := strconv.Atoi(line5parts[len(line5parts)-1])
		falseTarget, _ := strconv.Atoi(line6parts[len(line6parts)-1])

		monkey := Monkey{
			items:       items,
			operation:   operation,
			divisor:     divisor,
			trueTarget:  trueTarget,
			falseTarget: falseTarget,
		}
		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func calcNewWeight(oldWeight int, operation string) int {
	var newWeight int
	switch parts := strings.Split(operation, " "); {
	case parts[1] == "+":
		summand, _ := strconv.Atoi(parts[2])
		newWeight = oldWeight + summand
	case parts[1] == "*" && parts[2] == "old":
		newWeight = oldWeight * oldWeight
	case parts[1] == "*":
		factor, _ := strconv.Atoi(parts[2])
		newWeight = oldWeight * factor
	}
	return newWeight
}

func calcNewWeightA(oldWeight int, operation string) int {
	newWeight := calcNewWeight(oldWeight, operation)

	return int(math.Floor(float64(newWeight / 3)))
}

func calcNewWeightB(oldWeight int, operation string, moduloProduct int) int {
	newWeight := calcNewWeight(oldWeight, operation)

	return newWeight % moduloProduct
}

func evalA(monkeyInput []Monkey) int {
	monkeys := make([]Monkey, len(monkeyInput))
	copy(monkeys, monkeyInput)

	totalRounds := 20

	for round := 0; round < totalRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, itemWeight := range monkeys[i].items {
				monkeys[i].countInspections += 1
				newWeight := calcNewWeightA(itemWeight, monkeys[i].operation)

				if newWeight%monkeys[i].divisor == 0 {
					monkeys[monkeys[i].trueTarget].items = append(monkeys[monkeys[i].trueTarget].items, newWeight)
				} else {
					monkeys[monkeys[i].falseTarget].items = append(monkeys[monkeys[i].falseTarget].items, newWeight)
				}
			}
			monkeys[i].items = make([]int, 0)
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].countInspections > monkeys[j].countInspections
	})

	return monkeys[0].countInspections * monkeys[1].countInspections
}

func evalB(monkeyInput []Monkey) int {
	monkeys := make([]Monkey, len(monkeyInput))
	copy(monkeys, monkeyInput)
	moduloProduct := 1
	for _, monkey := range monkeys {
		moduloProduct *= monkey.divisor
	}

	totalRounds := 10000

	for round := 0; round < totalRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, itemWeight := range monkeys[i].items {
				monkeys[i].countInspections += 1
				newWeight := calcNewWeightB(itemWeight, monkeys[i].operation, moduloProduct)

				if newWeight%monkeys[i].divisor == 0 {
					monkeys[monkeys[i].trueTarget].items = append(monkeys[monkeys[i].trueTarget].items, newWeight)
				} else {
					monkeys[monkeys[i].falseTarget].items = append(monkeys[monkeys[i].falseTarget].items, newWeight)
				}
			}
			monkeys[i].items = make([]int, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].countInspections > monkeys[j].countInspections
	})

	return monkeys[0].countInspections * monkeys[1].countInspections
}

func eval(filename string, debug bool) {
	lines := AoC2022.ReadFile(filename)
	monkeys := parseMonkeys(lines)

	resA := evalA(monkeys)
	resB := evalB(monkeys)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 11
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
