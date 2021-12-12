package day11

import (
	"strconv"
	"strings"
)

type Octopus struct {
	energy       int
	needsToFlash bool
}

type Row []Octopus

type Octopuses []Row

func parseInput(input []string) (octopuses Octopuses) {
	for _, line := range input {
		split := strings.Split(line, "")
		row := Row{}
		for _, strNum := range split {
			num, err := strconv.Atoi(strNum)
			octopus := Octopus{
				energy:       num,
				needsToFlash: false,
			}
			if err == nil {
				row = append(row, octopus)
			}
		}
		octopuses = append(octopuses, row)
	}

	return
}

func updateUp(rowIdx int, colIdx int, octopuses *Octopuses) {
	if rowIdx == 0 {
		return
	}
	(*octopuses)[rowIdx-1][colIdx].energy++
	if (*octopuses)[rowIdx-1][colIdx].energy == 10 {
		(*octopuses)[rowIdx-1][colIdx].needsToFlash = true
	}
}

func updateDown(rowIdx int, colIdx int, octopuses *Octopuses) {
	if rowIdx == len((*octopuses))-1 {
		return
	}
	(*octopuses)[rowIdx+1][colIdx].energy++
	if (*octopuses)[rowIdx+1][colIdx].energy == 10 {
		(*octopuses)[rowIdx+1][colIdx].needsToFlash = true
	}
}

func updateDownAndRight(rowIdx int, colIdx int, octopuses *Octopuses) {
	if rowIdx == len((*octopuses))-1 || colIdx == len((*octopuses)[rowIdx])-1 {
		return
	}
	(*octopuses)[rowIdx+1][colIdx+1].energy++
	if (*octopuses)[rowIdx+1][colIdx+1].energy == 10 {
		(*octopuses)[rowIdx+1][colIdx+1].needsToFlash = true
	}
}

func updateRight(rowIdx int, colIdx int, octopuses *Octopuses) {
	if colIdx == len((*octopuses)[rowIdx])-1 {
		return
	}
	(*octopuses)[rowIdx][colIdx+1].energy++
	if (*octopuses)[rowIdx][colIdx+1].energy == 10 {
		(*octopuses)[rowIdx][colIdx+1].needsToFlash = true
	}
}

func updateUpAndRight(rowIdx int, colIdx int, octopuses *Octopuses) {
	if rowIdx == 0 || colIdx == len((*octopuses)[rowIdx])-1 {
		return
	}
	(*octopuses)[rowIdx-1][colIdx+1].energy++
	if (*octopuses)[rowIdx-1][colIdx+1].energy == 10 {
		(*octopuses)[rowIdx-1][colIdx+1].needsToFlash = true
	}
}

func updateUpAndLeft(rowIdx int, colIdx int, octopuses *Octopuses) {
	if rowIdx == len((*octopuses))-1 || colIdx == 0 {
		return
	}
	(*octopuses)[rowIdx+1][colIdx-1].energy++
	if (*octopuses)[rowIdx+1][colIdx-1].energy == 10 {
		(*octopuses)[rowIdx+1][colIdx-1].needsToFlash = true
	}
}

func updateDownAndLeft(rowIdx int, colIdx int, octopuses *Octopuses) {
	if rowIdx == 0 || colIdx == 0 {
		return
	}
	(*octopuses)[rowIdx-1][colIdx-1].energy++
	if (*octopuses)[rowIdx-1][colIdx-1].energy == 10 {
		(*octopuses)[rowIdx-1][colIdx-1].needsToFlash = true
	}
}

func updateLeft(rowIdx int, colIdx int, octopuses *Octopuses) {
	if colIdx == 0 {
		return
	}
	(*octopuses)[rowIdx][colIdx-1].energy++
	if (*octopuses)[rowIdx][colIdx-1].energy == 10 {
		(*octopuses)[rowIdx][colIdx-1].needsToFlash = true
	}
}

func updateAdjacentOctopuses(rowIdx int, colIdx int, octopuses *Octopuses) {
	updateUp(rowIdx, colIdx, octopuses)

	updateUpAndRight(rowIdx, colIdx, octopuses)

	updateUpAndLeft(rowIdx, colIdx, octopuses)

	updateLeft(rowIdx, colIdx, octopuses)

	updateRight(rowIdx, colIdx, octopuses)

	updateDown(rowIdx, colIdx, octopuses)

	updateDownAndLeft(rowIdx, colIdx, octopuses)

	updateDownAndRight(rowIdx, colIdx, octopuses)
}

func processStep(octopuses *Octopuses) {
	for rowIdx, row := range *octopuses {
		for colIdx := range row {
			(*octopuses)[rowIdx][colIdx].energy += 1
			if (*octopuses)[rowIdx][colIdx].energy == 10 {
				(*octopuses)[rowIdx][colIdx].needsToFlash = true
			}
		}
	}
}

func countFlashes(octopuses *Octopuses) (flashes int) {
	for {
		flashes += flash(octopuses)
		if checkIfAnyNeedToFlash(octopuses) {

			continue
		} else {

			return
		}
	}
}

func reset(octopuses *Octopuses) {
	for i := 0; i < len(*octopuses); i++ {
		for j := 0; j < len((*octopuses)[0]); j++ {
			if (*octopuses)[i][j].energy >= 10 {
				(*octopuses)[i][j].energy = 0
			}
		}
	}
}

func checkIfAllNeedToFlash(octopuses *Octopuses) bool {
	for i := 0; i < len(*octopuses); i++ {
		for j := 0; j < len((*octopuses)[0]); j++ {
			if (*octopuses)[i][j].energy != 0 {
				return false
			}
		}
	}
	return true
}

func checkIfAnyNeedToFlash(octopuses *Octopuses) bool {
	for i := 0; i < len(*octopuses); i++ {
		for j := 0; j < len((*octopuses)[0]); j++ {
			if (*octopuses)[i][j].needsToFlash {
				return true
			}
		}
	}
	return false
}

func flash(octopuses *Octopuses) int {
	flashes := 0
	for rIdx, row := range *octopuses {
		for cIdx, octopus := range row {
			if octopus.needsToFlash {
				updateAdjacentOctopuses(rIdx, cIdx, octopuses)
				flashes++
				(*octopuses)[rIdx][cIdx].needsToFlash = false
			}
		}
	}
	return flashes
}

func processSteps(steps int, octopuses Octopuses) (flashes int) {
	for i := 0; i < steps; i++ {

		processStep(&octopuses)
		flashes += countFlashes(&octopuses)
		reset(&octopuses)
	}
	return flashes
}

func findAllNeedToFlash(octopuses Octopuses) (count int) {
	allNeedToFlash := false
	for !allNeedToFlash {

		processStep(&octopuses)
		countFlashes(&octopuses)
		reset(&octopuses)
		allNeedToFlash = checkIfAllNeedToFlash(&octopuses)
		count++
	}
	return count
}

func Task1(input []string) int {
	parsed := parseInput(input)
	return processSteps(100, parsed)
}

func Task2(input []string) int {
	parsed := parseInput(input)
	return findAllNeedToFlash(parsed)
}
