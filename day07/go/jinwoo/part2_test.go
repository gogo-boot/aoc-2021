package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"math"
	"testing"
)

func TestPart2(t *testing.T) {
	genFloatNum(input)
	average, _ := stats.Mean(floatNum)
	average = math.Floor(average)
	fmt.Println(getRequiredSigmaFuel(&floatNum, average))
}

func getRequiredSigmaFuel(numberSlice *[]float64, position float64) (fuel int) {
	for _, num := range *numberSlice {
		n := int(math.Abs(float64(position - num)))
		fuel += (n * (n + 1)) / 2
	}
	return
}
