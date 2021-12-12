package day10

import (
	"sort"
)

type Bracket struct {
	open      rune
	close     rune
	points    int
	openCount int
	t2Points  int
}

type IllegalCharacters []Bracket

var ROUND = Bracket{
	open:     '(',
	close:    ')',
	points:   3,
	t2Points: 1,
}

var SQUARE = Bracket{
	open:     '[',
	close:    ']',
	points:   57,
	t2Points: 2,
}

var CURLY = Bracket{
	open:     '{',
	close:    '}',
	points:   1197,
	t2Points: 3,
}

var ANGLE = Bracket{
	open:     '<',
	close:    '>',
	points:   25137,
	t2Points: 4,
}

var BRACKETS = []*Bracket{&ROUND, &SQUARE, &CURLY, &ANGLE}

func resetOpenCount() {
	for _, bracket := range BRACKETS {
		bracket.openCount = 0
	}
}

func findFirstIllegalCharacterInEachLine(input []string) IllegalCharacters {
	illegal := IllegalCharacters{}
	openCharacters := []rune{}
	var shouldBreak = false
	for _, line := range input {
		for _, char := range line {
			if shouldBreak {
				shouldBreak = false
				break
			}
			for _, pointer := range BRACKETS {
				if char == pointer.open {
					openCharacters = append(openCharacters, pointer.open)
				}
				if char == pointer.close {
					if openCharacters[len(openCharacters)-1] == pointer.open {
						openCharacters = openCharacters[0 : len(openCharacters)-1]
					} else {
						illegal = append(illegal, *pointer)
						shouldBreak = true

					}
				}
			}

		}

	}
	return illegal
}

func getOpenCharacters(line string) []Bracket {
	openCharacters := []Bracket{}
	var shouldBreak = false
	for _, char := range line {
		if shouldBreak {
			shouldBreak = false
			break
		}
		for _, pointer := range BRACKETS {
			if char == pointer.open {
				openCharacters = append(openCharacters, *pointer)
			}
			if char == pointer.close {
				if openCharacters[len(openCharacters)-1].open == pointer.open {
					openCharacters = openCharacters[0 : len(openCharacters)-1]
				} else {
					shouldBreak = true
					return []Bracket{}
				}
			}
		}

	}

	return openCharacters
}

func reverse(slice []Bracket) []Bracket {
	reversed := []Bracket{}
	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}
	return reversed
}

func processEachLine(input []string) int {

	scores := []int{}
	for _, line := range input {
		open := getOpenCharacters(line)
		reversed := reverse(open)
		score := 0
		for _, char := range reversed {
			score *= 5
			score += char.t2Points
		}
		if score > 0 {
			scores = append(scores, score)
		}

	}

	sort.Ints(scores)
	length := len(scores)
	middle := ((length + 1) / 2) - 1
	return scores[middle]
}

func calcTotal(illegal IllegalCharacters) int {
	total := 0
	for _, char := range illegal {
		total += char.points
	}

	return total
}

func Task1(input []string) int {
	illegal := findFirstIllegalCharacterInEachLine(input)
	return calcTotal(illegal)
}

func Task2(input []string) int {
	return processEachLine(input)
}
