package day9

import (
	"sort"
	"strconv"
	"strings"
)

type Row []int

type Numbers []Row

type LowPoint struct {
	value int
	row   int
	col   int
}

type NumberInBasin struct {
	value int
	row   int
	col   int
}
type BasinNumbers []NumberInBasin

func parseInput(input []string) Numbers {
	// 3987894921
	// 9856789892
	// 8767896789
	numbers := Numbers{}

	for _, str := range input {
		split := strings.Split(str, "")
		row := Row{}
		for _, char := range split {
			num, _ := strconv.Atoi(char)
			row = append(row, num)
		}
		numbers = append(numbers, row)
	}
	return numbers
}

func calculateSumOfLowPoints(heightmap Numbers) int {
	sum := 0
	for rowIdx, row := range heightmap {
		for colIdx, num := range row {
			if locationIsLowPoint(heightmap, num, rowIdx, colIdx) {
				sum += num + 1
			}
		}
	}
	return sum
}

func locationIsLowPoint(heightmap Numbers, num int, rowIdx int, colIdx int) bool {
	if upIsHigher(heightmap, num, rowIdx, colIdx) && downIsHigher(heightmap, num, rowIdx, colIdx) && rightIsHigher(heightmap, num, rowIdx, colIdx) && leftIsHigher(heightmap, num, rowIdx, colIdx) {
		return true
	}
	return false
}

func upIsHigher(heightmap Numbers, num int, rowIdx int, colIdx int) bool {
	if rowIdx == 0 {
		return true
	}
	upperNum := heightmap[rowIdx-1][colIdx]

	return upperNum > num
}

func downIsHigher(heightmap Numbers, num int, rowIdx int, colIdx int) bool {
	if rowIdx == len(heightmap)-1 {
		return true
	}
	lowerNum := heightmap[rowIdx+1][colIdx]

	return lowerNum > num
}

func rightIsHigher(heightmap Numbers, num int, rowIdx int, colIdx int) bool {
	if colIdx == len(heightmap[0])-1 {
		return true
	}
	rightNum := heightmap[rowIdx][colIdx+1]

	return rightNum > num
}

func leftIsHigher(heightmap Numbers, num int, rowIdx int, colIdx int) bool {
	if colIdx == 0 {
		return true
	}
	leftNum := heightmap[rowIdx][colIdx-1]

	return leftNum > num
}

func getAllLowPoints(heightmap Numbers) []LowPoint {
	lowPoints := []LowPoint{}
	for rowIdx, row := range heightmap {
		for colIdx, num := range row {
			if locationIsLowPoint(heightmap, num, rowIdx, colIdx) {
				point := LowPoint{
					value: num,
					row:   rowIdx,
					col:   colIdx,
				}
				lowPoints = append(lowPoints, point)
			}
		}
	}
	return lowPoints
}

func generateUDLR(currNum NumberInBasin, heightmap Numbers) (up NumberInBasin, down NumberInBasin, left NumberInBasin, right NumberInBasin) {
	up = NumberInBasin{
		row: currNum.row + 1,
		col: currNum.col,
	}
	down = NumberInBasin{
		row: currNum.row - 1,
		col: currNum.col,
	}
	left = NumberInBasin{
		row: currNum.row,
		col: currNum.col - 1,
	}

	right = NumberInBasin{
		row: currNum.row,
		col: currNum.col + 1,
	}
	//setting the value to 9 so it doesnt get check if its out of range as 9 is not added
	if currNum.row == len(heightmap)-1 {
		up.value = 9
	} else {
		up.value = heightmap[currNum.row+1][currNum.col]
	}
	if currNum.col == len(heightmap[0])-1 {
		right.value = 9
	} else {
		right.value = heightmap[currNum.row][currNum.col+1]
	}

	if currNum.col == 0 {
		left.value = 9
	} else {
		left.value = heightmap[currNum.row][currNum.col-1]
	}

	if currNum.row == 0 {
		down.value = 9
	} else {
		down.value = heightmap[currNum.row-1][currNum.col]
	}

	return
}

func checkAdjacentNumbers(currNum NumberInBasin, basin *BasinNumbers, heightmap Numbers) bool {
	up, down, left, right := generateUDLR(currNum, heightmap)
	returnValue := false
	if up.value != 9 && !pointAlreadyInBasin(up, *basin) {
		(*basin) = append((*basin), up)
		returnValue = true
	}
	if down.value != 9 && !pointAlreadyInBasin(down, *basin) {
		(*basin) = append((*basin), down)
		returnValue = true
	}
	if left.value != 9 && !pointAlreadyInBasin(left, *basin) {
		(*basin) = append((*basin), left)
		returnValue = true
	}
	if right.value != 9 && !pointAlreadyInBasin(right, *basin) {
		(*basin) = append((*basin), right)
		returnValue = true
	}
	return returnValue
}

func getBasinSize(lowPoint LowPoint, heightmap Numbers) int {
	init := NumberInBasin{
		value: lowPoint.value,
		row:   lowPoint.row,
		col:   lowPoint.col,
	}
	basin := BasinNumbers{init}
	for i := 0; i < len(basin); i++ {
		number := basin[i]
		checkAdjacentNumbers(number, &basin, heightmap)
	}
	return len(basin)
}

func pointAlreadyInBasin(point NumberInBasin, basin BasinNumbers) bool {
	for _, testPoint := range basin {
		if point.col == testPoint.col && point.row == testPoint.row {
			return true
		}
	}
	return false
}

func Task1(input []string) int {
	parsed := parseInput(input)
	return calculateSumOfLowPoints(parsed)
}

func Task2(input []string) int {
	parsed := parseInput(input)
	lowPoints := getAllLowPoints(parsed)
	largestBasins := []int{0, 0, 0}
	for _, lowPoint := range lowPoints {
		basinSize := getBasinSize(lowPoint, parsed)
		if basinSize > largestBasins[0] {
			largestBasins[0] = basinSize
			sort.Ints(largestBasins)
		}
	}
	return largestBasins[0] * largestBasins[1] * largestBasins[2]
}
