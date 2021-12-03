package day3

import (
	"strconv"
	"strings"
)

type BinaryCounter map[string]int

type BinaryCounterByPosition map[int]BinaryCounter

func createBinaryCounterIfDoesNotExist(idx int, countOfDigits *BinaryCounterByPosition) {
	_, idxExists := (*countOfDigits)[idx]
	if !idxExists {
		(*countOfDigits)[idx] = make(BinaryCounter)
	}
}

func countOfBinaryDigitsByPosition(binaryDigits []string) BinaryCounterByPosition {
	var countOfDigits = make(BinaryCounterByPosition)
	for _, digits := range binaryDigits {
		//looping over each set of bindary digits ie 100110
		for idx, digit := range digits {
			createBinaryCounterIfDoesNotExist(idx, &countOfDigits)
			//looping through each individual digit to get the index and the digit itself
			strDigit := string(digit)
			//digit is a rune so converting to string
			_, exists := countOfDigits[idx][strDigit]
			if !exists {
				countOfDigits[idx][strDigit] = 0
			}
			//incrementing the count by 1
			countOfDigits[idx][strDigit] = countOfDigits[idx][strDigit] + 1
		}
	}
	//by the end of the loop we have a map with all the indexes and their counts of 1s and 0s
	return countOfDigits
}

func calculateTheGammaAndEpsilonBinaryValues(binaryDigits []string) (gamma string, epsilon string) {
	countOfDigits := countOfBinaryDigitsByPosition(binaryDigits)
	gammaArr := make([]string, len(countOfDigits))
	epsilonArr := make([]string, len(countOfDigits))
	for idx, counter := range countOfDigits {
		zeroCount := counter["0"]
		oneCount := counter["1"]
		if zeroCount > oneCount {
			gammaArr[idx] = "0"
			epsilonArr[idx] = "1"
		} else {
			gammaArr[idx] = "1"
			epsilonArr[idx] = "0"
		}
	}
	gamma = strings.Join(gammaArr, "")
	epsilon = strings.Join(epsilonArr, "")
	return
}

func convertBinaryStringToInt(binary string) int64 {
	i, _ := strconv.ParseInt(binary, 2, 64)
	return i
}

func Task1(binaryDigits []string) int64 {
	gammaStr, epsilonStr := calculateTheGammaAndEpsilonBinaryValues(binaryDigits)
	gamma := convertBinaryStringToInt(gammaStr)
	epsilon := convertBinaryStringToInt(epsilonStr)
	return gamma * epsilon
}

// count the digits at index -> filter -> count -> filter etc.

func countDigitsAtIndex(binaryDigits []string, idx int) (countZero int, countOne int) {
	for _, digits := range binaryDigits {
		if digits[idx] == '0' {
			countZero++
		} else {
			countOne++
		}
	}
	return
}

func filterByDigit(binaryDigits []string, idx int, targetRune rune) (filteredSlc []string) {
	for _, digits := range binaryDigits {
		if digits[idx] == byte(targetRune) {
			filteredSlc = append(filteredSlc, digits)
		}
	}
	return
}

func highLowCount(countZero int, countOne int, preference int) (low rune, high rune) {
	if countZero == countOne {
		if preference == 1 {
			return '1', '1'
		}
		return '0', '0'
	}
	if countZero > countOne {
		return '1', '0'
	}
	return '0', '1'
}

func countDigitsAndFilter(binaryDigits []string, highOrLow string) string {
	currentTargetSlc := binaryDigits
	preference := 0
	if highOrLow == "high" {
		preference = 1
	}
	for i := 0; len(currentTargetSlc) > 1; i++ {
		countZero, countOne := countDigitsAtIndex(currentTargetSlc, i)
		low, high := highLowCount(countZero, countOne, preference)
		if highOrLow == "high" {
			currentTargetSlc = filterByDigit(currentTargetSlc, i, high)
		} else {
			currentTargetSlc = filterByDigit(currentTargetSlc, i, low)
		}
	}
	return currentTargetSlc[0]
}

func Task2(binaryDigits []string) int64 {
	oxygenStr := countDigitsAndFilter(binaryDigits, "high")
	co2Str := countDigitsAndFilter(binaryDigits, "low")
	oxygen := convertBinaryStringToInt(oxygenStr)
	co2 := convertBinaryStringToInt(co2Str)
	return oxygen * co2
}
