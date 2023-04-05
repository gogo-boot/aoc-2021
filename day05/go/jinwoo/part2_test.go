package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	filteredLineSegment := getLines(inputLine, true)
	thermalVent := drawline(&filteredLineSegment)
	fmt.Printf("Part2 Answer : %d \n", countLineOver(&thermalVent, 1))
}

func getLines(inputLine []string, hasDiagonal bool) (filteredLineSegment []lineSegment) {
	for _, s := range inputLine {
		xyList := strings.Replace(s, " -> ", ",", -1) // 111,111 -> 222,222 to 111,111,222,222
		xy := strings.Split(xyList, ",")
		x1, _ := strconv.Atoi(xy[0])
		y1, _ := strconv.Atoi(xy[1])
		x2, _ := strconv.Atoi(xy[2])
		y2, _ := strconv.Atoi(xy[3])
		if x1 == x2 || y1 == y2 || hasDiagonal && math.Abs(float64(x2-x1)) == math.Abs(float64(y2-y1)) {
			filteredLineSegment = append(filteredLineSegment, lineSegment{x1, y1, x2, y2})
		}
	}
	return
}

func drawDigonalLine(segment *lineSegment, thermalVent *[1000][1000]int) {
	j := segment.y1
	for i := segment.x1; i <= segment.x2 && j <= segment.y2; i++ { //increase x value and increase y value
		(*thermalVent)[i][j]++
		j++
	}
	j = segment.y1
	for i := segment.x1; i <= segment.x2 && j >= segment.y2; i++ { //increase x value and decrease y value
		(*thermalVent)[i][j]++
		j--
	}
	j = segment.y1
	for i := segment.x1; i >= segment.x2 && j <= segment.y2; i-- { //decrease x value and increase y value
		(*thermalVent)[i][j]++
		j++
	}
	j = segment.y1
	for i := segment.x1; i >= segment.x2 && j >= segment.y2; i-- { //decrease x value and decrease y value
		(*thermalVent)[i][j]++
		j--
	}
}
