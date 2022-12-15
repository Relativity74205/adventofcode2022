package main

import (
	"AoC2022"
	"container/heap"
	"errors"
	"fmt"
)

const start = 19
const end = 5
const a = 33
const z = 58

var neighborDeltas = []Pos{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

type PriorityQueue []Path

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance() < pq[j].distance()
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Path))
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type Pos struct {
	x, y int
}

type Path struct {
	visitedPos []Pos
}

func (p Path) lastPos() Pos {
	return p.visitedPos[len(p.visitedPos)-1]
}

func (p Path) distance() int {
	return len(p.visitedPos)
}

func getPositionByRuneID(valleyMap [][]int, runeId int, height int) (Pos, error) {
	for y, row := range valleyMap {
		for x, val := range row {
			if val == runeId {
				valleyMap[y][x] = height
				return Pos{x, y}, nil
			}
		}
	}

	return Pos{}, errors.New("position not found")
}

func getStartPosition(valleyMap [][]int) (Pos, error) {
	return getPositionByRuneID(valleyMap, start, a)
}

func getEndPosition(valleyMap [][]int) (Pos, error) {
	return getPositionByRuneID(valleyMap, end, z)
}

func getNewPaths(valleyMap [][]int, path Path) []Path {
	var newPaths []Path

	for _, delta := range neighborDeltas {
		lastPos := path.lastPos()
		newPos := Pos{lastPos.x + delta.x, lastPos.y + delta.y}

		if AoC2022.CheckInBounds(valleyMap, newPos.x, newPos.y) {
			currentHeight := getHeight(valleyMap, lastPos)
			newHeight := getHeight(valleyMap, newPos)
			if currentHeight+1 >= newHeight {
				copiedPath := make([]Pos, len(path.visitedPos))
				copy(copiedPath, path.visitedPos)
				newPath := Path{
					append(copiedPath, newPos),
				}
				newPaths = append(newPaths, newPath)
			}
		}
	}

	return newPaths
}

func getHeight(valleyMap [][]int, pos Pos) int {
	return valleyMap[pos.y][pos.x]
}

func traverse(startPos Pos, endPos Pos, valleyMap [][]int) (int, error) {

	h := &PriorityQueue{}
	heap.Init(h)
	heap.Push(h, Path{[]Pos{startPos}})
	visitedPos := make(map[Pos]bool)
	toVisitPos := make(map[Pos]int)
	toVisitPos[startPos] = 0

	for h.Len() > 0 {
		path := heap.Pop(h).(Path)
		if path.lastPos() == endPos {
			return path.distance() - 1, nil
		}

		visitedPos[path.lastPos()] = true

		newPaths := getNewPaths(valleyMap, path)
		for _, newPath := range newPaths {
			_, alreadyVisited := visitedPos[newPath.lastPos()]
			oldDistance, onVisitPos := toVisitPos[newPath.lastPos()]
			if alreadyVisited {
				continue
			}
			if onVisitPos && newPath.distance() >= oldDistance {
				continue
			}
			toVisitPos[newPath.lastPos()] = newPath.distance()
			heap.Push(h, newPath)
		}
	}

	return -1, errors.New("no path found")
}

func evalA(valleyMap [][]int, startPosition Pos, endPosition Pos) int {
	minDistanceTraveled, _ := traverse(startPosition, endPosition, valleyMap)

	return minDistanceTraveled
}

func getAllPotentialStartPositions(valleyMap [][]int) []Pos {
	var potentialStartPositions []Pos
	for y, row := range valleyMap {
		for x, val := range row {
			if val == a {
				potentialStartPositions = append(potentialStartPositions, Pos{x, y})
			}
		}
	}
	return potentialStartPositions
}

func evalB(valleyMap [][]int, endPosition Pos) int {
	potentialStartPositions := getAllPotentialStartPositions(valleyMap)

	bestMinDistanceTraveled := 1000000
	for _, startPosition := range potentialStartPositions {
		minDistanceTraveled, err := traverse(startPosition, endPosition, valleyMap)
		if err == nil {
			bestMinDistanceTraveled = AoC2022.MinInt(bestMinDistanceTraveled, minDistanceTraveled)
		}
	}

	return bestMinDistanceTraveled
}

func createValleyMap(lines []string) [][]int {
	var valleyMap [][]int
	for _, line := range lines {
		var valleyRow []int
		for _, c := range line {
			valleyRow = append(valleyRow, int(c)-64)
		}
		valleyMap = append(valleyMap, valleyRow)
	}

	return valleyMap
}

func eval(filename string, debug bool) {
	lines := AoC2022.ReadFile(filename)
	valleyMap := createValleyMap(lines)
	startPosition, _ := getStartPosition(valleyMap)
	endPosition, _ := getEndPosition(valleyMap)

	resA := evalA(valleyMap, startPosition, endPosition)
	resB := evalB(valleyMap, endPosition)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 12
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
