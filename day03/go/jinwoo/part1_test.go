package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	dat, _ := os.ReadFile("input.txt")
	line := strings.Split(string(dat[:]), "\n")

	gamma, _ := strconv.ParseInt(frequentDigit(line, "1"), 2, 32)
	epsilon, _ := strconv.ParseInt(frequentDigit(line, "0"), 2, 32)

	fmt.Printf("Answer : %d\n", epsilon*gamma)
}

func frequentDigit(a []string, digit string) string {

	counter := make(map[int]int)
	lineCount := len(a)
	bitCount := len(a[0])

	for i := 0; i < lineCount; i++ {
		for j := 0; j < bitCount; j++ {
			if a[i][j:j+1] == digit {
				counter[j]++
			}
		}
	}

	var sb strings.Builder
	for j := 0; j < bitCount; j++ {
		if counter[j] > lineCount/2 {
			sb.WriteString("0")
		} else {
			sb.WriteString("1")
		}
	}
	return sb.String()
}
