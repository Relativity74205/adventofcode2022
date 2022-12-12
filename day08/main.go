package main

import (
	"AoC2022"
	"fmt"
	"math"
)

func checkYTop(x, y int, lines [][]int) bool {
	for i := x - 1; i >= 0; i-- {
		if lines[i][y] >= lines[x][y] {
			return true
		}
	}

	return false
}

func checkYBottom(x, y int, lines [][]int) bool {
	for i := x + 1; i < len(lines); i++ {
		if lines[i][y] >= lines[x][y] {
			return true
		}
	}

	return false
}

func checkXLeft(x, y int, lines [][]int) bool {
	for i := y - 1; i >= 0; i-- {
		if lines[x][i] >= lines[x][y] {
			return true
		}
	}

	return false

}

func checkXRight(x, y int, lines [][]int) bool {
	for i := y + 1; i < len(lines); i++ {
		if lines[x][i] >= lines[x][y] {
			return true
		}
	}

	return false

}

func checkYTopDist(x, y int, lines [][]int) int {
	for i := x - 1; i >= 0; i-- {
		if lines[i][y] >= lines[x][y] {
			return int(math.Abs(float64(x - i)))
		}
	}

	return x
}

func checkYBottomDist(x, y int, lines [][]int) int {
	for i := x + 1; i < len(lines); i++ {
		if lines[i][y] >= lines[x][y] {
			return int(math.Abs(float64(x - i)))
		}
	}

	return len(lines) - x - 1
}

func checkXLeftDist(x, y int, lines [][]int) int {
	for i := y - 1; i >= 0; i-- {
		if lines[x][i] >= lines[x][y] {
			return int(math.Abs(float64(y - i)))
		}
	}

	return y

}

func checkXRightDist(x, y int, lines [][]int) int {
	for i := y + 1; i < len(lines); i++ {
		if lines[x][i] >= lines[x][y] {
			return int(math.Abs(float64(y - i)))
		}
	}

	return len(lines) - y - 1

}

func evalA(lines [][]int) int {
	var visible int
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if checkXLeft(i, j, lines) && checkXRight(i, j, lines) && checkYTop(i, j, lines) && checkYBottom(i, j, lines) {
				visible += 1
			}
		}

	}
	return int(math.Pow(float64(len(lines)), 2)) - visible
}

func evalB(lines [][]int) int {
	var maxDistance float64
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			distance := checkXLeftDist(i, j, lines) * checkXRightDist(i, j, lines) * checkYTopDist(i, j, lines) * checkYBottomDist(i, j, lines)
			maxDistance = math.Max(maxDistance, float64(distance))
		}
	}
	return int(maxDistance)
}

func eval(filename string, debug bool) {
	lines := AoC2022.ReadFile(filename)
	intLines := AoC2022.LinesToInt(lines)

	resA := evalA(intLines)
	resB := evalB(intLines)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 8
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
