package main

import (
	"fmt"

	"github.com/ross-mc/Advent-of-Code-2021/day1"
	"github.com/ross-mc/Advent-of-Code-2021/utils"
)

func main() {
	day1Input := utils.ReadFileIntoStringSlice("./day1/input.txt")
	d1t1 := day1.CalculateNumOfMeasurementsBiggerThanPrev(day1Input)
	fmt.Printf("Day 1 Task 1: %v\n", d1t1)
	d1t2 := day1.CalculateThreeMeasurementSlidingWindowBiggerThanPrevWindow(day1Input)
	fmt.Printf("Day 1 Task 2: %v\n", d1t2)
}
