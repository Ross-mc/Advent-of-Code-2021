package day2

import (
	"strconv"
	"strings"
)

func Task1(instructions []string) int {
	depth, horizontal := calculateDepthAndHorizontalDelta(instructions)
	return depth * horizontal
}

func Task2(instructions []string) int {
	depth, horizontal := calculateDepthAndHorizontalWithAim(instructions)
	return depth * horizontal
}

func calculateMultiplierOfDepthAndHoriztonal(depth int, horizontal int) int {
	return depth * horizontal
}

func calculateDepthAndHorizontalDelta(instructions []string) (depth int, horizontal int) {
	for _, instruction := range instructions {
		direction, delta := getDirectionAndDelta(instruction)
		switch direction {
		case "forward":
			horizontal += delta
		case "down":
			depth += delta
		case "up":
			depth -= delta
		}
	}
	return
}

func getDirectionAndDelta(instruction string) (direction string, delta int) {
	slc := strings.Split(instruction, " ")
	direction = slc[0]
	num, _ := strconv.Atoi(slc[1])
	delta = num
	return
}

func calculateDepthAndHorizontalWithAim(instructions []string) (depth int, horizontal int) {
	aim := 0
	for _, instruction := range instructions {
		direction, delta := getDirectionAndDelta(instruction)
		switch direction {
		case "forward":
			horizontal += delta
			depth += (delta * aim)
		case "down":
			aim += delta
		case "up":
			aim -= delta
		}
	}
	return
}
