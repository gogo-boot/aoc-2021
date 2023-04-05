package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var depth, forward, aim int64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		operation := strings.Split(line, " ")
		factor, _ := strconv.ParseInt(operation[1], 10, 16)
		switch operation[0] {
		case "up":
			aim -= factor
		case "down":
			aim += factor
		case "forward":
			forward += factor
			depth += aim * factor
		}
	}
	fmt.Printf("Part 2 Answer : %d\n", depth*forward)
}
