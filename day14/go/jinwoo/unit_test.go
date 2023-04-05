package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var testData, _ = os.ReadFile("testInput.txt")
var testLine = strings.Split(string(testData), "\n")

func TestUnit1(t *testing.T) {

	polymer, pairInsertion := parseInput(&testLine)
	length, mostCnt, leastCnt := polymerization(polymer, pairInsertion, 10)
	fmt.Printf("length : %d, most common count : %d, least common count : %d\n", length, mostCnt, leastCnt)
	if mostCnt-leastCnt != 1588 {
		t.Errorf("expected 1588 , but it is %d \n", mostCnt-leastCnt)
	}
}

func TestUnit2(t *testing.T) {

	polymer, pairInsertion := parseInput(&testLine)
	length, mostCnt, leastCnt := polymerization2(polymer, pairInsertion, 10)
	fmt.Printf("length : %d, most common count : %d, least common count : %d\n", length, mostCnt, leastCnt)
	if mostCnt-leastCnt != 1588 {
		t.Errorf("expected 1588 , but it is %d \n", mostCnt-leastCnt)
	}
}

func TestUnit3(t *testing.T) {

	polymer, pairInsertion := parseInput(&testLine)
	length, mostCnt, leastCnt := polymerization2(polymer, pairInsertion, 40)
	fmt.Printf("length : %d, most common count : %d, least common count : %d\n", length, mostCnt, leastCnt)
	if mostCnt-leastCnt != 2188189693529 {
		t.Errorf("expected 2188189693529 , but it is %d \n", mostCnt-leastCnt)
	}
}
