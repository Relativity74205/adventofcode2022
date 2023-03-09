package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

type Coords struct {
	x, y, z int
}

func evalA(lines []string) int {
	drops := make(map[Coords]bool)

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		drops[Coords{x, y, z}] = true
	}

	totalSurface := 0
	for drop, _ := range drops {
		surface := 6
		if drops[Coords{drop.x + 1, drop.y, drop.z}] {
			surface--
		}
		if drops[Coords{drop.x - 1, drop.y, drop.z}] {
			surface--
		}
		if drops[Coords{drop.x, drop.y + 1, drop.z}] {
			surface--
		}
		if drops[Coords{drop.x, drop.y - 1, drop.z}] {
			surface--
		}
		if drops[Coords{drop.x, drop.y, drop.z + 1}] {
			surface--
		}
		if drops[Coords{drop.x, drop.y, drop.z - 1}] {
			surface--
		}
		totalSurface += surface
	}

	return totalSurface
}

func evalB(lines []string) int {

	return 0
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
	day := 18
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
