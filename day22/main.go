package main

import (
	"AoC2022"
	"errors"
	"fmt"
	"strconv"
)

type Instruction struct {
	amountsMoves int
	rotate       string
}

type Board struct {
	boardMap    [][]int
	boardWidth  int
	boardHeight int
	x           int
	y           int
	orientation int
}

func (b *Board) checkY(newY int, startY int) error {

	if b.checkForEmpty(b.x, newY) {
		reappearPointY := b.getReappearPointY(startY, b.x)
		if b.checkForWall(b.x, reappearPointY) {
			return errors.New("hit wall")
		} else {
			b.y = reappearPointY
			return nil
		}
	}
	b.y = newY
	return nil
}

func (b *Board) move() error {
	switch b.orientation {
	case 0:
		newY := b.y - 1
		if b.checkForEmpty(b.x, newY) {
			newY = b.getReappearPointY(b.boardHeight, b.x)
		}
		if b.checkForWall(b.x, newY) {
			return errors.New("hit wall")
		}
		b.y = newY
	case 90:
		newX := b.x + 1
		if b.checkForEmpty(newX, b.y) {
			newX = b.getReappearPointX(0, b.y)
		}
		if b.checkForWall(newX, b.y) {
			return errors.New("hit wall")
		}
		b.x = newX
	case 180:
		newY := b.y + 1
		if b.checkForEmpty(b.x, newY) {
			newY = b.getReappearPointY(0, b.x)
		}
		if b.checkForWall(b.x, newY) {
			return errors.New("hit wall")
		}
		b.y = newY
	case 270:
		newX := b.x - 1
		if b.checkForEmpty(newX, b.y) {
			newX = b.getReappearPointX(b.boardWidth, b.y)
		}
		if b.checkForWall(newX, b.y) {
			return errors.New("hit wall")
		}
		b.x = newX
	}

	return nil
}

func (b *Board) checkOutsideOfBoard(x, y int) bool {
	if x < 0 || x >= b.boardWidth || y < 0 || y >= b.boardHeight {
		return true
	}
	return false
}

func (b *Board) checkForWall(x, y int) bool {
	if b.checkOutsideOfBoard(x, y) {
		return false
	}
	return b.boardMap[y][x] == 2
}

func (b *Board) checkForEmpty(x, y int) bool {
	if b.checkOutsideOfBoard(x, y) {
		return true
	}
	return b.boardMap[y][x] == 0
}

func (b *Board) getReappearPointX(startSearchX int, y int) int {
	if startSearchX == 0 {
		for x := 0; x < b.boardWidth; x++ {
			if b.boardMap[y][x] != 0 {
				return x
			}
		}
	} else {
		for x := b.boardWidth - 1; x >= 0; x-- {
			if b.boardMap[y][x] != 0 {
				return x
			}
		}
	}

	return -1
}

func (b *Board) getReappearPointY(startSearchY int, x int) int {
	if startSearchY == 0 {
		for y := 0; y < b.boardHeight; y++ {
			if b.boardMap[y][x] != 0 {
				return y
			}
		}
	} else {
		for y := b.boardHeight - 1; y >= 0; y-- {
			if b.boardMap[y][x] != 0 {
				return y
			}
		}
	}

	return -1
}

func (b *Board) rotate(rotateCommand string) {
	if rotateCommand == "R" {
		b.orientation += 90
	} else {
		b.orientation -= 90
	}

	b.orientation = (b.orientation + 360) % 360
}

func evalA(lines []string) int {
	board, instructions := parseInput(lines)
	for _, instruction := range instructions {
		if instruction.rotate != "" {
			board.rotate(instruction.rotate)
		} else {
			for i := 0; i < instruction.amountsMoves; i++ {
				err := board.move()
				if err != nil {
					break
				}
			}
		}
	}
	orientationScore := getOrientationScore(board.orientation)

	return 1000*(board.y+1) + 4*(board.x+1) + orientationScore
}

func getOrientationScore(orientation int) int {
	switch orientation {
	case 0:
		return 3
	case 90:
		return 0
	case 180:
		return 1
	case 270:
		return 2
	default:
		return -1
	}
}

func evalB(lines []string) int {

	return 0
}

func getStartPosX(firstRow []int) int {
	for x, tile := range firstRow {
		if tile == 1 {
			return x
		}
	}
	return -1
}

func parseInput(lines []string) (Board, []Instruction) {
	instructionsString := lines[len(lines)-1]
	var instructions []Instruction
	var newMoveInstruction string
	for _, char := range instructionsString {
		if char == 'R' || char == 'L' {
			amountMoves, _ := strconv.Atoi(newMoveInstruction)
			instructions = append(instructions, Instruction{amountsMoves: amountMoves})
			newMoveInstruction = ""
			instructions = append(instructions, Instruction{rotate: string(char)})
		} else {
			newMoveInstruction += string(char)
		}
	}
	if newMoveInstruction != "" {
		amountMoves, _ := strconv.Atoi(newMoveInstruction)
		instructions = append(instructions, Instruction{amountsMoves: amountMoves})
	}

	boardHeight := len(lines) - 2
	boardWidth := 0
	for _, line := range lines[:len(lines)-2] {
		boardWidth = util.MaxInt(boardWidth, len(line))
	}

	boardMap := make([][]int, boardHeight)
	for y, line := range lines[:len(lines)-2] {
		boardMap[y] = make([]int, boardWidth)
		for x, char := range line {
			tileKind := 0
			switch char {
			case '.':
				tileKind = 1
			case '#':
				tileKind = 2
			}
			boardMap[y][x] = tileKind
		}
	}

	startPosX := getStartPosX(boardMap[0])

	direction := 1
	for _, instruction := range instructions {
		switch instruction.rotate {
		case "R":
			direction += 1
		case "L":
			direction -= 1
		default:
		}
	}

	return Board{
		boardMap:    boardMap,
		boardWidth:  boardWidth,
		boardHeight: boardHeight,
		x:           startPosX,
		y:           0,
		orientation: 90,
	}, instructions
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
	day := 22
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
