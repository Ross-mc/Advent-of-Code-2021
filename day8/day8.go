package day8

import (
	"sort"
	"strconv"
	"strings"
)

type SegmentsRequired map[int]int

var segmentsRequired = SegmentsRequired{0: 6, 1: 2, 2: 5, 3: 5, 4: 4, 5: 5, 6: 6, 7: 3, 8: 7, 9: 6}

type Display struct {
	top    string
	tl     string
	tr     string
	middle string
	bl     string
	br     string
	bottom string
}

type Patterns map[int]string

func stringContains(str string, char string) bool {
	for _, c := range str {
		if string(c) == char {
			return true
		}
	}
	return false
}

func everyStringContains(strings []string, char string) bool {
	for _, str := range strings {
		if stringContains(str, char) == false {
			return false
		}
	}
	return true
}

func parseInput(input []string) (patterns []string, output []string) {
	for _, line := range input {
		split := strings.Split(line, " | ")
		patterns = append(patterns, split[0])
		output = append(output, split[1])
	}
	return
}

func findNumOfEasyDigits(output []string) int {
	//"easy" digits to find are those that require a unique number of segments
	//in order to build and are therefore easy to work out
	//the unique digits are 1,4,7 and 8

	count := 0

	for _, outputString := range output {
		digits := strings.Split(outputString, " ")
		for _, digit := range digits {
			digitLength := len(digit)
			if (digitLength == segmentsRequired[1]) ||
				(digitLength == segmentsRequired[4]) ||
				(digitLength == segmentsRequired[7]) ||
				(digitLength == segmentsRequired[8]) {
				count++
			}
		}
	}

	return count
}

func getTopLetter(onePattern string, sevenPattern string) string {
	//the seven digit uniquely contains the top segment out of 1 and 7
	for _, c := range sevenPattern {
		if stringContains(onePattern, string(c)) == false {
			return string(c)
		}
	}
	return ""
}

func getMiddleLetter(twoThreeFive []string, fourPattern string) string {
	//the middle segment is the only segment that exists in 2,3,4 and 5
	// we dont know yet which ones 2,3 and 5 are but it doesnt matter as they all contain middle
	var allPatterns = []string{}
	allPatterns = append(allPatterns, twoThreeFive...)
	allPatterns = append(allPatterns, fourPattern)
	letters := ""
	for _, str := range twoThreeFive {
		for _, c := range str {
			if !stringContains(letters, string(c)) {
				letters += string(c)
			}
		}
	}
	for _, letter := range letters {
		if everyStringContains(allPatterns, string(letter)) {
			return string(letter)
		}
	}
	return ""
}

func getTopLeftLetter(onePattern string, fourPattern string, middleLetter string) string {
	//four has top left and middle and everything in one
	//therefore if we middle, and everything in one we can find topleft
	middlePlusOne := onePattern + middleLetter
	for _, c := range fourPattern {
		if !stringContains(middlePlusOne, string(c)) {
			return string(c)
		}
	}
	return ""
}

func getRightHandSideAndFive(onePattern string, twoThreeFive []string, topLeft string) (tr string, br string, fivePattern string) {
	//as we know topleft, 5 is the only one of two, three and five which has the top left
	for _, pattern := range twoThreeFive {
		if stringContains(pattern, topLeft) {
			fivePattern = pattern
		}
	}
	//now compare against 1
	for _, c := range onePattern {
		//one only has tr and br whereas 5 only has br
		//therefore we find the char that is on both 5 and 1 to get br
		//and the otherone must be tr
		if stringContains(fivePattern, string(c)) {
			br = string(c)
		} else {
			tr = string(c)
		}
	}
	return
}

func getTwoAndThree(twoAndThree []string, tr string, br string) (two string, three string) {
	//three has tr and br, two just has tr
	first := twoAndThree[0]
	second := twoAndThree[1]
	if stringContains(first, tr) && stringContains(first, br) {
		three = first
		two = second
	} else {
		three = second
		two = first
	}

	return
}

func getZeroSixAndNine(zeroSixNine []string, middle string, tr string) (zero string, six string, nine string) {
	//zero only one to not have middle
	//six only one to have top right
	for _, pattern := range zeroSixNine {
		if !stringContains(pattern, middle) {
			zero = pattern
		} else if !stringContains(pattern, tr) {
			six = pattern
		} else {
			nine = pattern
		}
	}
	return
}

func decodePatterns(pattern string) Patterns {
	display := Display{}
	patterns := Patterns{0: "", 1: "", 2: "", 3: "", 4: "", 5: "", 6: "", 7: "", 8: "", 9: ""}
	segments := strings.Split(pattern, " ")
	twoThreeFive := []string{}
	zeroSixNine := []string{}
	//1,7,4,8 are unique, 235 are grouped, 069 are grouped
	for _, segment := range segments {
		switch len(segment) {
		case 2:
			patterns[1] = segment
		case 3:
			patterns[7] = segment
		case 4:
			patterns[4] = segment
		case 5:
			twoThreeFive = append(twoThreeFive, segment)
		case 6:
			zeroSixNine = append(zeroSixNine, segment)
		case 7:
			patterns[8] = segment
		}
	}
	display.top = getTopLetter(patterns[1], patterns[7])
	display.middle = getMiddleLetter(twoThreeFive, patterns[4])
	display.tl = getTopLeftLetter(patterns[1], patterns[4], display.middle)
	tr, br, five := getRightHandSideAndFive(patterns[1], twoThreeFive, display.tl)
	display.tr = tr
	display.br = br
	patterns[5] = five
	twoAndThree := []string{}
	for _, str := range twoThreeFive {
		if str != patterns[5] {
			twoAndThree = append(twoAndThree, str)
		}
	}
	two, three := getTwoAndThree(twoAndThree, display.tr, display.br)
	patterns[2] = two
	patterns[3] = three
	zero, six, nine := getZeroSixAndNine(zeroSixNine, display.middle, display.tr)
	patterns[0] = zero
	patterns[6] = six
	patterns[9] = nine
	sortPatterns(&patterns)
	return patterns
}

func sortStr(str string) string {
	slc := strings.Split(str, "")
	sort.Strings(slc)
	return strings.Join(slc, "")
}

func sortPatterns(patterns *Patterns) {
	for key, pattern := range *patterns {
		(*patterns)[key] = sortStr(pattern)
	}
}

func calculateTotal(pattern string, output []string) int {
	total := ""
	decoded := decodePatterns(pattern)
	for _, str := range output {
		for key, value := range decoded {
			if str == value {
				total += strconv.Itoa(key)
			}
		}
	}
	converted, _ := strconv.Atoi(total)
	return converted
}

func Task1(input []string) int {
	_, output := parseInput(input)
	countOfEasyDigits := findNumOfEasyDigits(output)
	return countOfEasyDigits
}

func Task2(input []string) int {
	patterns, output := parseInput(input)
	runningTotal := 0
	for idx, pattern := range patterns {
		currOutput := strings.Split(output[idx], " ")
		for idx, output := range currOutput {
			currOutput[idx] = sortStr(output)
		}
		runningTotal += calculateTotal(pattern, currOutput)
	}
	return runningTotal
}
