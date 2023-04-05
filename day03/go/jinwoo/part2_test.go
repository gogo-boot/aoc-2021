package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	dat, _ := os.ReadFile("input.txt")
	line := strings.Split(string(dat[:]), "\n")

	oxygenRating := getFrequentDigitPerPos(line, 0, false)
	scrubberRating := getFrequentDigitPerPos(line, 0, true)

	oxy, _ := strconv.ParseInt(oxygenRating[0], 2, 16)
	scr, _ := strconv.ParseInt(scrubberRating[0], 2, 16)
	fmt.Printf("Answer : %d\n", oxy*scr)
}

func getFrequentDigitPerPos(input []string, pos int, reverse bool) []string {

	if len(input) == 1 || pos == len(input[0]) {
		return input
	}

	var group0, group1 []string
	for _, line := range input {
		if line[pos:pos+1] == "0" {
			group0 = append(group0, line)
		} else {
			group1 = append(group1, line)
		}
	}

	if len(group0) > len(group1) {
		if reverse {
			return getFrequentDigitPerPos(group1, pos+1, reverse)
		} else {
			return getFrequentDigitPerPos(group0, pos+1, reverse)
		}
	} else {
		if reverse {
			return getFrequentDigitPerPos(group0, pos+1, reverse)
		} else {
			return getFrequentDigitPerPos(group1, pos+1, reverse)
		}
	}
}
