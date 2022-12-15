package main

import (
	"AoC2022"
	"fmt"
	"strconv"
	"strings"
)

type inputPair struct {
	signal1, signal2 string
}

type SignalPart interface {
	partType() string
}

type SignalValue struct {
	val int
}

func (sv SignalValue) partType() string {
	return "value"
}

type SignalList struct {
	values []SignalPart
}

func (sl SignalList) partType() string {
	return "list"
}

func getDelimiterPositions(s string) []int {
	var delimiterPositions []int
	var openBrackets int

	for i, c := range s {
		switch string(c) {
		case "[":
			openBrackets += 1
		case "]":
			openBrackets -= 1
		case ",":
			if openBrackets == 0 {
				delimiterPositions = append(delimiterPositions, i)
			}
		}
	}

	return delimiterPositions
}

func parsePart(signalPart string) SignalPart {
	var parsedList []SignalPart

	// when the string consists of one number
	if number, err := strconv.Atoi(signalPart); err == nil {
		return SignalValue{number}
	}
	if strings.HasPrefix(signalPart, "[") {
		signalPart = signalPart[1 : len(signalPart)-1]
	}

	delimiterPositions := getDelimiterPositions(signalPart)
	// when list is empty or only consists of one element
	if len(delimiterPositions) == 0 {
		if signalPart == "" {
			return SignalList{[]SignalPart{}}
		} else {
			return SignalList{[]SignalPart{parsePart(signalPart)}}
		}
	}

	var sPart string
	for i := 0; i <= len(delimiterPositions); i++ {
		if i == 0 {
			sPart = signalPart[:delimiterPositions[0]]
		} else if i == len(delimiterPositions) {
			sPart = signalPart[delimiterPositions[i-1]+1:]
		} else {
			sPart = signalPart[delimiterPositions[i-1]+1 : delimiterPositions[i]]
		}
		parsedList = append(parsedList, parsePart(sPart))
	}

	return SignalList{parsedList}
}

func compareSignalValues(s1, s2 SignalValue) int {
	if s1.val > s2.val {
		return -1
	} else if s1.val < s2.val {
		return 1
	} else {
		return 0
	}
}

func compareSignalParts(s1, s2 SignalPart) int {
	if s1.partType() == "value" && s2.partType() == "value" {
		return compareSignalValues(s1.(SignalValue), s2.(SignalValue))
	} else if s1.partType() == "value" && s2.partType() == "list" {
		return compareSignalParts(SignalList{[]SignalPart{s1}}, s2)
	} else if s1.partType() == "list" && s2.partType() == "value" {
		return compareSignalParts(s1, SignalList{[]SignalPart{s2}})
	}

	// both are lists
	s1Values := s1.(SignalList).values
	s2Values := s2.(SignalList).values
	for i := 0; i < len(s1Values); i++ {
		if len(s2Values)-1 < i {
			return -1
		}
		if result := compareSignalParts(s1Values[i], s2Values[i]); result != 0 {
			return result
		}
	}
	if len(s2Values) > len(s1Values) {
		return 1
	}

	return 0
}

func packetPairInOrder(packet1, packet2 string) bool {
	signal1 := parsePart(packet1)
	signal2 := parsePart(packet2)

	compareResult := compareSignalParts(signal1, signal2)

	return compareResult == 0 || compareResult == 1
}

func evalA(inputPairs []inputPair) int {
	var sumIndex int

	for i, inputPair := range inputPairs {
		if packetPairInOrder(inputPair.signal1, inputPair.signal2) {
			sumIndex += i + 1
		}
	}

	return sumIndex
}

func evalB(packets []string) int {
	var sortedPackets []string
	dividerPacket1 := "[[2]]"
	dividerPacket2 := "[[6]]"
	packets = append(packets, dividerPacket1, dividerPacket2)
	sortedPackets = append(sortedPackets, packets[0])
	for _, packet := range packets[1:] {
		countSortedPackets := len(sortedPackets)
		for i := 0; i < countSortedPackets; i++ {
			if packetPairInOrder(packet, sortedPackets[i]) {
				sortedPackets = append(sortedPackets[:i+1], sortedPackets[i:]...)
				sortedPackets[i] = packet
				break
			}

			// last action in inner for loop
			if i+1 == len(sortedPackets) {
				sortedPackets = append(sortedPackets, packet)
			}
		}
	}

	factor1, _ := util.GetIndexInStringSlice(sortedPackets, dividerPacket1)
	factor2, _ := util.GetIndexInStringSlice(sortedPackets, dividerPacket2)

	return (factor1 + 1) * (factor2 + 1)
}

func getInputPairs(lines []string) []inputPair {
	var pairs []inputPair
	for i := 0; i < (len(lines)+1)/3; i += 1 {
		pairs = append(pairs, inputPair{lines[i*3], lines[i*3+1]})
	}

	return pairs
}

func getPackets(lines []string) []string {
	var packets []string
	for _, line := range lines {
		if line == "" {
			continue
		}
		packets = append(packets, line)
	}

	return packets
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	inputPairs := getInputPairs(lines)
	packets := getPackets(lines)

	resA := evalA(inputPairs)
	resB := evalB(packets)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
		fmt.Printf("B (debug): %v \n", resB)
	} else {
		fmt.Printf("A: %v \n", resA)
		fmt.Printf("B: %v \n", resB)
	}

}

func main() {
	day := 13
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
