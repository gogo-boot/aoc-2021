package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var data, _ = os.ReadFile("input.txt")
var input = strings.Split(string(data[:]), ",")

func TestPart1(t *testing.T) {
	fish := convertIntSlice(input)
	fishCnt := dayAfterDay(fish, 80)
	fmt.Println(fishCnt)
}

func convertIntSlice(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func dayAfterDay(fish []int, days int) (cnt int) {
	for i := 0; i < days; i++ {
		fishCnt := len(fish)
		for j := 0; j < fishCnt; j++ {
			fish[j]--
			if fish[j] < 0 {
				fish[j] = 6
				fish = append(fish, 8)
			}
		}
	}
	return len(fish)
}
