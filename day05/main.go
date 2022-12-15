package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

func getInstructions(lines []string) []string {
	for i, line := range lines {
		if line == "" {
			return lines[i+1:]
		}
	}

	var emptyInstructions []string
	return emptyInstructions
}

func getCrate(line string, i int) string {
	pos := i*4 + 1
	if pos+1 > len(line) {
		return " "
	} else {
		return string(line[pos])
	}

}

func getReversedLines(lines []string) []string {
	reversedLines := make([]string, len(lines))
	copy(reversedLines, lines)
	for i, j := 0, len(reversedLines)-1; i < j; i, j = i+1, j-1 {
		reversedLines[i], reversedLines[j] = reversedLines[j], reversedLines[i]
	}
	return reversedLines
}

func reverseStacks(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func parseStartingStacks(lines []string) []string {
	var startingStacks []string
	numStacks := (len(lines[len(lines)-1]) + 2) / 4
	reversedLines := getReversedLines(lines[:len(lines)-1])
	for i := 0; i < numStacks; i++ {
		stack := ""
		for _, line := range reversedLines {
			crate := getCrate(line, i)
			if crate != " " {
				stack += crate
			}
		}
		startingStacks = append(startingStacks, stack)
	}
	return startingStacks
}

func getStartingStacks(lines []string) []string {
	for i, line := range lines {
		if line == "" {
			return parseStartingStacks(lines[:i])
		}
	}
	var emptyStartingStacks []string
	return emptyStartingStacks
}

func runInstruction(instruction string, stacks []string, reverse bool) []string {
	instructionParts := strings.Split(instruction, " ")
	numCrates, _ := strconv.Atoi(instructionParts[1])
	source, _ := strconv.Atoi(instructionParts[3])
	target, _ := strconv.Atoi(instructionParts[5])
	source -= 1
	target -= 1
	crates := stacks[source][len(stacks[source])-numCrates : len(stacks[source])]
	if reverse {
		crates = reverseStacks(crates)
	}
	stacks[target] += crates
	stacks[source] = stacks[source][:len(stacks[source])-numCrates]

	return stacks
}

func getMessage(stacks []string) string {
	message := ""
	for _, stack := range stacks {
		message += string(stack[len(stack)-1])
	}
	return message
}

func evalA(lines []string) string {
	instructions := getInstructions(lines)
	stacks := getStartingStacks(lines)
	for _, instruction := range instructions {
		stacks = runInstruction(instruction, stacks, true)
	}

	return getMessage(stacks)
}

func evalB(lines []string) string {
	instructions := getInstructions(lines)
	stacks := getStartingStacks(lines)
	for _, instruction := range instructions {
		stacks = runInstruction(instruction, stacks, false)
	}

	return getMessage(stacks)
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
	day := 5
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
