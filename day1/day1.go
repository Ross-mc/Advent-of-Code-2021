package day1

import (
	"strconv"
)

func CalculateNumOfMeasurementsBiggerThanPrev(slc []string) (count int) {
	for i := 1; i < len(slc); i++ {
		prevNum, _ := strconv.Atoi(slc[i-1])
		currNum, _ := strconv.Atoi(slc[i])
		if currNum > prevNum {
			count++
		}
	}

	return
}

func calcSumOfMeasurements(slc []string) (sum int) {
	for _, str := range slc {
		num, _ := strconv.Atoi(str)
		sum += num
	}

	return
}

func CalculateThreeMeasurementSlidingWindowBiggerThanPrevWindow(slc []string) (count int) {
	prevSum := calcSumOfMeasurements(slc[0:3])
	for i := 1; i+2 < len(slc); i += 1 {
		currSum := calcSumOfMeasurements(slc[i : i+3])
		if currSum > prevSum {
			count++
		}
		prevSum = currSum
	}
	return
}
