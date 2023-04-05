package main

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	input, foldInstruction := parseInput(&line)
	for _, foldStr := range foldInstruction {
		fold(input, foldStr)
	}
	printStatus(input)
}

func printStatus(input *[][]bool) {
	for x := 0; x < len(*input); x++ {
		for y := 0; y < len((*input)[0]); y++ {
			if (*input)[x][y] {
				fmt.Printf("#")
			} else {
				fmt.Printf("_")
			}
		}
		fmt.Printf("\n")
	}
}
