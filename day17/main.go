package main

import (
	"AoC2022"
	"fmt"
)

type Coords struct {
	x int
	y int
}

type Rock struct {
	parts     []Coords
	refCoords Coords
}

func (r *Rock) Width() int {
	width := 0
	for _, part := range r.parts {
		width = util.MaxInt(width, part.x+1)
	}

	return width
}

type Cave struct {
	Map         [][]bool
	highestRock int
}

var rockSequence = map[int][]Coords{
	1: {{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	2: {{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
	3: {{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
	4: {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	0: {{0, 0}, {1, 0}, {0, 1}, {1, 1}},
}

const maxRock = 2022

func checkFallingPossible(cave *Cave, rock *Rock) bool {
	for _, part := range rock.parts {
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
	for _, part := range rock.parts {
		x := rock.refCoords.x + part.x + delta
		y := rock.refCoords.y + part.y
		if x >= 0 && x <= 6 && cave.Map[y][x] {
			return false
		}
	}

	return true
}

func solidifyRock(cave *Cave, rock *Rock) {
	for _, part := range rock.parts {
		x := rock.refCoords.x + part.x
		y := rock.refCoords.y + part.y
		cave.Map[y][x] = true
		cave.highestRock = util.MaxInt(cave.highestRock, y)
	}
}

func evalA(winds string) int {
	cave := &Cave{[][]bool{}, 0}
	cave.Map = append(cave.Map, []bool{true, true, true, true, true, true, true})

	round := 0
	for rockNumber := 1; rockNumber <= maxRock; rockNumber++ {
		startX := 2
		startY := cave.highestRock + 4
		rock := &Rock{rockSequence[rockNumber%5], Coords{startX, startY}}
		for i := cave.highestRock + 1; i <= startY; i++ {
			cave.Map = append(cave.Map, make([]bool, 7))
		}

		for true {
			switch string(winds[round%len(winds)]) {
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
	}

	return cave.highestRock
}

func evalB(winds string) int {

	return 0
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
