package main

import (
	"AoC2022"
	"fmt"
	"unicode"
)

func getCharsFromString(s string) map[rune]bool {
	chars := make(map[rune]bool)
	for _, c := range s {
		chars[c] = true
	}

	return chars
}

func calcPriority(c rune) int {
	if unicode.IsUpper(c) {
		return int(c) - 64 + 26
	} else {
		return int(c) - 96
	}
}

func getCommonRunes(chars1, chars2 map[rune]bool) map[rune]bool {
	commonChars := make(map[rune]bool)
	for c := range chars1 {
		if _, ok := chars2[c]; ok {
			commonChars[c] = true
		}
	}

	return commonChars
}

func evalA(lines []string) int {
	totalPriority := 0
	for _, line := range lines {
		backpack1 := line[:len(line)/2]
		backpack2 := line[len(line)/2:]

		backpack1Chars := getCharsFromString(backpack1)
		backpack2Chars := getCharsFromString(backpack2)

		commonChar := getCommonRunes(backpack1Chars, backpack2Chars)
		for c, _ := range commonChar {
			totalPriority += calcPriority(c)
		}
	}

	return totalPriority
}

func evalB(lines []string) int {
	totalPriority := 0
	for i := 0; i < len(lines); i += 3 {
		elf1Chars := getCharsFromString(lines[i])
		elf2Chars := getCharsFromString(lines[i+1])
		elf3Chars := getCharsFromString(lines[i+2])

		commonChars := getCommonRunes(elf1Chars, elf2Chars)
		commonChar := getCommonRunes(commonChars, elf3Chars)
		for c, _ := range commonChar {
			totalPriority += calcPriority(c)
		}
	}

	return totalPriority
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
	day := 3
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
