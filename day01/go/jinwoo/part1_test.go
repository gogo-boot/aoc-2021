package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	preNum := 9999999999
	count := 0

	for scanner.Scan() {
		curNum, _ := strconv.Atoi(scanner.Text())
		if curNum > preNum {
			count++
		}
		preNum = curNum
	}
	fmt.Printf("Part 1 Answer : %d\n", count)
}
