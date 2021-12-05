package day5

import (
	"strconv"
	"strings"
)

type Board [][]int

type Start struct {
	x1 int
	y1 int
}

type End struct {
	x2 int
	y2 int
}

type MoveSet struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Instructions []MoveSet

func generateBlankBoard(limiter int) Board {
	board := Board{}
	for i := 0; i < limiter; i++ {
		temp := make([]int, limiter)
		board = append(board, temp)
	}
	return board
}

func filterHorizontalAndVerticalLines(instructions Instructions) (straight Instructions, diagonal Instructions) {
	for _, instruction := range instructions {
		if (instruction.x1 == instruction.x2) || (instruction.y1 == instruction.y2) {
			straight = append(straight, instruction)
		} else {
			diagonal = append(diagonal, instruction)
		}
	}
	return straight, diagonal
}

func parseInput(input []string) Instructions {
	result := Instructions{}
	for _, moveInstuction := range input {
		splitOnArrow := strings.Split(moveInstuction, "->")
		//[["100,150", "200,150"]]
		start := strings.Split(splitOnArrow[0], ",")
		end := strings.Split(splitOnArrow[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])
		moveset := MoveSet{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
		result = append(result, moveset)
	}
	return result
}

func traverseVertical(board *Board, instruction MoveSet) {
	if instruction.y1 < instruction.y2 {
		for i := instruction.y1; i <= instruction.y2; i++ {
			(*board)[i][instruction.x1] += 1
		}
	} else {
		for i := instruction.y2; i <= instruction.y1; i++ {
			(*board)[i][instruction.x1] += 1
		}
	}
}

func traverseHorizontal(board *Board, instruction MoveSet) {
	if instruction.x1 < instruction.x2 {
		for i := instruction.x1; i <= instruction.x2; i++ {

			(*board)[instruction.y1][i] += 1
		}
	} else {
		for i := instruction.x2; i <= instruction.x1; i++ {
			(*board)[instruction.y1][i] += 1
		}
	}
}

func moveDownAndRight(start *Start, end End) {
	start.x1 += 1
	start.y1 += 1
}

func moveUpAndLeft(start *Start, end End) {
	start.x1 -= 1
	start.y1 -= 1
}

func moveDownAndLeft(start *Start, end End) {
	start.x1 -= 1
	start.y1 += 1
}

func moveUpAndRight(start *Start, end End) {
	start.x1 += 1
	start.y1 -= 1
}

func traverseDiagonal(board *Board, instruction MoveSet) {
	start := Start{
		x1: instruction.x1,
		y1: instruction.y1,
	}
	(*board)[start.y1][start.x1] += 1
	end := End{
		x2: instruction.x2,
		y2: instruction.y2,
	}
	for (start.x1 != end.x2) || (start.y1 != end.y2) {
		if start.x1 < end.x2 && start.y1 < end.y2 {
			moveDownAndRight(&start, end)
		} else if start.x1 > end.x2 && start.y1 > end.y2 {
			moveUpAndLeft(&start, end)
		} else if start.x1 > end.x2 && start.y1 < end.y2 {
			moveDownAndLeft(&start, end)
		} else {
			moveUpAndRight(&start, end)
		}
		(*board)[start.y1][start.x1] += 1
	}
}

func processHorizontalAndVerticalLines(board *Board, filteredInstructions Instructions) {
	//[
	//	[0,0,0]
	//	[0,0,0]
	//	[0,0,0]
	//]
	//the board is an array of arrays. To move horizontal, find the array (y coord) and move through the array
	//ie keep the array but change index
	//to move vertical, find the column (the x position) and move through the arrays but keep the index the same
	for _, instruction := range filteredInstructions {
		if instruction.x1 == instruction.x2 {
			traverseVertical(board, instruction)
		} else {
			traverseHorizontal(board, instruction)
		}
	}

}

func calculatePointsWithTwoOrMoreHits(board *Board) int {
	count := 0
	for _, row := range *board {
		for _, num := range row {
			if num > 1 {
				count++
			}
		}
	}
	return count
}

func processDiagonalInstructions(board *Board, instructions Instructions) {
	for _, instruction := range instructions {
		traverseDiagonal(board, instruction)
	}
}

func Task1(input []string) int {
	board := generateBlankBoard(1000)
	instructions := parseInput(input)
	straight, _ := filterHorizontalAndVerticalLines(instructions)
	processHorizontalAndVerticalLines(&board, straight)
	result := calculatePointsWithTwoOrMoreHits(&board)
	return result
}

func Task2(input []string) int {
	board := generateBlankBoard(1000)
	instructions := parseInput(input)
	straight, diagonal := filterHorizontalAndVerticalLines(instructions)
	processHorizontalAndVerticalLines(&board, straight)
	processDiagonalInstructions(&board, diagonal)
	result := calculatePointsWithTwoOrMoreHits(&board)
	return result
}
