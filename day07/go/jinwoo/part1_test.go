package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

var dat, _ = os.ReadFile("input.txt")
var input = strings.Split(string(dat[:]), ",")
var floatNum []float64

func TestPart1(t *testing.T) {
	genFloatNum(input)
	median, _ := stats.Median(floatNum)
	fmt.Println(getRequiredFuel(&floatNum, median))
}

func getRequiredFuel(numberSlice *[]float64, position float64) (fuel int) {
	for _, num := range *numberSlice {
		fuel += int(math.Abs(float64(position - num)))
	}
	return
}

func genFloatNum(input []string) {
	for _, stringNumber := range input {
		num, _ := strconv.Atoi(stringNumber)
		floatNum = append(floatNum, float64(num))
	}
}
