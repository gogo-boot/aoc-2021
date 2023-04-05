package main

import (
	"testing"
)

func TestBasic(t *testing.T) {
	input := []string{"0,0 -> 9,9", "0,0 -> 9,9"}
	filteredLineSegment := getLines(input, true)
	thermalVent := drawline(&filteredLineSegment)
	value := countLineOver(&thermalVent, 1)
	if value != 10 {
		t.Errorf("crossline count was incorrect, got: %d, want: %d.", value, 10)
	}
}

func TestBasic2(t *testing.T) {
	input := []string{"0,0 -> 9,9", "0,0 -> 9,9", "0,0 -> 9,9",
		"0,0 -> 9,9", "0,0 -> 9,9", "0,0 -> 9,9",
		"0,0 -> 9,9", "0,0 -> 9,9", "0,0 -> 9,9",
		"0,0 -> 8,8"}
	filteredLineSegment := getLines(input, true)
	thermalVent := drawline(&filteredLineSegment)
	value := countLineOver(&thermalVent, 1)
	if value != 10 {
		t.Errorf("crossline count was incorrect, got: %d, want: %d.", value, 10)
	}

}

func TestBasic3(t *testing.T) {
	input := []string{"0,0 -> 9,9", "0,0 -> 9,9", "0,0 -> 9,9"}
	filteredLineSegment := getLines(input, true)
	thermalVent := drawline(&filteredLineSegment)
	value := countLineOver(&thermalVent, 1)
	if value != 10 {
		t.Errorf("crossline count was incorrect, got: %d, want: %d.", value, 10)
	}
}

func TestBasic4(t *testing.T) {
	input := []string{"0,0 -> 9,9", "1,0 -> 9,8"}
	filteredLineSegment := getLines(input, true)
	thermalVent := drawline(&filteredLineSegment)
	value := countLineOver(&thermalVent, 1)
	if value != 0 {
		t.Errorf("crossline count was incorrect, got: %d, want: %d.", value, 0)
	}
}

func TestBasic5(t *testing.T) {
	input := []string{"0,0 -> 2,2", "2,0 -> 0,2"}
	filteredLineSegment := getLines(input, true)
	thermalVent := drawline(&filteredLineSegment)
	value := countLineOver(&thermalVent, 1)
	if value != 1 {
		t.Errorf("crossline count was incorrect, got: %d, want: %d.", value, 1)
	}
}

func TestBasic6(t *testing.T) {
	input := []string{"0,0 -> 2,2", "0,2 -> 2,0"}
	filteredLineSegment := getLines(input, true)
	thermalVent := drawline(&filteredLineSegment)
	value := countLineOver(&thermalVent, 1)
	if value != 1 {
		t.Errorf("crossline count was incorrect, got: %d, want: %d.", value, 1)
	}
}
