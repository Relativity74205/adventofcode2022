package main

import (
	"AoC2022"
	"fmt"
	"strconv"
)

type number struct {
	value         int
	position      int
	amountNumbers int
	left          *number
	right         *number
}

func (n *number) move() {
	valueRemainder := n.value % (n.amountNumbers - 1)

	if valueRemainder < 0 {
		for i := 0; i > valueRemainder; i-- {
			n.moveLeft()
		}
	} else {
		for i := 0; i < valueRemainder; i++ {
			n.moveRight()
		}
	}
}

func (n *number) moveRight() {
	currentLeft := n.left
	currentRight := n.right
	currentRightRight := currentRight.right
	currentPosition := n.position

	n.position = currentRight.position
	currentRight.position = currentPosition

	currentLeft.right = currentRight
	currentRight.left = currentLeft
	currentRight.right = n
	n.right = currentRightRight
	currentRightRight.left = n
	n.left = currentRight
}

func (n *number) moveLeft() {
	currentLeft := n.left
	currentLeftLeft := currentLeft.left
	currentRight := n.right
	currentPosition := n.position

	n.position = currentLeft.position
	currentLeft.position = currentPosition

	currentLeftLeft.right = n
	n.left = currentLeftLeft
	n.right = currentLeft
	currentLeft.left = n
	currentLeft.right = currentRight
	currentRight.left = currentLeft
}

func parseInput(lines []string) ([]*number, int) {
	var zeroPosition int
	var numbers []*number
	for i, line := range lines {
		value, _ := strconv.Atoi(line)
		if value == 0 {
			zeroPosition = i
		}
		number := &number{value, i, len(lines), nil, nil}

		numbers = append(numbers, number)
	}

	for i := 0; i < len(numbers); i++ {
		var leftIndex, rightIndex int
		if i == 0 {
			leftIndex = len(numbers) - 1
		} else {
			leftIndex = i - 1
		}
		if i == len(numbers)-1 {
			rightIndex = 0
		} else {
			rightIndex = i + 1
		}
		numbers[i].left = numbers[leftIndex]
		numbers[i].right = numbers[rightIndex]
	}

	return numbers, zeroPosition
}

func evalA(lines []string) int {
	numbers, zeroPosition := parseInput(lines)
	for _, n := range numbers {
		n.move()
	}

	return getGroveCoordinates(numbers, zeroPosition)
}

func getGroveCoordinates(numbers []*number, zeroPosition int) int {
	var sum int
	currentNumber := numbers[zeroPosition]
	for i := 1; i <= 3000; i++ {
		currentNumber = currentNumber.right
		if i == 1000 || i == 2000 || i == 3000 {
			sum += currentNumber.value
		}
	}
	return sum
}

func evalB(lines []string) int {
	numbers, zeroPosition := parseInput(lines)
	for _, n := range numbers {
		n.value *= 811589153
	}
	for i := 0; i < 10; i++ {
		for _, n := range numbers {
			n.move()
		}
	}

	return getGroveCoordinates(numbers, zeroPosition)
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
	day := 20
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
