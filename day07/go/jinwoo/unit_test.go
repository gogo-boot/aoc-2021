package main

import (
	"testing"
)

func TestUnit(t *testing.T) {
	input := []float64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	fuel := getRequiredFuel(&input, 2)
	if fuel != 37 {
		t.Errorf("want 37 fuel and position 2, but got %d fuel", fuel)
	}
	fuel = getRequiredFuel(&input, 3)
	if fuel != 39 {
		t.Errorf("want 39 fuel and position 3, but got %d fuel", fuel)
	}

	fuel = getRequiredFuel(&input, 10)
	if fuel != 71 {
		t.Errorf("want 71 fuel and position 10, but got %d fuel", fuel)
	}
}

func TestUnit2(t *testing.T) {
	input := []float64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	fuel := getRequiredSigmaFuel(&input, 2)
	if fuel != 206 {
		t.Errorf("want 206 fuel and position 2, but got %d fuel", fuel)
	}
	fuel = getRequiredSigmaFuel(&input, 5)
	if fuel != 168 {
		t.Errorf("want 168 fuel and position 5, but got %d fuel", fuel)
	}
}
