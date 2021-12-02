package main

import (
	"fmt"

	"github.com/ross-mc/Advent-of-Code-2021/day1"
	"github.com/ross-mc/Advent-of-Code-2021/day2"
	"github.com/ross-mc/Advent-of-Code-2021/utils"
)

func main() {
	day1Tasks()
	day2Tasks()
}

func day1Tasks() {
	day1Input := utils.ReadFileIntoStringSlice("./day1/input.txt")
	d1t1 := day1.CalculateNumOfMeasurementsBiggerThanPrev(day1Input)
	fmt.Printf("Day 1 Task 1: %v\n", d1t1)
	d1t2 := day1.CalculateThreeMeasurementSlidingWindowBiggerThanPrevWindow(day1Input)
	fmt.Printf("Day 1 Task 2: %v\n", d1t2)
}

func day2Tasks() {
	day2Input := utils.ReadFileIntoStringSlice("./day2/input.txt")
	d2t1 := day2.Task1(day2Input)
	fmt.Printf("Day 2 Task 1: %v\n", d2t1)
	d2t2 := day2.Task2(day2Input)
	fmt.Printf("Day 2 Task 2: %v\n", d2t2)
}
