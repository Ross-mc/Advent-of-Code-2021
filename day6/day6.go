package day6

import (
	"strconv"
	"strings"
)

type School []int

func parseInput(input []string) School {
	inputStr := input[0]
	csv := strings.Split(inputStr, ",")
	school := School{}
	for _, strInt := range csv {
		parsed, _ := strconv.Atoi(strInt)
		school = append(school, parsed)
	}
	return school
}
func shift(slc *School) (int, School) {
	copy := *slc
	first := copy[0]
	return first, copy[1:]
}

func simulateFishGrowth(initialSchool *School, iterations int) School {
	var school = School{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, num := range *initialSchool {
		school[num] += 1
	}
	for i := 0; i < iterations; i++ {
		prevValue, newSlc := shift(&school)
		school = newSlc
		school[6] += prevValue
		school = append(school, prevValue)
	}
	return school
}

func countTheFish(School School) int {
	count := 0
	for _, fish := range School {
		count += fish
	}
	return count
}

func Task1(input []string) int {
	initial := parseInput(input)
	school := simulateFishGrowth(&initial, 80)
	return countTheFish(school)
}

func Task2(input []string) int {
	initial := parseInput(input)
	school := simulateFishGrowth(&initial, 256)
	return countTheFish(school)
}
