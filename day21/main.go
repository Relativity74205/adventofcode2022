package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

type Monkey interface {
	result() int
	isImmutable() bool
}

type YellMonkey struct {
	name      string
	number    int
	immutable bool
}

func (y YellMonkey) result() int {
	return y.number
}

func (y YellMonkey) isImmutable() bool {
	return y.immutable
}

type OpMonkey struct {
	name         string
	monkey1Name  string
	monkey2Name  string
	monkey1      Monkey
	monkey2      Monkey
	operation    string
	cachedNumber int
	immutable    bool
}

func (op *OpMonkey) result() int {
	if op.immutable {
		return op.cachedNumber
	}

	var result int
	switch op.operation {
	case "+":
		result = op.monkey1.result() + op.monkey2.result()
	case "-":
		result = op.monkey1.result() - op.monkey2.result()
	case "*":
		result = op.monkey1.result() * op.monkey2.result()
	case "/":
		result = op.monkey1.result() / op.monkey2.result()
	default:
		if op.monkey1.result() == op.monkey2.result() {
			result = 1
		} else {
			result = 0
		}
	}
	if op.monkey1.isImmutable() && op.monkey2.isImmutable() {
		op.immutable = true
		op.cachedNumber = result
	}

	return result
}

func (op OpMonkey) isImmutable() bool {
	return op.immutable
}

func parseInput(lines []string) map[string]Monkey {
	monkeys := make(map[string]Monkey)
	for _, line := range lines {
		lineParts := strings.Split(line, ": ")
		monkeyName := lineParts[0]
		operationsParts := strings.Split(lineParts[1], " ")

		if len(operationsParts) == 1 {
			number, _ := strconv.Atoi(operationsParts[0])
			monkeys[monkeyName] = &YellMonkey{name: monkeyName, number: number, immutable: true}
		} else {
			monkeys[monkeyName] = &OpMonkey{
				name:        monkeyName,
				monkey1Name: operationsParts[0],
				monkey1:     nil,
				monkey2Name: operationsParts[2],
				monkey2:     nil,
				operation:   operationsParts[1],
				immutable:   false,
			}
		}
	}

	for _, monkey := range monkeys {
		opMonkey, isOpMonkey := monkey.(*OpMonkey)
		if isOpMonkey {
			opMonkey.monkey1 = monkeys[opMonkey.monkey1Name]
			opMonkey.monkey2 = monkeys[opMonkey.monkey2Name]
		}
	}

	return monkeys
}

func evalA(lines []string) int {
	monkeys := parseInput(lines)

	return monkeys["root"].result()
}

func evalB(lines []string) int {
	monkeys := parseInput(lines)
	rootMonkey := monkeys["root"].(*OpMonkey)
	rootMonkey.operation = "="
	for myNumber := -1000000000; myNumber <= 1000000000; myNumber += 1 {
		//for myNumber := -10000000; myNumber <= 10000000; myNumber += 1 {
		humn := monkeys["humn"].(*YellMonkey)
		humn.number = myNumber
		humn.immutable = false

		if monkeys["root"].result() == 1 {
			return myNumber
		}
	}

	return -1
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)

	//resA := evalA(lines)
	resB := evalB(lines)
	if debug {
		//fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		//fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 21
	//debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	//filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	//eval(filenameDebug, true)
	eval(filename, false)
}
