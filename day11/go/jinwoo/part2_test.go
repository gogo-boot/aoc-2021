package main

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {

	octopuses := initOctopus(data)
	syncStep := step(500, &octopuses)
	fmt.Println(syncStep)
}
