package day7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Counter map[int]int

type Nums []int

var nums = Nums{}

func parseInput(input []string) Nums {
	inputStr := input[0]
	csv := strings.Split(inputStr, ",")
	nums := Nums{}
	for _, strInt := range csv {
		parsed, _ := strconv.Atoi(strInt)
		nums = append(nums, parsed)
	}
	return nums
}

func findMedian(nums Nums) int {
	sort.Ints(nums)
	length := len(nums)
	if length%2 == 0 {
		return nums[length/2]
	} else {
		rounded := math.Floor(float64(length / 2))
		return nums[int(rounded)]
	}
}

func totalUpDifferences(nums Nums, mostCommon int) int {
	totalDifference := 0
	for _, num := range nums {
		difference := num - mostCommon
		if difference < 0 {
			difference *= -1
		}
		totalDifference += difference
	}
	return totalDifference
}

func findMaxAndMin(nums Nums) (max int, min int) {
	max = math.MinInt32
	min = math.MaxInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return
}

func findBestCost(max int, min int, nums Nums) int {
	bestCost := math.MaxInt32
	iterations := 0
	for i := min; i < max; i++ {
		testCost := totalDifferenceWithIncreasingCost(nums, i)
		if testCost < bestCost {
			bestCost = testCost
		}
		iterations++
	}
	return bestCost
}

func checkIfAnyValuesAreAlreadyAtTarget(nums Nums, target int) int {
	hitTarget := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			hitTarget++
		}
	}
	return hitTarget
}

func loopThroughCurrentValuesAndMoveCrabsAtCost(nums *Nums, target int, totalCost *int, hitTarget *int, iteration int) {
	currNums := (*nums)
	for j := 0; j < len(currNums); j++ {
		if currNums[j] == target {
			continue
		}
		currCostToMove := iteration + 1
		if currNums[j] < target {
			currNums[j]++
			(*totalCost) += currCostToMove
		} else if currNums[j] > target {
			currNums[j]--
			(*totalCost) += currCostToMove
		}
		if currNums[j] == target {
			(*hitTarget)++
		}
	}
}

func totalDifferenceWithIncreasingCost(nums Nums, target int) int {
	copiedSlc := make(Nums, len(nums))
	copy(copiedSlc, nums)
	totalCost := 0
	hitTarget := checkIfAnyValuesAreAlreadyAtTarget(copiedSlc, target)
	for i := 0; hitTarget < len(copiedSlc); i++ {
		loopThroughCurrentValuesAndMoveCrabsAtCost(&copiedSlc, target, &totalCost, &hitTarget, i)
	}
	return totalCost
}

func Task1(input []string) int {
	nums = parseInput(input)
	median := findMedian(nums)
	difference := totalUpDifferences(nums, median)
	return difference
}

func Task2(input []string) int {
	max, min := findMaxAndMin(nums)
	cost := findBestCost(max, min, nums)
	return cost
}
