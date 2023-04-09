package main

import (
	"AoC2022"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type monkey interface {
	result(monkeyMap map[string]monkey) int
}

type yellMonkey struct {
	number int
}

func (y *yellMonkey) result(_ map[string]monkey) int {
	return y.number
}

type opMonkey struct {
	monkey1   string
	monkey2   string
	operation string
}

func (op *opMonkey) result(monkeyMap map[string]monkey) int {
	switch op.operation {
	case "+":
		return monkeyMap[op.monkey1].result(monkeyMap) + monkeyMap[op.monkey2].result(monkeyMap)
	case "-":
		return monkeyMap[op.monkey1].result(monkeyMap) - monkeyMap[op.monkey2].result(monkeyMap)
	case "*":
		return monkeyMap[op.monkey1].result(monkeyMap) * monkeyMap[op.monkey2].result(monkeyMap)
	case "/":
		return monkeyMap[op.monkey1].result(monkeyMap) / monkeyMap[op.monkey2].result(monkeyMap)
	default:
		return 0
	}
}

func parseInput(lines []string) map[string]monkey {
	monkeys := make(map[string]monkey)
	for _, line := range lines {
		lineParts := strings.Split(line, ": ")
		monkeyName := lineParts[0]
		pattern := regexp.MustCompile(`[\+\-\*/]`)
		operationsParts := pattern.Split(lineParts[1], -1)
		if len(operationsParts) == 1 {
			number, _ := strconv.Atoi(operationsParts[0])
			monkeys[monkeyName] = &yellMonkey{number: number}
		} else {
			monkeys[monkeyName] = &opMonkey{
				monkey1:   operationsParts[0],
				monkey2:   operationsParts[2],
				operation: operationsParts[1],
			}
		}
	}

	return monkeys
}

func evalA(lines []string) int {
	monkeys := parseInput(lines)

	return 0
}

func evalB(lines []string) int {

	return 0
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)

	resA := evalA(lines)
	resB := evalB(lines)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 21
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
