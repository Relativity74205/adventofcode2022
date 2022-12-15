package main

import (
	"AoC2022"
	"fmt"
)

func checkForMarker(s string) bool {
	set := make(map[rune]bool)
	for _, c := range s {
		set[c] = true
	}

	return len(set) == len(s)
}

func lookForStart(input string, length int) int {
	for i := 0; i < len(input)-1-length; i++ {
		if checkForMarker(input[i : i+length]) {
			return i + length
		}
	}

	return -1
}

func evalA(input string) int {
	return lookForStart(input, 4)
}

func evalB(input string) int {
	return lookForStart(input, 14)
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)

	resA := evalA(lines[0])
	resB := evalB(lines[0])
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 6
	//debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	//filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	//eval(filenameDebug, true)
	eval(filename, false)
}
