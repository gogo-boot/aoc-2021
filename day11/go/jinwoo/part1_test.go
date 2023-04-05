package main

import (
	"fmt"
	"os"
	"testing"
)

var data, _ = os.ReadFile("input.txt")
var colCnt = 10

func TestPart1(t *testing.T) {
	octopuses := initOctopus(data)
	step(100, &octopuses)
	fmt.Println(countFlash(&octopuses))
}

func countFlash(octopuses *[]octopus) (sum int) {
	for _, oct := range *octopuses {
		sum += oct.flashCnt
	}
	return
}

func initOctopus(data []byte) (octopuses []octopus) {
	for _, data := range data {
		if data != '\n' {
			octopuses = append(octopuses, octopus{int(data) - 48, 0, false})
		}
	}
	return
}

func step(step int, octopuses *[]octopus) (syncStep int) {
	for i := 0; i < step; i++ {
		var sum int
		for k, _ := range *octopuses {
			(*octopuses)[k].isFlashed = false
		}
		for k, _ := range *octopuses {
			energyUp(k, octopuses)
		}
		for k, _ := range *octopuses {
			sum += (*octopuses)[k].energy
		}
		if sum == 0 {
			return i + 1
		}
	}
	return -1
}

func flash(idx int, octopuses *[]octopus) {
	upleft, up, upright := idx-1-colCnt, idx-colCnt, idx+1-colCnt
	left, right := idx-1, idx+1
	downleft, down, downright := idx-1+colCnt, idx+colCnt, idx+1+colCnt

	if idx%colCnt == 0 { //left side octopus
		// ignore left
		energyUp(up, octopuses)
		energyUp(upright, octopuses)
		energyUp(right, octopuses)
		energyUp(downright, octopuses)
		energyUp(down, octopuses)
	} else if (idx+1)%colCnt == 0 { //right side octopus
		//ignore right
		energyUp(up, octopuses)
		energyUp(upleft, octopuses)
		energyUp(left, octopuses)
		energyUp(downleft, octopuses)
		energyUp(down, octopuses)
	} else {
		energyUp(upleft, octopuses)
		energyUp(up, octopuses)
		energyUp(upright, octopuses)
		energyUp(left, octopuses)
		energyUp(right, octopuses)
		energyUp(downleft, octopuses)
		energyUp(down, octopuses)
		energyUp(downright, octopuses)
	}
}

func energyUp(idx int, octopuses *[]octopus) {
	if idx > -1 && idx < len(*octopuses) && (*octopuses)[idx].isFlashed == false {
		(*octopuses)[idx].energy++
		if (*octopuses)[idx].energy%10 == 0 {
			(*octopuses)[idx].energy = 0
			(*octopuses)[idx].flashCnt++
			(*octopuses)[idx].isFlashed = true
			flash(idx, octopuses)
		}
	}
}

type octopus struct {
	energy, flashCnt int
	isFlashed        bool
}
