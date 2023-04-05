package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"testing"
)

func TestPart2(t *testing.T) {
	fish := convertIntSlice(input)
	fishCnt := dayAfterDay2(fish, 256)
	fmt.Println(fishCnt)
}

func dayAfterDay2(fish []int, days int) (cnt float64) {
	bucket := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, fish := range fish {
		bucket[fish]++
	}
	var bucket0 float64
	for i := 0; i < days; i++ {
		bucket0 = bucket[0]
		bucket[0] = bucket[1]
		bucket[1] = bucket[2]
		bucket[2] = bucket[3]
		bucket[3] = bucket[4]
		bucket[4] = bucket[5]
		bucket[5] = bucket[6]
		bucket[6] = bucket[7]
		bucket[7] = bucket[8]
		bucket[8] = bucket0
		bucket[6] += bucket0
	}
	sum, _ := stats.Sum(bucket)
	return sum
}
