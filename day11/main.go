package main

import (
	"AoC2022"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items            []big.Int
	operation        string
	operationNumber  *big.Int
	divisor          big.Int
	trueTarget       int
	falseTarget      int
	countInspections int
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey
	var operationNumber int
	for i := 0; i < (len(lines)+1)/7; i++ {
		line2parts := strings.Split(lines[i*7+1], ": ")
		line4parts := strings.Split(lines[i*7+3], " ")
		line5parts := strings.Split(lines[i*7+4], " ")
		line6parts := strings.Split(lines[i*7+5], " ")
		items, _ := AoC2022.StringSliceToBigIntSlice(strings.Split(line2parts[1], ", "))
		operation := strings.Split(lines[i*7+2], " = ")[1]
		divisor, _ := strconv.Atoi(line4parts[len(line4parts)-1])
		trueTarget, _ := strconv.Atoi(line5parts[len(line5parts)-1])
		falseTarget, _ := strconv.Atoi(line6parts[len(line6parts)-1])
		operationParts := strings.Split(operation, " ")
		if len(operationParts) == 3 {
			operationNumber, _ = strconv.Atoi(operationParts[2])
		}

		var bigIntDivisor big.Int
		bigIntDivisor.SetInt64(int64(divisor))
		monkey := Monkey{
			items:           items,
			operation:       operation,
			operationNumber: big.NewInt(int64(operationNumber)),
			divisor:         bigIntDivisor,
			trueTarget:      trueTarget,
			falseTarget:     falseTarget,
		}
		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func calcNewWeight(oldWeight int, operation string, divideByThree bool) int {
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

	if divideByThree {
		return int(math.Floor(float64(newWeight / 3)))
	} else {
		return newWeight
	}

}

func calcNewWeightBigInt(oldWeight *big.Int, operation string, operationNumber *big.Int) *big.Int {
	var newWeight *big.Int
	switch parts := strings.Split(operation, " "); {
	case parts[1] == "+":
		newWeight = big.NewInt(0).Add(oldWeight, operationNumber)
	case parts[1] == "*" && parts[2] == "old":
		newWeight = big.NewInt(0).Mul(oldWeight, oldWeight)
	case parts[1] == "*":
		newWeight = big.NewInt(0).Mul(oldWeight, operationNumber)
	}

	return newWeight
}

func evalA(monkeyInput []Monkey) int {
	monkeys := make([]Monkey, len(monkeyInput))
	copy(monkeys, monkeyInput)

	totalRounds := 20

	for round := 0; round < totalRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, itemWeight := range monkeys[i].items {
				monkeys[i].countInspections += 1
				newWeight := big.NewInt(int64(calcNewWeight(int(itemWeight.Int64()), monkeys[i].operation, true)))

				mod := big.NewInt(0).Mod(newWeight, &monkeys[i].divisor)
				if mod.Cmp(big.NewInt(0)) == 0 {
					monkeys[monkeys[i].trueTarget].items = append(monkeys[monkeys[i].trueTarget].items, *newWeight)
				} else {
					monkeys[monkeys[i].falseTarget].items = append(monkeys[monkeys[i].falseTarget].items, *newWeight)
				}
			}
			monkeys[i].items = make([]big.Int, 0)
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

	totalRounds := 20

	for round := 0; round < totalRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			for _, itemWeight := range monkeys[i].items {
				monkeys[i].countInspections += 1
				newWeight := calcNewWeightBigInt(&itemWeight, monkeys[i].operation, monkeys[i].operationNumber)

				mod := big.NewInt(0).Mod(newWeight, &monkeys[i].divisor)
				if mod.Cmp(big.NewInt(0)) == 0 {
					monkeys[monkeys[i].trueTarget].items = append(monkeys[monkeys[i].trueTarget].items, *newWeight)
				} else {
					monkeys[monkeys[i].falseTarget].items = append(monkeys[monkeys[i].falseTarget].items, *newWeight)
				}
			}
			monkeys[i].items = make([]big.Int, 0)
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
