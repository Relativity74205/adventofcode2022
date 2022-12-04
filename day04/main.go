package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

func getStartStop(assignment string) (int, int) {
	startStop := strings.Split(assignment, "-")
	start, _ := strconv.Atoi(startStop[0])
	stop, _ := strconv.Atoi(startStop[1])
	return start, stop
}

func evalA(lines []string) int {
	counter := 0
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		start1, stop1 := getStartStop(assignments[0])
		start2, stop2 := getStartStop(assignments[1])
		if start1 <= start2 && stop1 >= stop2 || start1 >= start2 && stop1 <= stop2 {
			counter += 1
		}
	}

	return counter
}

func evalB(lines []string) int {
	counter := 0
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		start1, stop1 := getStartStop(assignments[0])
		start2, stop2 := getStartStop(assignments[1])
		if !(stop1 < start2 || stop2 < start1) {
			counter += 1
		}
	}

	return counter
}

func eval(filename string, debug bool) {
	lines := AoC2022.ReadFile(filename)

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
	day := 4
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
