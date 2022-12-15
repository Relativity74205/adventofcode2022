package main

import (
	"AoC2022"
	"fmt"
)

var myMove = map[string]int{
	"X": 1, //Rock A
	"Y": 2, //Paper B
	"Z": 3, //Scissors C
}

var outcome = map[string]int{
	"A X": 3, //Rock Rock
	"A Y": 6, //Rock Paper
	"A Z": 0, //Rock Scissors
	"B X": 0, //Paper Rock
	"B Y": 3, //Paper Paper
	"B Z": 6, //Paper Scissors
	"C X": 6, //Scissors Rock
	"C Y": 0, //Scissors Paper
	"C Z": 3, //Scissors Scissors
}

var translate = map[string]string{
	"A X": "A Z", //Rock Scissors
	"A Y": "A X", //Rock Rock
	"A Z": "A Y", //Rock Paper
	"B X": "B X", //Paper Rock
	"B Y": "B Y", //Paper Paper
	"B Z": "B Z", //Paper Scissors
	"C X": "C Y", //Scissors Paper
	"C Y": "C Z", //Scissors Scissors
	"C Z": "C X", //Scissors Rock
}

func evalA(lines []string) int {
	score := 0
	for _, line := range lines {
		score += outcome[line]
		score += myMove[string(line[2])]
	}

	return score
}

func evalB(lines []string) int {
	score := 0
	for _, line := range lines {
		transformed := translate[line]
		score += outcome[transformed]
		score += myMove[string(transformed[2])]
	}

	return score
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
	day := 2
	//debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	//filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	//eval(filenameDebug, true)
	eval(filename, false)
}
