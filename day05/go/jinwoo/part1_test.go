package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var dat, _ = os.ReadFile("input.txt")
var inputLine = strings.Split(string(dat[:]), "\n")

func TestPart1(t *testing.T) {
	filteredLineSegment := getLines(inputLine, false)
	thermalVent := drawline(&filteredLineSegment)
	fmt.Printf("Part1 Answer : %d \n", countLineOver(&thermalVent, 1))
}

func countLineOver(thermalVent *[1000][1000]int, limit int) (counter int) {
	for i := 0; i < len(*thermalVent); i++ {
		for j := 0; j < len((*thermalVent)[i]); j++ {
			if (*thermalVent)[i][j] > limit {
				counter++
			}
		}
	}
	return
}

func drawline(lineSegment *[]lineSegment) (thermalVent [1000][1000]int) {
	for _, segment := range *lineSegment {
		if segment.x1 == segment.x2 || segment.y1 == segment.y2 {
			drawHoriVertiLine(&segment, &thermalVent)
		} else {
			drawDigonalLine(&segment, &thermalVent)
		}
	}
	return
}

func drawHoriVertiLine(segment *lineSegment, thermalVent *[1000][1000]int) {

	if (*segment).x1 <= (*segment).x2 && (*segment).y1 == (*segment).y2 { //increase x value
		for i := (*segment).x1; i <= (*segment).x2; i++ {
			(*thermalVent)[i][(*segment).y1]++
		}
	} else if (*segment).x2 <= (*segment).x1 && (*segment).y1 == (*segment).y2 { //decrease x value
		for i := (*segment).x2; i <= (*segment).x1; i++ {
			(*thermalVent)[i][(*segment).y1]++
		}
	} else if (*segment).y1 <= (*segment).y2 && (*segment).x1 == (*segment).x2 { //increase y value
		for i := (*segment).y1; i <= (*segment).y2; i++ {
			(*thermalVent)[(*segment).x1][i]++
		}
	} else if (*segment).y2 <= (*segment).y1 && (*segment).x1 == (*segment).x2 { //decrease y value
		for i := (*segment).y2; i <= (*segment).y1; i++ {
			(*thermalVent)[(*segment).x1][i]++
		}
	}
}

type lineSegment struct {
	x1, y1, x2, y2 int
}
