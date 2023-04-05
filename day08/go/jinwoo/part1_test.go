package main

import (
	"os"
	"strings"
	"testing"
)

var dat, _ = os.ReadFile("input.txt")
var inputLine = strings.Split(string(dat[:]), "\n")

func TestPart1(t *testing.T) {
	var digits []string
	var count int
	for _, line := range inputLine {
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
