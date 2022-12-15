package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func calcTail(head Position, tail Position) Position {
	if util.AbsInt(head.x-tail.x) >= 2 {
		tail.x += util.SignInt(head.x - tail.x)
		tail.y += util.SignInt(head.y - tail.y)
	} else if util.AbsInt(head.y-tail.y) >= 2 {
		tail.x += util.SignInt(head.x - tail.x)
		tail.y += util.SignInt(head.y - tail.y)
	}

	return tail
}

func evalA(lines []string) int {
	visited := make(map[Position]bool)
	head := Position{x: 0, y: 0}
	tail := Position{x: 0, y: 0}

	for _, instruction := range lines {
		direction := strings.Split(instruction, " ")[0]
		times, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
		for i := 0; i < times; i++ {
			switch direction {
			case "L":
				head.x -= 1
			case "R":
				head.x += 1
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			}

			tail = calcTail(head, tail)
			visited[tail] = true
		}
	}

	return len(visited)
}

func evalB(lines []string) int {
	visited := make(map[Position]bool)
	positions := make([]Position, 10)

	for _, instruction := range lines {
		direction := strings.Split(instruction, " ")[0]
		times, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
		for i := 0; i < times; i++ {
			switch direction {
			case "L":
				positions[0].x -= 1
			case "R":
				positions[0].x += 1
			case "U":
				positions[0].y += 1
			case "D":
				positions[0].y -= 1
			}
			for i := 0; i < len(positions)-1; i++ {
				positions[i+1] = calcTail(positions[i], positions[i+1])
			}

			visited[positions[len(positions)-1]] = true
		}
	}

	return len(visited)
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
	day := 9
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
