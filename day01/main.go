package main

import (
	"AoC2022"
	"fmt"
	"strconv"
)

func getMaxElf(elves map[int]int) int {
	maxElfId := 0
	for elfId, sum := range elves {
		if sum > elves[maxElfId] {
			maxElfId = elfId
		}
	}
	return maxElfId
}

func getElves(lines []string) map[int]int {
	elves := make(map[int]int)

	elfId := 0
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			elfId += 1
			continue
		}

		elves[elfId] += number
	}

	return elves
}

func evalA(elves map[int]int) int {
	return elves[getMaxElf(elves)]
}

func evalB(elves map[int]int) int {
	calories := 0
	for i := 0; i < 3; i++ {
		maxElfId := getMaxElf(elves)
		calories += elves[maxElfId]
		delete(elves, maxElfId)
	}

	return calories
}

func eval(filename string, debug bool) {
	lines := AoC2022.ReadFile(filename)
	elves := getElves(lines)

	resA := evalA(elves)
	resB := evalB(elves)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 1
	//debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	//filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	//eval(filenameDebug, true)
	eval(filename, false)
}
