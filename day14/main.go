package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

const caveWidth = 1000
const caveHeight = 200

type coordinates struct {
	x, y int
}

type Cave struct {
	caveMap  [][]uint
	maxDepth int
}

func printCaveMap(cave Cave) {
	for y, line := range cave.caveMap {
		if y > cave.maxDepth+2 {
			continue
		}

		for x, val := range line {
			if x >= 480 && x <= 520 {
				fmt.Print(val)
			}
		}
		fmt.Println("")
	}
}

func makeCaveMap(lines []string) Cave {
	var totalYMax int
	caveMap := make([][]uint, caveHeight)
	for i := range caveMap {
		caveMap[i] = make([]uint, caveWidth)
	}

	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		for i := 0; i <= len(coords)-2; i++ {
			coordsStart := getCoordinates(coords[i])
			coordsEnd := getCoordinates(coords[i+1])
			if coordsStart.x == coordsEnd.x {
				yMin := util.MinInt(coordsStart.y, coordsEnd.y)
				yMax := util.MaxInt(coordsStart.y, coordsEnd.y)
				totalYMax = util.MaxInt(totalYMax, yMax)
				for y := yMin; y <= yMax; y++ {
					caveMap[y][coordsStart.x] = 1
				}
			} else {
				xMin := util.MinInt(coordsStart.x, coordsEnd.x)
				xMax := util.MaxInt(coordsStart.x, coordsEnd.x)
				for x := xMin; x <= xMax; x++ {
					caveMap[coordsStart.y][x] = 1
				}
			}
		}
	}

	for x := 0; x < caveWidth; x++ {
		caveMap[totalYMax+2][x] = 1
	}

	return Cave{caveMap, totalYMax}
}

func getCoordinates(coordsString string) coordinates {
	coordsStringParts := strings.Split(coordsString, ",")
	x, _ := strconv.Atoi(coordsStringParts[0])
	y, _ := strconv.Atoi(coordsStringParts[1])
	return coordinates{x, y}
}

func runSandUnitA(cave Cave) bool {
	x := 500
	y := 0
	for true {
		if y > cave.maxDepth {
			return true
		} else if cave.caveMap[y+1][x] == 0 {
			y += 1
		} else if cave.caveMap[y+1][x] != 0 && cave.caveMap[y+1][x-1] == 0 {
			x -= 1
			y += 1
		} else if cave.caveMap[y+1][x] != 0 && cave.caveMap[y+1][x+1] == 0 {
			x += 1
			y += 1
		} else {
			cave.caveMap[y][x] = 2
			break
		}
	}

	return false
}

func runSandUnitB(cave Cave) bool {
	x := 500
	y := 0
	for true {
		if cave.caveMap[y+1][x] == 0 {
			y += 1
		} else if cave.caveMap[y+1][x] != 0 && cave.caveMap[y+1][x-1] == 0 {
			x -= 1
			y += 1
		} else if cave.caveMap[y+1][x] != 0 && cave.caveMap[y+1][x+1] == 0 {
			x += 1
			y += 1
		} else {
			cave.caveMap[y][x] = 2
			if x == 500 && y == 0 {
				return true
			} else {
				break
			}
		}
	}

	return false
}

func evalA(lines []string) int {
	cave := makeCaveMap(lines)
	var countSandUnits int
	for true {
		if runSandUnitA(cave) {
			break
		}
		countSandUnits += 1
	}

	return countSandUnits
}

func evalB(lines []string) int {
	cave := makeCaveMap(lines)
	var countSandUnits int
	for true {
		countSandUnits += 1
		if runSandUnitB(cave) {
			break
		}
	}

	return countSandUnits
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
	day := 14
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
