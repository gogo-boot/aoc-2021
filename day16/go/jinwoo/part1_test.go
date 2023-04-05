package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

var data, _ = os.ReadFile("input.txt")
var line = strings.Split(string(data), "\n")

func TestPart1(t *testing.T) {

	polymer, pairInsertion := parseInput(&line)
	length, mostCnt, leastCnt := polymerization(polymer, pairInsertion, 10)
	fmt.Printf("length : %d, most common count : %d, least common count : %d\n", length, mostCnt, leastCnt)
	fmt.Printf("answer : %d\n", mostCnt-leastCnt)
	if mostCnt-leastCnt != 3230 {
		t.Errorf("expected 1588 , but it is %d \n", mostCnt-leastCnt)
	}
}

func polymerization(polymer []rune, insertion map[string]rune, step int) (length, mostCnt, leastCnt int) {

	i := 0
	for i < step {
		todo := make(map[int]rune)
		for j := 0; j < len(polymer)-1; j++ {
			key := fmt.Sprintf("%c", polymer[j]) + fmt.Sprintf("%c", polymer[j+1])
			if value, ok := insertion[key]; ok {
				todo[j+1] = value
			}
		}
		for j := len(polymer) - 1; j > -1; j-- {
			if value, ok := todo[j]; ok {
				polymer = insert(polymer, j, value)
			}
		}
		//fmt.Printf("%s \n", string(polymer))
		i++
	}
	length = len(string(polymer))
	mostCnt, leastCnt = getMostLeastCnt(polymer)
	return
}

func getMostLeastCnt(r []rune) (mostCnt, leastCnt int) {

	counter := make(map[rune]int)
	for _, rune := range r {
		counter[rune]++
	}
	values := make([]int, 0, len(counter))
	for _, v := range counter {
		values = append(values, v)
	}
	sort.Ints(values)

	return values[len(values)-1], values[0]
}

func insert(a []rune, index int, value rune) []rune {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func parseInput(line *[]string) (polymer []rune, pairInsertion map[string]rune) {

	pairInsertion = make(map[string]rune)
	for i, line := range *line {
		if i == 0 {
			polymer = make([]rune, len(line))
			for j, char := range line {
				polymer[j] = char
			}
		}
		if i > 1 {
			line := strings.Fields(line)
			pairInsertion[line[0]] = ([]rune(line[2]))[0]
		}
	}
	return
}
