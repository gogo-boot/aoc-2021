package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestPart2(t *testing.T) {

	polymer, pairInsertion := parseInput(&line)
	length, mostCnt, leastCnt := polymerization2(polymer, pairInsertion, 40)
	fmt.Printf("length : %d, most common count : %d, least common count : %d\n", length, mostCnt, leastCnt)
	fmt.Printf("answer : %d\n", mostCnt-leastCnt)
	if mostCnt-leastCnt != 3230 {
		t.Errorf("expected 1588 , but it is %d \n", mostCnt-leastCnt)
	}
}

func polymerization2(polymer []rune, insertion map[string]rune, step int) (length, mostCnt, leastCnt int) {

	var pair map[string]int
	pair = make(map[string]int)

	for j := 0; j < len(polymer)-1; j++ {
		firstChar := fmt.Sprintf("%c", polymer[j])
		lastChar := fmt.Sprintf("%c", polymer[j+1])
		key := firstChar + lastChar
		pair[key]++
	}
	//fmt.Printf("%v \n", pair)
	var newpair map[string]int

	i := 0
	newpair = nextStep(pair, insertion)
	for i < step-1 {
		newpair = nextStep(newpair, insertion)
		i++
	}
	mostCnt, leastCnt = getMostLeastCnt2(newpair)
	return
}

func getMostLeastCnt2(newpair map[string]int) (int, int) {

	firstcount := true
	countMap := make(map[string]int)
	for k, v := range newpair {
		if k[:1] == k[1:2] {
			countMap[k[1:2]] += v
		} else {
			if firstcount {
				countMap[k[0:1]]++
				firstcount = false
			}
			countMap[k[1:2]] += v
		}
	}

	values := make([]int, 0)
	for _, v := range countMap {
		values = append(values, v)
	}
	sort.Ints(values)

	return values[len(values)-1], values[0]
}

func nextStep(pair map[string]int, insertion map[string]rune) (newpair map[string]int) {

	newpair = make(map[string]int)

	for k, v := range pair {
		if value, ok := insertion[k]; ok {
			strValue := fmt.Sprintf("%c", value)
			newpair[k[0:1]+strValue] += v
			newpair[strValue+k[1:2]] += v
		} else {
			newpair[k]++
		}
	}
	//fmt.Printf("%v \n", newpair)
	return
}
