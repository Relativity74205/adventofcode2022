package util

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

const ResourcePath = "resources"

func ReadFile(filename string) []string {
	file, err := os.Open(ResourcePath + "/" + filename)

	if err != nil {
		log.Fatalf("failed to open: %v", err)

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()

	return lines
}

func LinesToInt(lines []string) [][]int {
	var intLines [][]int
	for _, line := range lines {
		var intLine []int
		for _, c := range line {
			val, _ := strconv.Atoi(string(c))
			intLine = append(intLine, val)
		}
		intLines = append(intLines, intLine)
	}

	return intLines
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func SignInt(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	} else {
		return 0
	}
}

func PowInt(base int, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result
}

func MinInt(i1, i2 int) int {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func MaxInt(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func StringSliceToIntSlice(s []string) ([]int, error) {
	var integers []int
	for _, ele := range s {
		integer, err := strconv.Atoi(ele)
		if err != nil {
			return nil, err
		}
		integers = append(integers, integer)
	}

	return integers, nil
}

func CheckInBounds(mapArray [][]int, x, y int) bool {
	maxHeight := len(mapArray) - 1
	maxWidth := len(mapArray[0]) - 1

	if x < 0 || y < 0 || x > maxWidth || y > maxHeight {
		return false
	}

	return true
}

func GetIndexInStringSlice(slice []string, searchedValue string) (int, error) {
	for i, val := range slice {
		if val == searchedValue {
			return i, nil
		}
	}

	return -1, errors.New("searchedValue not found")
}
