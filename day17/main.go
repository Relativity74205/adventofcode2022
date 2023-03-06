package main

import (
	"AoC2022"
	"fmt"
)

type RockParts []Coords

type Coords struct {
	x int
	y int
}

type Rock struct {
	rockParts RockParts
	refCoords Coords
}

func (r *Rock) Width() int {
	width := 0
	for _, part := range r.rockParts {
		width = util.MaxInt(width, part.x+1)
	}

	return width
}

type Cave struct {
	Map           [][]bool
	highestRock   int
	skippedHeight int
}

var rockVersions = map[int]RockParts{
	1: {{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	2: {{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
	3: {{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
	4: {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	0: {{0, 0}, {1, 0}, {0, 1}, {1, 1}},
}

func checkFallingPossible(cave *Cave, rock *Rock) bool {
	for _, part := range rock.rockParts {
		x := rock.refCoords.x + part.x
		y := rock.refCoords.y + part.y - 1
		if cave.Map[y][x] {
			return false
		}
	}

	return true
}

func checkMoveSide(cave *Cave, rock *Rock, moveRight bool) bool {
	var delta int
	if moveRight {
		delta = 1
	} else {
		delta = -1
	}
	for _, part := range rock.rockParts {
		x := rock.refCoords.x + part.x + delta
		y := rock.refCoords.y + part.y
		if x >= 0 && x <= 6 && cave.Map[y][x] {
			return false
		}
	}

	return true
}

func solidifyRock(cave *Cave, rock *Rock) {
	for _, part := range rock.rockParts {
		x := rock.refCoords.x + part.x
		y := rock.refCoords.y + part.y
		cave.Map[y][x] = true
		cave.highestRock = util.MaxInt(cave.highestRock, y)
	}
}

func evalA(winds string) int {
	cave := &Cave{[][]bool{}, 0, 0}
	cave.Map = append(cave.Map, []bool{true, true, true, true, true, true, true})

	play(winds, cave, 2022)

	return cave.highestRock + cave.skippedHeight
}

type IndexCombination struct {
	windsIndex int
	rockIndex  int
}

type Observation struct {
	rockNumber  int
	highestRock int
}

func play(winds string, cave *Cave, maxRocks int) {
	round := 0
	seen := make(map[IndexCombination][]Observation)
	for rockNumber := 1; rockNumber <= maxRocks; rockNumber++ {
		rockIndex := rockNumber % 5
		rockParts := rockVersions[rockIndex]
		rock := &Rock{rockParts, Coords{2, cave.highestRock + 4}}
		for i := cave.highestRock + 1; i <= cave.highestRock+4; i++ {
			cave.Map = append(cave.Map, make([]bool, 7))
		}

		for true {
			windsIndex := round % len(winds)
			switch string(winds[windsIndex]) {
			case ">":
				if checkMoveSide(cave, rock, true) {
					rock.refCoords.x = util.MinInt(rock.refCoords.x+1, 6-rock.Width()+1)
				}
			case "<":
				if checkMoveSide(cave, rock, false) {
					rock.refCoords.x = util.MaxInt(rock.refCoords.x-1, 0)
				}
			}
			round++

			if checkFallingPossible(cave, rock) {
				rock.refCoords.y--
			} else {
				solidifyRock(cave, rock)
				break
			}
		}

		// save seen rock and windIndex combination
		indexCombination := IndexCombination{round % len(winds), rockIndex}
		seen[indexCombination] = append(seen[indexCombination], Observation{rockNumber, cave.highestRock})

		// in case a rock/windIndex combination have been seen the third time, THE cycle has been detected,
		// skipping all possible cycles.
		if len(seen[indexCombination]) == 3 {
			deltaRocks := seen[indexCombination][2].rockNumber - seen[indexCombination][1].rockNumber
			deltaHeight := seen[indexCombination][2].highestRock - seen[indexCombination][1].highestRock
			skippedCycles := (maxRocks - rockNumber) / deltaRocks
			maxRocks -= deltaRocks * skippedCycles
			cave.skippedHeight += deltaHeight * skippedCycles
		}
	}
}

func evalB(winds string) int {
	cave := &Cave{[][]bool{}, 0, 0}
	cave.Map = append(cave.Map, []bool{true, true, true, true, true, true, true})

	play(winds, cave, 1_000_000_000_000)

	return cave.highestRock + cave.skippedHeight
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
	day := 17
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
