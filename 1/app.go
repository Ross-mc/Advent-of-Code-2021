package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	timeStart := time.Now().UnixMilli()
	input := readFileIntoStringSlice()
	result := calculateNumOfMeasurementsBiggerThanPrev(input)
	fmt.Println(result)
	result2 := calculateThreeMeasurementSlidingWindowBiggerThanPrevWindow(input)
	fmt.Println(result2)
	timeEnd := time.Now().UnixMilli()
	difference := timeEnd - timeStart
	fmt.Printf("the program took %v ms to run\n", difference)
}

func readFileIntoStringSlice() []string {
	file, _ := os.Open("input.txt")
	bytes, _ := ioutil.ReadAll(file)
	str := string(bytes)
	slice := strings.Split(str, "\n")
	return slice
}

func calculateNumOfMeasurementsBiggerThanPrev(slc []string) (count int) {
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

func calculateThreeMeasurementSlidingWindowBiggerThanPrevWindow(slc []string) (count int) {
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
