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

type Space struct {
	xMin, xMax, yMin, yMax, zMin, zMax int
}

type ToVisitStructure struct {
	list          []Coords
	alreadyOnList map[Coords]bool
}

func newToVisit() ToVisitStructure {
	var toVisit ToVisitStructure
	toVisit.alreadyOnList = make(map[Coords]bool)
	return toVisit
}

func (v *ToVisitStructure) add(coordinate Coords) {
	v.list = append(v.list, coordinate)
	v.alreadyOnList[coordinate] = true
}

func (v *ToVisitStructure) pop() Coords {
	coordinate := v.list[0]
	v.list = v.list[1:]
	return coordinate
}

func (v *ToVisitStructure) finished() bool {
	return len(v.list) == 0
}

func evalA(lines []string) int {
	drops := parseCoords(lines)
	totalSurface := getTotalSurface(drops)

	return totalSurface
}

func getTotalSurface(drops map[Coords]bool) int {
	totalSurface := 6 * len(drops)
	for drop := range drops {
		for _, coords := range getNeighbors(drop) {
			if drops[coords] {
				totalSurface--
			}
		}
	}
	return totalSurface
}

func getExternalSurface(drops map[Coords]bool, outsideSpace map[Coords]bool) int {
	totalSurface := 6 * len(drops)
	for drop := range drops {
		for _, coords := range getNeighbors(drop) {
			if drops[coords] || !outsideSpace[coords] {
				totalSurface--
			}
		}
	}
	return totalSurface
}

func parseCoords(lines []string) map[Coords]bool {
	drops := make(map[Coords]bool)

	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		drops[Coords{x, y, z}] = true
	}

	return drops
}

func evalB(lines []string) int {
	drops := parseCoords(lines)
	outsideSpace := getOutsideSpace(drops)
	externalSurface := getExternalSurface(drops, outsideSpace)

	return externalSurface
}

func getNeighbors(coordinate Coords) []Coords {
	var neighbors []Coords
	neighbors = append(neighbors, Coords{coordinate.x - 1, coordinate.y, coordinate.z})
	neighbors = append(neighbors, Coords{coordinate.x + 1, coordinate.y, coordinate.z})
	neighbors = append(neighbors, Coords{coordinate.x, coordinate.y - 1, coordinate.z})
	neighbors = append(neighbors, Coords{coordinate.x, coordinate.y + 1, coordinate.z})
	neighbors = append(neighbors, Coords{coordinate.x, coordinate.y, coordinate.z - 1})
	neighbors = append(neighbors, Coords{coordinate.x, coordinate.y, coordinate.z + 1})

	return neighbors
}

func withinBounds(coordinate Coords, space *Space) bool {
	if coordinate.x < space.xMin || coordinate.x > space.xMax {
		return false
	}
	if coordinate.y < space.yMin || coordinate.y > space.yMax {
		return false
	}
	if coordinate.z < space.zMin || coordinate.z > space.zMax {
		return false
	}

	return true
}

func getOutsideSpace(drops map[Coords]bool) map[Coords]bool {
	maxSearchSpace := getSearchSpace(drops)
	outsideSpace := make(map[Coords]bool)
	toVisit := newToVisit()
	toVisit.add(Coords{maxSearchSpace.xMin, maxSearchSpace.yMin, maxSearchSpace.zMin})

	for !toVisit.finished() {
		coordinate := toVisit.pop()
		if outsideSpace[coordinate] || drops[coordinate] {
			continue
		}

		for _, neighbor := range getNeighbors(coordinate) {
			if withinBounds(neighbor, maxSearchSpace) {
				toVisit.add(neighbor)
			}
		}

		outsideSpace[coordinate] = true
	}

	return outsideSpace
}

func getSearchSpace(drops map[Coords]bool) *Space {
	xMin := 1
	yMin := 1
	zMin := 1
	var xMax, yMax, zMax int
	for coords := range drops {
		xMin = util.MinInt(xMin, coords.x)
		xMax = util.MaxInt(xMax, coords.x)
		yMin = util.MinInt(yMin, coords.y)
		yMax = util.MaxInt(yMax, coords.y)
		zMin = util.MinInt(zMin, coords.z)
		zMax = util.MaxInt(zMax, coords.z)
	}
	return &Space{xMin - 1, xMax + 1, yMin - 1, yMax + 1, zMin - 1, zMax + 1}
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
