package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

func evalCycle(cycle, registerX int) int {
	importantCycles := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}
	if _, ok := importantCycles[cycle]; ok {
		return cycle * registerX
	}
	return 0
}

func evalA(lines []string) int {
	var signalStrength int
	registerX := 1
	cycle := 0

	for _, line := range lines {
		switch line[:4] {
		case "noop":
			cycle += 1
			signalStrength += evalCycle(cycle, registerX)
		case "addx":
			delta, _ := strconv.Atoi(strings.Split(line, " ")[1])
			cycle += 1
			signalStrength += evalCycle(cycle, registerX)
			cycle += 1
			signalStrength += evalCycle(cycle, registerX)
			registerX += delta
		}
	}

	return signalStrength
}

func checkCrtPosition(registerX, cycle int) string {
	if registerX <= cycle%40 && registerX+2 >= cycle%40 {
		return "#"
	} else {
		return "."
	}
}

func evalB(lines []string) string {
	var crt string
	registerX := 1
	cycle := 0

	for _, line := range lines {
		switch line[:4] {
		case "noop":
			cycle += 1
			crt += checkCrtPosition(registerX, cycle)
		case "addx":
			delta, _ := strconv.Atoi(strings.Split(line, " ")[1])
			cycle += 1
			crt += checkCrtPosition(registerX, cycle)
			cycle += 1
			crt += checkCrtPosition(registerX, cycle)
			registerX += delta
		}
	}
	return crt
}

func printB(crt string) {
	for i := 0; i < len(crt); i += 40 {
		fmt.Println(crt[i : i+40])
	}
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)

	resA := evalA(lines)
	resB := evalB(lines)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		printB(resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		printB(resB)
	}

}

func main() {
	day := 10
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
