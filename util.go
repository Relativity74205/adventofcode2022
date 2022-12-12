package AoC2022

import (
	"bufio"
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
