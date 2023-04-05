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
	count, preSum := 0, 0

	for i := 0; i < len(line)-1; i++ {
		if i == 0 {
			continue
		}
		sum := sum(line[i-1], line[i], line[i+1])
		if i > 1 && sum > preSum {
			count++
		}
		preSum = sum
	}
	fmt.Printf("Part 2 Answer : %d\n", count)
}

func sum(a, b, c string) int {
	num1, _ := strconv.Atoi(a)
	num2, _ := strconv.Atoi(b)
	num3, _ := strconv.Atoi(c)
	return num1 + num2 + num3
}
