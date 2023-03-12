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

func evalA(lines []string) int {
	drops := getDropCoords(lines)
	totalSurface := getTotalSurface(drops)

	return totalSurface
}

func getTotalSurface(drops map[Coords]bool) int {
	totalSurface := 0
	for drop := range drops {
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

func getExternalSurface(drops map[Coords]bool, outsideSpace map[Coords]bool) int {
	totalSurface := 0
	for drop := range drops {
		surface := 6
		if drops[Coords{drop.x + 1, drop.y, drop.z}] || !outsideSpace[Coords{drop.x + 1, drop.y, drop.z}] {
			surface--
		}
		if drops[Coords{drop.x - 1, drop.y, drop.z}] || !outsideSpace[Coords{drop.x - 1, drop.y, drop.z}] {
			surface--
		}
		if drops[Coords{drop.x, drop.y + 1, drop.z}] || !outsideSpace[Coords{drop.x, drop.y + 1, drop.z}] {
			surface--
		}
		if drops[Coords{drop.x, drop.y - 1, drop.z}] || !outsideSpace[Coords{drop.x, drop.y - 1, drop.z}] {
			surface--
		}
		if drops[Coords{drop.x, drop.y, drop.z + 1}] || !outsideSpace[Coords{drop.x, drop.y, drop.z + 1}] {
			surface--
		}
		if drops[Coords{drop.x, drop.y, drop.z - 1}] || !outsideSpace[Coords{drop.x, drop.y, drop.z - 1}] {
			surface--
		}
		totalSurface += surface
	}
	return totalSurface
}

func getDropCoords(lines []string) map[Coords]bool {
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
	drops := getDropCoords(lines)
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
	if coordinate.x < space.xMin || coordinate.x > space.xMax || coordinate.y < space.yMin || coordinate.y > space.yMax || coordinate.z < space.zMin || coordinate.z > space.zMax {
		return false
	}

	return true
}

func getOutsideSpace(drops map[Coords]bool) map[Coords]bool {
	maxSearchSpace := getSearchSpace(drops)
	outsideSpace := make(map[Coords]bool)
	var coordinatesToVisit []Coords
	coordinatesToVisitAlreadyOnList := make(map[Coords]bool)
	coordinatesToVisit = append(coordinatesToVisit, Coords{maxSearchSpace.xMin, maxSearchSpace.yMin, maxSearchSpace.zMin})
	coordinatesToVisitAlreadyOnList[Coords{maxSearchSpace.xMin, maxSearchSpace.yMin, maxSearchSpace.zMin}] = true

	for len(coordinatesToVisit) > 0 {
		coordinate := coordinatesToVisit[0]
		coordinatesToVisit = coordinatesToVisit[1:]
		if outsideSpace[coordinate] || drops[coordinate] {
			continue
		}

		neighbors := getNeighbors(coordinate)
		for _, neighbor := range neighbors {
			if withinBounds(neighbor, maxSearchSpace) {
				coordinatesToVisit = append(coordinatesToVisit, neighbor)
			}
		}

		outsideSpace[coordinate] = true
	}

	return outsideSpace
}

//func getInternalSurface(drops map[Coords]bool) int {
//	internalSpace := make(map[Coords]bool)
//	emptySpace := make(map[Coords]bool)
//	maxSearchSpace := getSearchSpace(drops)
//
//	for x := maxSearchSpace.xMin; x <= maxSearchSpace.xMax; x++ {
//		for y := maxSearchSpace.yMin; y <= maxSearchSpace.yMax; y++ {
//			for z := maxSearchSpace.zMin; z <= maxSearchSpace.zMax; z++ {
//				coordinate := Coords{x, y, z}
//				if drops[coordinate] || internalSpace[coordinate] || emptySpace[coordinate] {
//					continue
//				}
//				searchedSpace, isInternalSpace := checkIfInternalSpace(coordinate, drops, emptySpace, &maxSearchSpace)
//				if isInternalSpace {
//					for _, coords := range searchedSpace {
//						internalSpace[coords] = true
//					}
//				} else {
//					for _, coords := range searchedSpace {
//						emptySpace[coords] = true
//					}
//				}
//			}
//		}
//	}
//
//	return getTotalSurface(internalSpace)
//}
//
//func checkIfAppendToCoordsTOInvestigate(coordinate Coords, coordsToInvestigate []Coords, coordsToInvestigateAlreadyInSlice map[Coords]bool, drops map[Coords]bool, emptySpace map[Coords]bool) []Coords {
//	if !drops[coordinate] && !emptySpace[coordinate] && !coordsToInvestigateAlreadyInSlice[coordinate] {
//		coordsToInvestigate = append(coordsToInvestigate, coordinate)
//		coordsToInvestigateAlreadyInSlice[coordinate] = true
//	}
//
//	return coordsToInvestigate
//}
//
//func checkIfInternalSpace(startCoordinate Coords, drops map[Coords]bool, emptySpace map[Coords]bool, maxSearchSpace *Space) ([]Coords, bool) {
//	var coordsToInvestigate, searchedSpace []Coords
//	coordsToInvestigateAlreadyInSlice := make(map[Coords]bool)
//	coordsToInvestigate = append(coordsToInvestigate, startCoordinate)
//
//	for len(coordsToInvestigate) > 0 {
//		coordinate := coordsToInvestigate[0]
//		coordsToInvestigate = coordsToInvestigate[1:]
//		searchedSpace = append(searchedSpace, coordinate)
//		if drops[coordinate] || emptySpace[coordinate] {
//			continue
//		}
//
//		if coordinate.x == maxSearchSpace.xMin {
//			return searchedSpace, false
//		} else {
//			coordsToInvestigate = checkIfAppendToCoordsTOInvestigate(Coords{coordinate.x - 1, coordinate.y, coordinate.z}, coordsToInvestigate, coordsToInvestigateAlreadyInSlice, drops, emptySpace)
//		}
//		if coordinate.x == maxSearchSpace.xMax {
//			return searchedSpace, false
//		} else {
//			coordsToInvestigate = checkIfAppendToCoordsTOInvestigate(Coords{coordinate.x + 1, coordinate.y, coordinate.z}, coordsToInvestigate, coordsToInvestigateAlreadyInSlice, drops, emptySpace)
//		}
//		if coordinate.y == maxSearchSpace.yMin {
//			return searchedSpace, false
//		} else {
//			coordsToInvestigate = checkIfAppendToCoordsTOInvestigate(Coords{coordinate.x, coordinate.y - 1, coordinate.z}, coordsToInvestigate, coordsToInvestigateAlreadyInSlice, drops, emptySpace)
//		}
//		if coordinate.y == maxSearchSpace.yMax {
//			return searchedSpace, false
//		} else {
//			coordsToInvestigate = checkIfAppendToCoordsTOInvestigate(Coords{coordinate.x, coordinate.y + 1, coordinate.z}, coordsToInvestigate, coordsToInvestigateAlreadyInSlice, drops, emptySpace)
//		}
//		if coordinate.z == maxSearchSpace.zMin {
//			return searchedSpace, false
//		} else {
//			coordsToInvestigate = checkIfAppendToCoordsTOInvestigate(Coords{coordinate.x, coordinate.y, coordinate.z - 1}, coordsToInvestigate, coordsToInvestigateAlreadyInSlice, drops, emptySpace)
//		}
//		if coordinate.z == maxSearchSpace.zMax {
//			return searchedSpace, false
//		} else {
//			coordsToInvestigate = checkIfAppendToCoordsTOInvestigate(Coords{coordinate.x, coordinate.y, coordinate.z + 1}, coordsToInvestigate, coordsToInvestigateAlreadyInSlice, drops, emptySpace)
//		}
//	}
//
//	return searchedSpace, true
//}

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
