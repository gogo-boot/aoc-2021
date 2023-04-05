package main

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	justCalledNumber, board := searchLastBingoBoard(calledNumber, &boardMap)
	sum := GetSumOfBoard(board)
	fmt.Printf("Part2 Answer : %d \n", sum*justCalledNumber)
}

func searchLastBingoBoard(number *[]int, boardMap *map[int]Board) (num int, board *Board) {
	var lastBoard Board
	var lastNum int
	for i := 0; i < len(*number); i++ {
		chooseNumber((*number)[i], boardMap)
		if is, idx, board := HasBingo(boardMap); is {
			lastNum = (*number)[i]
			lastBoard = *board
			delete(*boardMap, idx)
		}
	}
	return lastNum, &lastBoard
}
