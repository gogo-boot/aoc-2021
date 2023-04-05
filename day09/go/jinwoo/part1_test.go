package main

import (
	"os"
	"strings"
	"testing"
)

var dat, _ = os.ReadFile("input.txt")
var inputLine = strings.Split(string(dat[:]), "\n")

func TestPart1(t *testing.T) {
	sum := getSumOfLowRiskPoint(&inputLine)
	println(sum)
}

func getSumOfLowRiskPoint(inputLine *[]string) int32 {
	var digits []byte
	var byteCountPerLine int
	for _, line := range *inputLine {
		digits = append(digits, line...)
		if byteCountPerLine == 0 {
			byteCountPerLine = len([]byte(line))
		}
	}

	var sum int32
	for i, digit := range digits {
		left, right, up, down := getValueOfIndex(i, byteCountPerLine, digits)
		if digit < left && digit < right && digit < up && digit < down {
			sum += int32(int(digit) - '0' + 1)
		}
	}
	return sum
}

func getValueOfIndex(idx int, byteCountPerLine int, digits []byte) (left, right, up, down byte) {

	idxLeft, idxRight, idxUp, idxDown := idx-1, idx+1, idx-byteCountPerLine, idx+byteCountPerLine

	if idxLeft < 0 || idxLeft%byteCountPerLine == byteCountPerLine-1 {
		left = '9'
	} else {
		left = digits[idxLeft]
	}
	if idxRight > len(digits)-1 || idxRight%byteCountPerLine == 0 {
		right = '9'
	} else {
		right = digits[idxRight]
	}
	if idxUp < 0 {
		up = '9'
	} else {
		up = digits[idxUp]
	}
	if idxDown > len(digits)-1 {
		down = '9'
	} else {
		down = digits[idxDown]
	}
	return
}
