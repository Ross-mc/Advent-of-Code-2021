package day14

import (
	"math"
	"strings"
)

var RULES = make(map[string]Rule)

type Rule struct {
	pair      string
	insertion string
}

var PAIRS = make(map[string]Pair)

type Pair struct {
	letters    string
	count      int
	added      int
	subtracted int
}

type CountMap map[string]int

var last = ""

func parseInput(input []string) {

	for i := 1; i < len(input); i++ {
		split := strings.Split(input[i], " -> ")
		pairStr := split[0]
		rule := Rule{
			pair:      pairStr,
			insertion: split[1],
		}
		pair := Pair{
			letters: pairStr,
			count:   0,
		}
		RULES[pairStr] = rule
		PAIRS[pairStr] = pair
	}
	split := strings.Split(input[0], "")
	last = split[len(split)-1]
	for i := 0; i < len(split)-1; i++ {
		pairStr := split[i] + split[i+1]
		pair := PAIRS[pairStr]
		pair.count = 1
		PAIRS[pairStr] = pair
	}

}

func updateCounts() {
	for k := range PAIRS {
		pair := PAIRS[k]
		pair.count = pair.count + pair.added - pair.subtracted
		pair.added = 0
		pair.subtracted = 0
		PAIRS[k] = pair
	}
}

func getTargetPairsToUpdate(pair string, count int) {
	insertion := RULES[pair].insertion
	split := strings.Split(pair, "")
	for idx, ltr := range split {
		pairStr := ltr + insertion
		if idx == 1 {
			pairStr = insertion + ltr
		}
		pairToUpdate := PAIRS[pairStr]
		pairToUpdate.added += count
		PAIRS[pairStr] = pairToUpdate
	}
}

func updateOriginalPairToSubtract(pair string, count int) {
	orig := PAIRS[pair]
	orig.subtracted += count
	PAIRS[pair] = orig
}

func processStep() {
	for key, value := range PAIRS {
		if !(value.count == 0) {
			getTargetPairsToUpdate(key, value.count)
			updateOriginalPairToSubtract(key, value.count)
		}
	}
	updateCounts()
}

func countLeadingLetterOfEachPair() CountMap {
	COUNT_MAP := make(CountMap)
	for _, pair := range PAIRS {
		split := strings.Split(pair.letters, "")
		ltr1 := split[0]

		_, exists := COUNT_MAP[ltr1]
		if exists {
			COUNT_MAP[ltr1] += pair.count
		} else {
			COUNT_MAP[ltr1] = pair.count
		}
	}
	addLastLetterOfInitialString(&COUNT_MAP)
	return COUNT_MAP
}

func addLastLetterOfInitialString(cm *CountMap) {
	(*cm)[last]++
}

func findMaxAndMinValues(cm CountMap) (max int, min int) {
	max = math.MinInt
	min = math.MaxInt

	for _, v := range cm {
		if v > max {
			max = v
		}

		if v < min && v != 0 {
			min = v
		}
	}
	return
}

func getMinAndMaxDifference() int {

	COUNT_MAP := countLeadingLetterOfEachPair()
	max, min := findMaxAndMinValues(COUNT_MAP)
	return max - min
}

func Task(input []string) (task1 int, task2 int) {

	parseInput(input)
	for i := 0; i < 40; i++ {
		if i == 10 {
			task1 = getMinAndMaxDifference()
		}
		processStep()

	}
	task2 = getMinAndMaxDifference()
	return
}
