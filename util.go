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

func LinesToInt(lines []string) []int {
	var intLines []int
	for _, v := range lines {
		vint, _ := strconv.Atoi(v)
		intLines = append(intLines, vint)
	}

	return intLines
}
