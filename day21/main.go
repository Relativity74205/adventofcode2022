package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

type Monkey interface {
	result() int
	numberKnown() bool
	setNumber(int)
}

type YellMonkey struct {
	name   string
	number int
}

func (y *YellMonkey) result() int {
	return y.number
}

func (y *YellMonkey) numberKnown() bool {
	return y.number != 0
}

func (y *YellMonkey) setNumber(n int) {
	y.number = n
}

type OpMonkey struct {
	name        string
	monkey1Name string
	monkey2Name string
	monkey1     Monkey
	monkey2     Monkey
	operator    string
	number      int
}

func (op *OpMonkey) result() int {
	if op.numberKnown() {
		return op.number
	}

	var result int
	switch op.operator {
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
	if op.monkey1.numberKnown() && op.monkey2.numberKnown() {
		op.number = result
	}

	return result
}

func (op *OpMonkey) numberKnown() bool {
	return op.number != 0
}

func (op *OpMonkey) setNumber(n int) {
	op.number = n
}

func parseInput(lines []string) map[string]Monkey {
	monkeys := make(map[string]Monkey)
	for _, line := range lines {
		lineParts := strings.Split(line, ": ")
		monkeyName := lineParts[0]
		operationsParts := strings.Split(lineParts[1], " ")

		if len(operationsParts) == 1 {
			number, _ := strconv.Atoi(operationsParts[0])
			monkeys[monkeyName] = &YellMonkey{name: monkeyName, number: number}
		} else {
			monkeys[monkeyName] = &OpMonkey{
				name:        monkeyName,
				monkey1Name: operationsParts[0],
				monkey1:     nil,
				monkey2Name: operationsParts[2],
				monkey2:     nil,
				operator:    operationsParts[1],
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

func getHumnNumber(parentMonkey *OpMonkey) int {
	var unfinishedMonkey, finishedMonkey Monkey
	if parentMonkey.monkey1.numberKnown() {
		finishedMonkey = parentMonkey.monkey1
		unfinishedMonkey = parentMonkey.monkey2
	} else {
		finishedMonkey = parentMonkey.monkey2
		unfinishedMonkey = parentMonkey.monkey1
	}

	switch parentMonkey.operator {
	case "+":
		unfinishedMonkey.setNumber(parentMonkey.result() - finishedMonkey.result())
	case "-":
		if parentMonkey.monkey1.numberKnown() {
			unfinishedMonkey.setNumber(finishedMonkey.result() - parentMonkey.result())
		} else {
			unfinishedMonkey.setNumber(finishedMonkey.result() + parentMonkey.result())
		}
	case "*":
		unfinishedMonkey.setNumber(parentMonkey.result() / finishedMonkey.result())
	case "/":
		if parentMonkey.monkey1.numberKnown() {
			unfinishedMonkey.setNumber(finishedMonkey.result() / parentMonkey.result())
		} else {
			unfinishedMonkey.setNumber(finishedMonkey.result() * parentMonkey.result())
		}
	}

	if mutableOpMonkey, ok := unfinishedMonkey.(*OpMonkey); ok {
		return getHumnNumber(mutableOpMonkey)
	} else {
		// found answer for B :)
		return unfinishedMonkey.result()
	}
}

func evalB(lines []string) int {
	monkeys := parseInput(lines)
	rootMonkey := monkeys["root"].(*OpMonkey)
	rootMonkey.operator = "="
	humn := monkeys["humn"].(*YellMonkey)
	humn.number = 0
	// to pre-calculate immutable nodes aka monkeys
	rootMonkey.result()

	var firstTopMonkey *OpMonkey
	if rootMonkey.monkey1.numberKnown() {
		firstTopMonkey = rootMonkey.monkey2.(*OpMonkey)
		firstTopMonkey.number = rootMonkey.monkey1.result()
	} else {
		firstTopMonkey = rootMonkey.monkey1.(*OpMonkey)
		firstTopMonkey.number = rootMonkey.monkey2.result()
	}

	return getHumnNumber(firstTopMonkey)
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
