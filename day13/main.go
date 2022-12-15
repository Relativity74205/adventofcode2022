package main

import (
	"AoC2022"
	"fmt"
)

type inputPair struct {
	signal1, signal2 string
}

func validInputPair(inputPair inputPair) bool {

}

func evalA(inputPairs []inputPair) int {
	var sumIndex int

	for i, inputPair := range inputPairs {
		if validInputPair(inputPair) {
			sumIndex += i
		}
	}

	return sumIndex
}

func evalB(lines []string) int {

	return 0
}

func getInputPairs(lines []string) []inputPair {
	var pairs []inputPair
	for i := 0; i < (len(lines)+1)/3; i += 3 {
		pairs = append(pairs, inputPair{lines[0], lines[1]})
	}

	return pairs
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	inputPairs := getInputPairs(lines)

	resA := evalA(inputPairs)
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
	day := 13
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
