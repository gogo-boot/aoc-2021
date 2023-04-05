package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestPart2(t *testing.T) {
	multiplyLagestTop3BasinSize(&inputLine)
}

func multiplyLagestTop3BasinSize(inputLine *[]string) {
	var digits []byte
	var byteCountPerLine int
	for _, line := range *inputLine {
		digits = append(digits, line...)
		if byteCountPerLine == 0 {
			byteCountPerLine = len([]byte(line))
		}
	}

	basinMap := make([]int, len(digits))
	for i, digit := range digits {
		left, right, up, down := getValueOfIndex(i, byteCountPerLine, digits)
		if digit < left && digit < right && digit < up && digit < down {
			var checkedIdx = make(map[int]bool)
			var count = 0
			basinMap[i] = getBasinSize(i, byteCountPerLine, &digits, &count, &checkedIdx)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinMap)))
	fmt.Println(basinMap[0] * basinMap[1] * basinMap[2])
}

func getBasinSize(lowpointIdx int, byteCountPerLine int, digits *[]byte, basinSize *int, checkedIdx *map[int]bool) (newSize int) {

	(*checkedIdx)[lowpointIdx] = true
	idxLeft, idxRight, idxUp, idxDown := lowpointIdx-1, lowpointIdx+1, lowpointIdx-byteCountPerLine, lowpointIdx+byteCountPerLine
	left, right, up, down := getValueOfIndex(lowpointIdx, byteCountPerLine, *digits)

	if left != '9' && !(*checkedIdx)[idxLeft] {
		*basinSize = *basinSize + 1
		getBasinSize(idxLeft, byteCountPerLine, digits, basinSize, checkedIdx)
	}
	if right != '9' && !(*checkedIdx)[idxRight] {
		*basinSize = *basinSize + 1
		getBasinSize(idxRight, byteCountPerLine, digits, basinSize, checkedIdx)
	}
	if up != '9' && !(*checkedIdx)[idxUp] {
		*basinSize = *basinSize + 1
		getBasinSize(idxUp, byteCountPerLine, digits, basinSize, checkedIdx)
	}
	if down != '9' && !(*checkedIdx)[idxDown] {
		*basinSize = *basinSize + 1
		getBasinSize(idxDown, byteCountPerLine, digits, basinSize, checkedIdx)
	}
	return *basinSize + 1
}
