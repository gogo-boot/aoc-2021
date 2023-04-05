package main

import (
	"os"
	"testing"
)

var testData, _ = os.ReadFile("testInput.txt")

func TestUnit1(t *testing.T) {

	var sum int
	octopuses := initOctopus(testData)
	step(10, &octopuses)
	sum = countFlash(&octopuses)

	if sum != 204 {
		t.Errorf("Expected 204, but it was %d", sum)
	}
	step(90, &octopuses)
	sum = countFlash(&octopuses)
	if sum != 1656 {
		t.Errorf("Expected 1656, but it was %d", sum)
	}
}

func TestUnit2(t *testing.T) {

	var sum int
	octopuses := initOctopus(testData)
	syncStep := step(200, &octopuses)

	if syncStep != 195 {
		t.Errorf("Expected 195, but it was %d", sum)
	}
}
