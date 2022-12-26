package main

import (
	"AoC2022"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type sensorArea struct {
	x, y, dist int
}

type refLineCoverage struct {
	x1, x2 int
}

func getRefLineCoverages(sensorAreas []sensorArea, yRef int) []refLineCoverage {
	var refLineCoverages []refLineCoverage
	for _, sensorArea := range sensorAreas {
		deltaX := sensorArea.dist - util.AbsInt(sensorArea.y-yRef)
		if deltaX > 0 {
			refLineCoverages = append(refLineCoverages, refLineCoverage{sensorArea.x - deltaX, sensorArea.x + deltaX})
		}
	}

	sort.Slice(refLineCoverages, func(i, j int) bool {
		return refLineCoverages[i].x1 < refLineCoverages[j].x1
	})

	for i := 0; i < len(refLineCoverages)-1; i++ {
		if refLineCoverages[i].x2 < refLineCoverages[i+1].x1 {
			continue
		}
		refLineCoverages[i].x2 = util.MaxInt(refLineCoverages[i].x2, refLineCoverages[i+1].x2)
		refLineCoverages = append(refLineCoverages[:i+1], refLineCoverages[i+2:]...)
		i -= 1
	}

	return refLineCoverages
}

func calcCoverage(refLineCoverages []refLineCoverage) int {
	sum := 0
	for _, refLineCoverage := range refLineCoverages {
		sum += refLineCoverage.x2 - refLineCoverage.x1
	}

	return sum
}

func evalA(sensorAreas []sensorArea, yRef int) int {
	refLineCoverages := getRefLineCoverages(sensorAreas, yRef)

	return calcCoverage(refLineCoverages)
}

func evalB(sensorAreas []sensorArea, yRefMax int) int {
	for yRef := 0; yRef <= yRefMax; yRef++ {
		refLineCoverages := getRefLineCoverages(sensorAreas, yRef)
		for i := range refLineCoverages {
			refLineCoverages[i].x1 = util.MinInt(util.MaxInt(refLineCoverages[i].x1, 0), yRefMax)
			refLineCoverages[i].x2 = util.MinInt(util.MaxInt(refLineCoverages[i].x2, 0), yRefMax)
		}
		if calcCoverage(refLineCoverages) != yRefMax {
			return yRef + (refLineCoverages[0].x2+1)*4000000 // assumes that refLineCoverages will most probably only consist of two elements
		}
	}

	return 0
}

func readReadings(lines []string) []sensorArea {
	var sensorAreas []sensorArea
	for _, line := range lines {
		sensorData := strings.Split(line, ": ")[0][10:]
		sensorX, _ := strconv.Atoi(strings.Split(sensorData, ", ")[0][2:])
		sensorY, _ := strconv.Atoi(strings.Split(sensorData, ", ")[1][2:])
		beaconData := strings.Split(line, ": ")[1][21:]
		beaconX, _ := strconv.Atoi(strings.Split(beaconData, ", ")[0][2:])
		beaconY, _ := strconv.Atoi(strings.Split(beaconData, ", ")[1][2:])
		dist := util.AbsInt(sensorX-beaconX) + util.AbsInt(sensorY-beaconY)
		sensorAreas = append(sensorAreas, sensorArea{sensorX, sensorY, dist})
	}

	return sensorAreas
}

func eval(filename string, debug bool, yRef int) {
	lines := util.ReadFile(filename)
	readings := readReadings(lines)

	resA := evalA(readings, yRef)
	resB := evalB(readings, yRef*2)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 15
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true, 10)
	eval(filename, false, 2000000)
}
