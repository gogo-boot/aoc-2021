package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"testing"
)

func TestPart2(t *testing.T) {

	var sumSlice []float64
	for _, line := range input {
		var sum int
		idx, _, stack := checkValidate(line)
		if idx < 0 {
			for len(stack) > 0 {
				n := len(stack) - 1
				for _, bracket := range brackets {
					if stack[n] == bracket.open {
						sum *= 5
						sum += bracket.p2
					}
				}
				stack = stack[:n] // Pop
			}
			sumSlice = append(sumSlice, float64(sum))
		}
	}
	fmt.Println(stats.Median(sumSlice))
}
