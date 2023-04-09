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
	getOperator() string
	setNumber(int)
	setMutable(bool)
}

type YellMonkey struct {
	name      string
	number    int
	immutable bool
}

func (y *YellMonkey) result() int {
	return y.number
}

func (y *YellMonkey) isImmutable() bool {
	return y.immutable
}

func (y *YellMonkey) getOperator() string {
	return ""
}

func (y *YellMonkey) setNumber(n int) {
	y.number = n
}

func (y *YellMonkey) setMutable(m bool) {
	y.immutable = m
}

type OpMonkey struct {
	name        string
	monkey1Name string
	monkey2Name string
	monkey1     Monkey
	monkey2     Monkey
	operator    string
	number      int
	immutable   bool
}

func (op *OpMonkey) result() int {
	if op.immutable {
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
	if op.monkey1.isImmutable() && op.monkey2.isImmutable() {
		op.immutable = true
		op.number = result
	}

	return result
}

func (op *OpMonkey) isImmutable() bool {
	return op.immutable
}

func (op *OpMonkey) getOperator() string {
	return op.operator
}

func (op *OpMonkey) setNumber(n int) {
	op.number = n
}

func (op *OpMonkey) setMutable(m bool) {
	op.immutable = m
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
				operator:    operationsParts[1],
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

func getHumnNumber(monkey *OpMonkey) int {
	if monkey.name == "humn" {
		return monkey.result()
	}
	var mutableMonkey, immutableMonkey Monkey
	if monkey.monkey1.isImmutable() {
		immutableMonkey = monkey.monkey1
		mutableMonkey = monkey.monkey2
	} else {
		immutableMonkey = monkey.monkey2
		mutableMonkey = monkey.monkey1
	}

	switch monkey.getOperator() {
	case "+":
		mutableMonkey.setNumber(monkey.result() - immutableMonkey.result())
	case "-":
		if monkey.monkey1.isImmutable() {
			mutableMonkey.setNumber(monkey.monkey1.result() - monkey.result())
		} else {
			mutableMonkey.setNumber(monkey.monkey2.result() + monkey.result())
		}
	case "*":
		mutableMonkey.setNumber(monkey.result() / immutableMonkey.result())
	case "/":
		if monkey.monkey1.isImmutable() {
			mutableMonkey.setNumber(monkey.monkey1.result() / monkey.result())
		} else {
			mutableMonkey.setNumber(monkey.monkey2.result() * monkey.result())
		}
	}
	mutableMonkey.setMutable(true)
	if mutableOpMonkey, ok := mutableMonkey.(*OpMonkey); ok {
		return getHumnNumber(mutableOpMonkey)
	} else {
		return mutableMonkey.result()
	}
}

func evalB(lines []string) int {
	monkeys := parseInput(lines)
	rootMonkey := monkeys["root"].(*OpMonkey)
	rootMonkey.operator = "="
	humn := monkeys["humn"].(*YellMonkey)
	humn.number = 0
	humn.immutable = false
	// to pre-calculate immutable nodes aka monkeys
	rootMonkey.result()

	var topMonkey *OpMonkey
	if rootMonkey.monkey1.isImmutable() {
		topMonkey = rootMonkey.monkey2.(*OpMonkey)
		topMonkey.number = rootMonkey.monkey1.result()
	} else {
		topMonkey = rootMonkey.monkey1.(*OpMonkey)
		topMonkey.number = rootMonkey.monkey2.result()
	}
	topMonkey.immutable = true

	return getHumnNumber(topMonkey)
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
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
