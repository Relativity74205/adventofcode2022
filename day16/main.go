package main

import (
	"AoC2022"
	"container/heap"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const availableTime = 30

type Strategy struct {
	visitedValves    []string
	currentFlow      int
	totalFlow        int
	time             int
	currentValveName string
	openedValves     []string
}

func (s *Strategy) currentValveOpened() bool {
	for _, openedValve := range s.openedValves {
		if openedValve == s.currentValveName {
			return true
		}
	}

	return false
}

func (s *Strategy) countValveVisits() int {
	visits := 0
	for _, visitedValve := range s.visitedValves {
		if visitedValve == s.currentValveName {
			visits++
		}
	}

	return visits
}

type Valve struct {
	name     string
	flowRate int
	tunnels  []string
}

type ValveVisit struct {
	valveName string
	times     int
}

type PriorityQueue []Strategy

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	//if pq[i].time == pq[j].time {
	//	if pq[i].totalFlow == pq[j].totalFlow {
	//		return pq[i].currentFlow > pq[j].currentFlow
	//	} else {
	//		return pq[i].totalFlow > pq[j].totalFlow
	//	}
	//} else {
	//	return pq[i].time < pq[j].time
	//}
	return pq[i].time < pq[j].time
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Strategy))
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func searchStrategy(valves map[string]Valve) (int, error) {
	visitedValves := make(map[ValveVisit]bool)
	//toVisitValves := make(map[ValveVisit]int)
	h := &PriorityQueue{}
	heap.Init(h)

	startValve := valves["AA"]
	for _, valveName := range startValve.tunnels {
		heap.Push(h, Strategy{visitedValves: []string{valveName}, time: 1, currentFlow: 0, totalFlow: 0, currentValveName: valveName})
		//toVisitValves[ValveVisit{valveName, 1}] = 0
	}

	for h.Len() > 0 {
		strategy := heap.Pop(h).(Strategy)
		if strategy.time > availableTime {
			return strategy.totalFlow, nil
		}
		visitedValves[ValveVisit{strategy.currentValveName, strategy.countValveVisits()}] = true

		currentValve := valves[strategy.currentValveName]
		if currentValve.flowRate > 0 && !strategy.currentValveOpened() {
			openedValvesCopy := make([]string, len(strategy.openedValves))
			copy(openedValvesCopy, strategy.openedValves)
			newStrategy := Strategy{
				visitedValves:    strategy.visitedValves,
				currentFlow:      strategy.currentFlow + currentValve.flowRate,
				totalFlow:        strategy.totalFlow + strategy.currentFlow,
				time:             strategy.time + 1,
				currentValveName: strategy.currentValveName,
				openedValves:     append(openedValvesCopy, strategy.currentValveName),
			}
			heap.Push(h, newStrategy)
		}
		for _, nextValve := range currentValve.tunnels {
			_, alreadyVisited := visitedValves[ValveVisit{nextValve, strategy.countValveVisits() + 1}]
			//oldDistance, onVisitPos := toVisitValves[ValveVisit{nextValve, strategy.countValveVisits() + 1}]
			if alreadyVisited {
				continue
			}

			openedValvesCopy := make([]string, len(strategy.openedValves))
			copy(openedValvesCopy, strategy.openedValves)
			newStrategy := Strategy{
				visitedValves:    append(strategy.visitedValves, nextValve),
				currentFlow:      strategy.currentFlow,
				totalFlow:        strategy.totalFlow + strategy.currentFlow,
				time:             strategy.time + 1,
				currentValveName: nextValve,
				openedValves:     openedValvesCopy,
			}
			heap.Push(h, newStrategy)
		}
	}

	return -1, errors.New("nothing found")
}

func evalA(valves map[string]Valve) int {
	totalFlow, _ := searchStrategy(valves)

	return totalFlow
}

func evalB(lines []string) int {

	return 0
}

func getValves(lines []string) map[string]Valve {
	valves := make(map[string]Valve)

	for _, line := range lines {
		valveName := line[6:8]
		rate, _ := strconv.Atoi(strings.Split(strings.Split(line, "; ")[0], "=")[1])
		tunnels := strings.Split(line, "; ")[1][23:]
		valves[valveName] = Valve{valveName, rate, strings.Split(tunnels, ", ")}
	}

	return valves
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)
	scanOutput := getValves(lines)

	resA := evalA(scanOutput)
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
	day := 16
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
