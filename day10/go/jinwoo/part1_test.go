package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var data, _ = os.ReadFile("input.txt")
var input = strings.Split(string(data[:]), "\n")

type bracket struct {
	open, close string
	p1, p2      int
}

var roundBracket = bracket{"(", ")", 3, 1}
var squareBracket = bracket{"[", "]", 57, 2}
var curlyBracket = bracket{"{", "}", 1197, 3}
var angleBracket = bracket{"<", ">", 25137, 4}
var brackets = []bracket{roundBracket, squareBracket, curlyBracket, angleBracket}

func TestPart1(t *testing.T) {
	var sum int
	for _, line := range testInput {
		idx, expected, _ := checkValidate(line)
		if idx > 0 {
			sum += expected
		}
	}
	fmt.Println(sum)
}

func checkValidate(line string) (idx int, expect int, stack []string) {

	for _, char := range line {
		strChar := string(char)
		if closeParenthese(strChar) && len(stack) > 0 {
			n := len(stack) - 1 // Top element
			last := stack[n]
			stack = stack[:n] // Pop
			for _, bracket := range brackets {
				if strChar == bracket.close && last != bracket.open {
					return n, bracket.p1, stack
				}
			}
		} else {
			stack = append(stack, strChar)
		}
	}
	return -1, 0, stack
}

func closeParenthese(char string) bool {
	return char == "}" || char == ")" || char == "]" || char == ">"
}
