package day4

import (
	"strconv"
	"strings"
)

type NumbersDrawn []int

type BingoCard [][]int

type BingoCardCollection []BingoCard

func convertStrSlcToNumSlc(strSlc []string) (numSlc []int) {
	for _, str := range strSlc {
		num, _ := strconv.Atoi(str)
		numSlc = append(numSlc, num)
	}
	return
}

func translateInputIntoNumbersDrawnAndBingoBoard(input []string) (NumbersDrawn, BingoCardCollection) {
	firstStr := input[0]
	strSlc := strings.Split(firstStr, ",")
	numbersDrawn := convertStrSlcToNumSlc(strSlc)
	bingoBoard := BingoCardCollection{}
	for i := 2; i < len(input); i += 6 {
		bingoCard := BingoCard{}
		for j := i; j < i+5; j++ {
			row := strings.Split(input[j], " ")
			num := convertStrSlcToNumSlc(row)
			bingoCard = append(bingoCard, num)
		}
		bingoBoard = append(bingoBoard, bingoCard)
	}
	return numbersDrawn, bingoBoard
}

func contains(numToCheck int, numSlcToCheck NumbersDrawn) bool {
	for _, num := range numSlcToCheck {
		if numToCheck == num {
			return true
		}
	}
	return false
}

func everyContains(curentSelection NumbersDrawn, numSlcToCheck NumbersDrawn) bool {
	for _, num := range curentSelection {
		if contains(num, numSlcToCheck) == false {
			return false
		}
	}
	return true
}

func checkRow(row NumbersDrawn, calledNumbers NumbersDrawn) bool {
	return everyContains(row, calledNumbers)
}
func checkColumn(column NumbersDrawn, calledNumbers NumbersDrawn) bool {
	return everyContains(column, calledNumbers)
}

func checkCard(bingoCard BingoCard, calledNumbers NumbersDrawn) bool {
	for idx, row := range bingoCard {
		isRowComplete := checkRow(row, calledNumbers)
		if isRowComplete {
			return true
		}
		var column = []int{bingoCard[0][idx], bingoCard[1][idx], bingoCard[2][idx], bingoCard[3][idx], bingoCard[4][idx]}
		isColumnComplete := checkColumn(column, calledNumbers)
		if isColumnComplete {
			return true
		}
	}
	return false
}

func playBingo(numbers NumbersDrawn, cards BingoCardCollection) (NumbersDrawn, BingoCard) {
	playedNumbers := NumbersDrawn{}
	for _, num := range numbers {
		playedNumbers = append(playedNumbers, num)
		for _, card := range cards {
			hasCardWon := checkCard(card, playedNumbers)
			if hasCardWon {
				return playedNumbers, card
			}
		}
	}
	return playedNumbers, BingoCard{}
}

func createFlatArr(card BingoCard) []int {
	var flat = []int{}
	for _, numSlc := range card {
		for _, num := range numSlc {
			flat = append(flat, num)
		}
	}
	return flat
}

func sum(numSlc []int) int {
	total := 0
	for _, num := range numSlc {
		total += num
	}
	return total
}

func getUnmarkedNumbersAndMultiplyByFinalNumber(finalNumbers NumbersDrawn, winningCard BingoCard) int {
	flatArr := createFlatArr(winningCard)
	var unmarkedNumbers = []int{}
	for _, num := range flatArr {
		if contains(num, finalNumbers) == false {
			unmarkedNumbers = append(unmarkedNumbers, num)
		}
	}
	total := sum(unmarkedNumbers)
	lastCalled := finalNumbers[len(finalNumbers)-1]
	return total * lastCalled
}

func getIndex(num int, numSlc []int) int {
	for i, v := range numSlc {
		if v == num {
			return i
		}
	}
	return -1
}

func playBingoGetLastWinner(numbers NumbersDrawn, cards BingoCardCollection) (NumbersDrawn, BingoCard) {
	playedNumbers := NumbersDrawn{}
	var winners = []BingoCard{}
	lastWonNum := 0
	for _, num := range numbers {
		playedNumbers = append(playedNumbers, num)
		for idx, card := range cards {
			hasCardWon := checkCard(card, playedNumbers)
			if hasCardWon {
				winners = append(winners, card)
				cards[idx] = BingoCard{{-1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1}}
				lastWonNum = num
			}
		}
	}
	lastWonIdx := getIndex(lastWonNum, playedNumbers)
	playedNumbers = playedNumbers[0 : lastWonIdx+1]
	return playedNumbers, winners[len(winners)-1]
}

func Task1(input []string) int {
	numbersDrawn, cards := translateInputIntoNumbersDrawnAndBingoBoard(input)
	finalNumbers, winningCard := playBingo(numbersDrawn, cards)
	return getUnmarkedNumbersAndMultiplyByFinalNumber(finalNumbers, winningCard)
}

func Task2(input []string) int {
	numbersDrawn, cards := translateInputIntoNumbersDrawnAndBingoBoard(input)
	finalNumbers, lastWinningCard := playBingoGetLastWinner(numbersDrawn, cards)
	return getUnmarkedNumbersAndMultiplyByFinalNumber(finalNumbers, lastWinningCard)
}
