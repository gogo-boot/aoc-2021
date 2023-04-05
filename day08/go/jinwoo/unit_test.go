package main

import (
	"os"
	"strings"
	"testing"
)

var datTest, _ = os.ReadFile("inputTest.txt")
var inputTestLine = strings.Split(string(datTest[:]), "\n")

func TestUnit(t *testing.T) {

	var digits []string
	var count int
	for _, line := range inputTestLine {
		digits = strings.Fields(line)
		if len(digits[11]) >= 2 && len(digits[11]) <= 4 || len(digits[11]) == 7 {
			count++
		}
		if len(digits[12]) >= 2 && len(digits[12]) <= 4 || len(digits[12]) == 7 {
			count++
		}
		if len(digits[13]) >= 2 && len(digits[13]) <= 4 || len(digits[13]) == 7 {
			count++
		}
		if len(digits[14]) >= 2 && len(digits[14]) <= 4 || len(digits[14]) == 7 {
			count++
		}
	}
	println(count)
}

func TestUnit2(t *testing.T) {

	var sum int
	for _, line := range inputTestLine {
		digits := strings.Fields(line)
		numberPattern := getNumberPattern(digits)
		sevenSegment := getSevenSegment(digits, numberPattern)
		sum += interpreter(digits, sevenSegment)
	}
	if 61229 != sum {
		t.Errorf("expected 61229, but return %d", sum)
	}
}
