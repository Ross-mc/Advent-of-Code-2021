package main

import (
	"fmt"

	"github.com/ross-mc/Advent-of-Code-2021/day1"
	"github.com/ross-mc/Advent-of-Code-2021/day2"
	"github.com/ross-mc/Advent-of-Code-2021/day3"
	"github.com/ross-mc/Advent-of-Code-2021/day4"
	"github.com/ross-mc/Advent-of-Code-2021/day5"
	"github.com/ross-mc/Advent-of-Code-2021/day6"
	"github.com/ross-mc/Advent-of-Code-2021/day7"
	"github.com/ross-mc/Advent-of-Code-2021/utils"
)

func main() {
	// day1Tasks()
	// day2Tasks()
	// day3Tasks()
	// day4Tasks()
	// day5Tasks()
	// day6Tasks()
	day7Tasks()
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

func day3Tasks() {
	day3Input := utils.ReadFileIntoStringSlice("./day3/input.txt")
	d3t1 := day3.Task1(day3Input)
	fmt.Printf("Day 3 Task 1: %v\n", d3t1)
	d3t2 := day3.Task2(day3Input)
	fmt.Printf("Day 3 Task 2: %v\n", d3t2)
}

func day4Tasks() {
	day4Input := utils.ReadFileIntoStringSlice("./day4/input.txt")
	d4t1 := day4.Task1(day4Input)
	fmt.Printf("Day 4 Task 1: %v\n", d4t1)
	d4t2 := day4.Task2(day4Input)
	fmt.Printf("Day 4 Task 2: %v\n", d4t2)
}

func day5Tasks() {
	day5Input := utils.ReadFileIntoStringSlice("./day5/input.txt")
	d5t1 := day5.Task1(day5Input)
	fmt.Printf("Day 5 Task 1: %v\n", d5t1)
	d5t2 := day5.Task2(day5Input)
	fmt.Printf("Day 5 Task 2: %v\n", d5t2)
}

func day6Tasks() {
	day6Input := utils.ReadFileIntoStringSlice("./day6/input.txt")
	d6t1 := day6.Task1(day6Input)
	fmt.Printf("Day 6 Task 1: %v\n", d6t1)
	d6t2 := day6.Task2(day6Input)
	fmt.Printf("Day 6 Task 2: %v\n", d6t2)
}

func day7Tasks() {
	day7Input := utils.ReadFileIntoStringSlice("./day7/input.txt")
	d7t1 := day7.Task1(day7Input)
	fmt.Printf("Day 7 Task 1: %v\n", d7t1)
	d7t2 := day7.Task2(day7Input)
	fmt.Printf("Day 7 Task 2: %v\n", d7t2)
}
