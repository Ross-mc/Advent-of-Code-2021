package day6

import (
	"strconv"
	"strings"
)

type LanternFishSchool []int

type CleverSchool map[int]int

func parseInput(input []string) LanternFishSchool {
	inputStr := input[0]
	csv := strings.Split(inputStr, ",")
	school := LanternFishSchool{}
	for _, strInt := range csv {
		parsed, _ := strconv.Atoi(strInt)
		school = append(school, parsed)
	}
	return school
}

func itsANewDawnItsaNewDayAndImFeelingLanternFish(school *LanternFishSchool) {
	freshFish := LanternFishSchool{}
	for i := 0; i < len(*school); i++ {
		if (*school)[i] == 0 {
			(*school)[i] = 6
			freshFish = append(freshFish, 8)
		} else {
			(*school)[i]--
		}
	}
	(*school) = append((*school), freshFish...)
}

func runItBack80Times(school *LanternFishSchool) {
	for i := 0; i < 80; i++ {
		itsANewDawnItsaNewDayAndImFeelingLanternFish(school)
	}
}

func runItBack256Times(initialSchool *LanternFishSchool) CleverSchool {
	var cleverSchool = CleverSchool{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, num := range *initialSchool {
		cleverSchool[num] += 1
	}
	for i := 0; i < 256; i++ {
		holdZeroCount := cleverSchool[0]
		cleverSchool[0] = cleverSchool[1]
		cleverSchool[1] = cleverSchool[2]
		cleverSchool[2] = cleverSchool[3]
		cleverSchool[3] = cleverSchool[4]
		cleverSchool[4] = cleverSchool[5]
		cleverSchool[5] = cleverSchool[6]
		cleverSchool[6] = holdZeroCount + cleverSchool[7]
		cleverSchool[7] = cleverSchool[8]
		cleverSchool[8] = holdZeroCount
	}
	return cleverSchool
}

func countTheFishIntelligently(cleverSchool CleverSchool) int {
	count := 0
	for _, fish := range cleverSchool {
		count += fish
	}
	return count
}

func countTheFish(school LanternFishSchool) int {
	return len(school)
}

func Task1(input []string) int {
	school := parseInput(input)
	runItBack80Times(&school)
	return countTheFish(school)
}

func Task2(input []string) int {
	school := parseInput(input)
	finalSchool := runItBack256Times(&school)
	return countTheFishIntelligently(finalSchool)
}
