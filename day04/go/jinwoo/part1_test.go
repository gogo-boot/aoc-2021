package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var dat, _ = os.ReadFile("input.txt")
var line = strings.Split(string(dat[:]), "\n")
var calledNumber, boardMap = GenBoard(&line)

func TestPart1(t *testing.T) {
	justCalledNumber, board := searchBingoBoard(calledNumber, &boardMap)
	sum := GetSumOfBoard(board)
	fmt.Printf("Part1 Answer : %d \n", sum*justCalledNumber)
}

func GetSumOfBoard(board *Board) int {
	sum := 0
	for i := 0; i < len(board.cells); i++ {
		cell := board.cells
		for j := 0; j < len(cell[i]); j++ {
			if !cell[i][j].isChosen {
				sum = sum + cell[i][j].value
			}
		}
	}
	return sum
}

type Board struct {
	cells [][]Cell
}
type Cell struct {
	value    int
	isChosen bool
}

func searchBingoBoard(number *[]int, boardMap *map[int]Board) (num int, board *Board) {
	for i := 0; i < len(*number); i++ {
		chooseNumber((*number)[i], boardMap)
		if is, _, board := HasBingo(boardMap); is {
			return (*number)[i], board
		}
	}
	return
}

func HasBingo(boardMap *map[int]Board) (hasBingo bool, boardIdx int, board *Board) {

	for i := 0; i < len(*boardMap); i++ {
		board := (*boardMap)[i]
		for j := 0; j < len(board.cells); j++ {
			colBingCount := 0
			rowBingCount := 0
			for k := 0; k < len(board.cells[j]); k++ {
				if board.cells[j][k].isChosen {
					rowBingCount++
				}
				if board.cells[k][j].isChosen {
					colBingCount++
				}
				if rowBingCount == 5 || colBingCount == 5 {
					return true, i, &board
				}
			}
		}
	}
	return false, 0, board
}

func chooseNumber(number int, boardMap *map[int]Board) {
	for i := 0; i < len(*boardMap); i++ {
		board := (*boardMap)[i]
		for j := 0; j < len(board.cells); j++ {
			for k := 0; k < len(board.cells[j]); k++ {
				if board.cells[j][k].value == number {
					board.cells[j][k].isChosen = true
				}
			}
		}
	}
}

func GenBoard(input *[]string) (calledNumber *[]int, boardMap map[int]Board) {
	retmap := make(map[int]Board)
	for i := 2; i < len(*input); i = i + 6 {
		var cells [][]Cell
		for j := 0; j < 5; j++ {
			line := (*input)[i+j]
			cells = append(cells, getCellSlicePerLine(line))
		}
		retmap[len(retmap)] = Board{cells}
	}
	numbersInString := strings.Split((*input)[0], ",")
	number := make([]int, len(numbersInString))
	for i, s := range numbersInString {
		number[i], _ = strconv.Atoi(s)
	}
	return &number, retmap
}

func getCellSlicePerLine(line string) (retCell []Cell) {
	numbersInString := strings.Fields(line)
	number := make([]int, len(numbersInString))

	for i, s := range numbersInString {
		number[i], _ = strconv.Atoi(s)
	}

	for i := 0; i < len(numbersInString); i++ {
		retCell = append(retCell, Cell{number[i], false})
	}
	return
}
