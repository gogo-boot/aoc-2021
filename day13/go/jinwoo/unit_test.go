package main

import (
	"os"
	"strings"
	"testing"
)

var testData, _ = os.ReadFile("testInput.txt")
var testLine = strings.Split(string(testData[:]), "\n")

func TestUnit1(t *testing.T) {

	input, _ := parseInput(&testLine)
	printStatus(input)

	foldStr := "fold along y=7"
	fold(input, foldStr)
	cnt := countVisible(input)
	if cnt != 17 {
		t.Errorf("expected 17 , but it is %d \n", cnt)
		printStatus(input)
	}

	fold(input, "fold along x=5")
	cnt = countVisible(input)
	if cnt != 16 {
		t.Errorf("expected 16 , but it is %d \n", cnt)
		printStatus(input)
	}
}
