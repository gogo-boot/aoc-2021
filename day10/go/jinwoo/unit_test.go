package main

import (
	"os"
	"strings"
	"testing"
)

var testData, _ = os.ReadFile("testInput.txt")
var testInput = strings.Split(string(testData[:]), "\n")

func TestUnit1(t *testing.T) {

	var sum int
	for _, line := range testInput {
		idx, expected, _ := checkValidate(line)
		if idx > 0 {
			sum += expected
		}
	}
	if sum != 26397 {
		t.Errorf("Expected 26397, but it was %d", sum)
	}
}
