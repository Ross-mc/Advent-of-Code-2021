package day14

import (
	"fmt"
	"math"
	"strings"
)

var POLYMER = ""

var START = ""

var RULES []Rule

type Rule struct {
	pair      string
	insertion string
}

type LetterCount struct {
	char  string
	count int
}

type Pair struct {
	letters string
	count   int
}

func parseInput(input []string) {
	POLYMER = input[0]
	START = input[0]
	for i := 1; i < len(input); i++ {
		split := strings.Split(input[i], " -> ")
		rule := Rule{
			pair:      split[0],
			insertion: split[1],
		}
		if !strings.Contains(START, rule.insertion) {
			START += rule.insertion
		}
		RULES = append(RULES, rule)
	}
}

func getPairs() []string {
	pairs := []string{}
	split := strings.Split(POLYMER, "")
	for i := 0; i < len(split)-1; i++ {
		pair := split[i] + split[i+1]
		pairs = append(pairs, pair)
	}
	return pairs
}

func processStep() {
	holding := ""
	pairs := getPairs()
	for idx, pair := range pairs {
		split := strings.Split(pair, "")
		first := split[0]
		second := split[1]
		for _, rule := range RULES {
			if pair == rule.pair {
				holding += first + rule.insertion
			}
		}
		if idx == len(pairs)-1 {
			holding += second
		}
	}
	POLYMER = holding
}

func getMinAndMaxDifference() int {
	letters := strings.Split(START, "")
	letterCounts := []LetterCount{}
	for _, letter := range letters {
		lc := LetterCount{
			char:  letter,
			count: strings.Count(POLYMER, letter),
		}
		letterCounts = append(letterCounts, lc)
	}
	max := math.MinInt
	min := math.MaxInt
	for _, lc := range letterCounts {
		if lc.count > max {
			max = lc.count
		}
		if lc.count < min {
			min = lc.count
		}
	}
	fmt.Println(letterCounts)
	return max - min
}

func Task1(input []string) int {

	parseInput(input)
	for i := 0; i < 10; i++ {
		processStep()
	}
	return getMinAndMaxDifference()
}
