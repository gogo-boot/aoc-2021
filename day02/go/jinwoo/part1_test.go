package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var depth, forward int64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		operation := strings.Split(line, " ")
		factor, _ := strconv.ParseInt(operation[1], 10, 16)
		switch operation[0] {
		case "up":
			depth -= factor
		case "down":
			depth += factor
		case "forward":
			forward += factor
		}
	}
	fmt.Printf("Part 1 Answer : %d\n", depth*forward)
}
