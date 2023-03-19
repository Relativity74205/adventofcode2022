package main

import (
	"AoC2022"
	"fmt"
	profile "github.com/pkg/profile"
	"regexp"
	"strconv"
	"strings"
)

type Resource int
type resourceCount struct {
	countOre      int
	countClay     int
	countObsidian int
	countGeode    int
}

const (
	Ore Resource = iota
	Clay
	Obsidian
	Geode
)

func (r *Resource) String() string {
	return []string{"ore", "clay", "obsidian"}[*r]
}

type Blueprint struct {
	id                int
	maxNumberGeodes   int
	costOreRobot      resourceCount
	costClayRobot     resourceCount
	costObsidianRobot resourceCount
	costGeodeRobot    resourceCount
	maxNeededOre      int
	maxNeededClay     int
	maxNeededObsidian int
}

func getCosts(stringPart string) resourceCount {
	return resourceCount{
		countOre:      getCost(stringPart, Ore),
		countClay:     getCost(stringPart, Clay),
		countObsidian: getCost(stringPart, Obsidian),
	}
}

func getCost(stringPart string, resource Resource) int {
	var cost int
	oreRegex := regexp.MustCompile(`.* (\d+) ` + regexp.QuoteMeta(resource.String()))
	match := oreRegex.FindStringSubmatch(stringPart)
	if len(match) > 0 {
		cost, _ = strconv.Atoi(match[1])
	} else {
		cost = 0
	}
	return cost
}

func parseBlueprint(line string, id int) Blueprint {
	relevantPart := strings.Split(line, ": ")[1]
	parts := strings.Split(relevantPart, ". ")
	costsOreRobot := getCosts(parts[0])
	costsClayRobot := getCosts(parts[1])
	costsObsidianRobot := getCosts(parts[2])
	costsGeodeRobot := getCosts(parts[3])
	maxNeededOre := util.MaxIntegers(costsOreRobot.countOre, costsClayRobot.countOre, costsObsidianRobot.countOre, costsGeodeRobot.countOre)
	maxNeededClay := util.MaxIntegers(costsOreRobot.countClay, costsClayRobot.countClay, costsObsidianRobot.countClay, costsGeodeRobot.countClay)
	maxNeededObsidian := util.MaxIntegers(costsOreRobot.countObsidian, costsClayRobot.countObsidian, costsObsidianRobot.countObsidian, costsGeodeRobot.countObsidian)

	return Blueprint{
		id:                id,
		costOreRobot:      costsOreRobot,
		costClayRobot:     costsClayRobot,
		costObsidianRobot: costsObsidianRobot,
		costGeodeRobot:    costsGeodeRobot,
		maxNeededOre:      maxNeededOre,
		maxNeededClay:     maxNeededClay,
		maxNeededObsidian: maxNeededObsidian,
	}
}

type SimulationState struct {
	minute     int
	maxMinute  int
	production resourceCount
	resources  resourceCount
}

func (s *SimulationState) isFinished() bool {
	return s.minute >= s.maxMinute
}

func nextSimulationRound(simulationState *SimulationState, blueprint *Blueprint, openStates *statesDataStructure, results *Results) {
	buildRobot(Ore, &blueprint.costOreRobot, simulationState, blueprint, openStates, results)
	buildRobot(Clay, &blueprint.costClayRobot, simulationState, blueprint, openStates, results)
	buildRobot(Obsidian, &blueprint.costObsidianRobot, simulationState, blueprint, openStates, results)
	buildRobot(Geode, &blueprint.costGeodeRobot, simulationState, blueprint, openStates, results)
}

func canBeBuild(availableResources *resourceCount, costs *resourceCount) bool {
	if availableResources.countOre < costs.countOre {
		return false
	}
	if availableResources.countClay < costs.countClay {
		return false
	}
	if availableResources.countObsidian < costs.countObsidian {
		return false
	}

	return true
}

func payResource(availableResources *resourceCount, costs *resourceCount) {
	availableResources.countOre -= costs.countOre
	availableResources.countClay -= costs.countClay
	availableResources.countObsidian -= costs.countObsidian
	availableResources.countGeode -= costs.countGeode
}

func addResource(availableResources *resourceCount, harvest *resourceCount) {
	availableResources.countOre += harvest.countOre
	availableResources.countClay += harvest.countClay
	availableResources.countObsidian += harvest.countObsidian
	availableResources.countGeode += harvest.countGeode
}

func harvestResource(production *resourceCount) *resourceCount {
	return &resourceCount{production.countOre,
		production.countClay,
		production.countObsidian,
		production.countGeode}
}

func increaseProduction(production *resourceCount, robotType Resource) {
	switch robotType {
	case Ore:
		production.countOre += 1
	case Clay:
		production.countClay += 1
	case Obsidian:
		production.countObsidian += 1
	case Geode:
		production.countGeode += 1
	}
}

func buildRobot(robotType Resource, costs *resourceCount, simulationState *SimulationState, blueprint *Blueprint, openStates *statesDataStructure, results *Results) {
	minute := simulationState.minute
	availableResources := resourceCount{
		countOre:      simulationState.resources.countOre,
		countClay:     simulationState.resources.countClay,
		countObsidian: simulationState.resources.countObsidian,
		countGeode:    simulationState.resources.countGeode,
	}
	for minute < simulationState.maxMinute {
		minute++
		harvest := harvestResource(&simulationState.production)
		if canBeBuild(&availableResources, costs) && !maxNeededProductionReached(robotType, &simulationState.production, blueprint) {

			production := resourceCount{
				countOre:      simulationState.production.countOre,
				countClay:     simulationState.production.countClay,
				countObsidian: simulationState.production.countObsidian,
				countGeode:    simulationState.production.countGeode,
			}
			increaseProduction(&production, robotType)
			payResource(&availableResources, costs)
			addResource(&availableResources, harvest)
			state := SimulationState{minute: minute, maxMinute: simulationState.maxMinute, production: production, resources: availableResources}
			openStates.queue = append([]*SimulationState{&state}, openStates.queue...)
			return
		}
		addResource(&availableResources, harvest)
	}

	results.maxGeodes = util.MaxInt(results.maxGeodes, availableResources.countGeode)
}

func maxNeededProductionReached(robotType Resource, production *resourceCount, blueprint *Blueprint) bool {
	switch robotType {
	case Ore:
		return production.countOre >= blueprint.maxNeededOre
	case Clay:
		return production.countClay >= blueprint.maxNeededClay
	case Obsidian:
		return production.countObsidian >= blueprint.maxNeededObsidian
	}
	return false
}

type statesDataStructure struct {
	queue       []*SimulationState
	alreadySeen map[SimulationState]bool
}

type Results struct {
	maxGeodes int
}

func getMaxNumberOfGeodesFromBlueprint(blueprint *Blueprint, maxMinute int) int {
	results := Results{}

	startState := SimulationState{
		minute:     0,
		maxMinute:  maxMinute,
		production: resourceCount{countOre: 1, countClay: 0, countObsidian: 0, countGeode: 0},
		resources:  resourceCount{countOre: 0, countClay: 0, countObsidian: 0, countGeode: 0},
	}
	openStates := statesDataStructure{
		queue:       []*SimulationState{&startState},
		alreadySeen: map[SimulationState]bool{startState: true},
	}

	for len(openStates.queue) > 0 {
		state := openStates.queue[0]
		openStates.queue = openStates.queue[1:]
		nextSimulationRound(state, blueprint, &openStates, &results)
	}

	return results.maxGeodes
}

func evalA(lines []string) int {
	var blueprints []Blueprint
	for id, line := range lines {
		blueprints = append(blueprints, parseBlueprint(line, id+1))
	}

	var sumQualityLevel int
	for _, blueprint := range blueprints {
		blueprint.maxNumberGeodes = getMaxNumberOfGeodesFromBlueprint(&blueprint, 24)
		//fmt.Printf("blueprint.maxNumberGeodes = %d\n", blueprint.maxNumberGeodes)
		sumQualityLevel += blueprint.id * blueprint.maxNumberGeodes
	}

	return sumQualityLevel
}

func evalB(lines []string) int {
	var blueprints []Blueprint
	maxBlueprints := util.MinInt(len(lines), 3)
	for id, line := range lines[:maxBlueprints] {
		blueprints = append(blueprints, parseBlueprint(line, id+1))
	}

	multipliedMaxNumberGeodes := 1
	for _, blueprint := range blueprints {
		blueprint.maxNumberGeodes = getMaxNumberOfGeodesFromBlueprint(&blueprint, 32)
		//fmt.Printf("blueprint.maxNumberGeodes = %d\n", blueprint.maxNumberGeodes)
		multipliedMaxNumberGeodes *= blueprint.maxNumberGeodes
	}

	return multipliedMaxNumberGeodes
}

func eval(filename string, debug bool) {
	lines := util.ReadFile(filename)

	resA := evalA(lines)
	if debug {
		fmt.Printf("A (debug): %v \n", resA)
	} else {
		fmt.Printf("A: %v \n", resA)
	}
	//resB := evalB(lines)
	//if debug {
	//	fmt.Printf("B (debug): %v \n", resB)
	//} else {
	//	fmt.Printf("B: %v \n", resB)
	//}
}

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()
	day := 19
	debugSuffix := "_debug"
	filename := fmt.Sprintf("input%02d.txt", day)
	filenameDebug := fmt.Sprintf("input%02d%v.txt", day, debugSuffix)

	fmt.Printf("Day %02d \n", day)
	eval(filenameDebug, true)
	eval(filename, false)
}
