package main

import (
	"os"
	"strings"
	"testing"
)

func TestUnit(t *testing.T) {
	datTest, _ := os.ReadFile("inputTest.txt")
	inputTestLine := strings.Split(string(datTest[:]), "\n")
	sum := getSumOfLowRiskPoint(&inputTestLine)
	if sum != 15 {
		t.Errorf("exptected 15. but it is %d", sum)
	}
}

func TestUnit2(t *testing.T) {
	datTest, _ := os.ReadFile("inputTest2.txt")
	inputTestLine := strings.Split(string(datTest[:]), "\n")
	sum := getSumOfLowRiskPoint(&inputTestLine)
	if sum != 35 {
		t.Errorf("exptected 35. but it is %d", sum)
	}
}

func TestUnit3(t *testing.T) {
	datTest, _ := os.ReadFile("inputTest3.txt")
	inputTestLine := strings.Split(string(datTest[:]), "\n")
	sum := getSumOfLowRiskPoint(&inputTestLine)
	if sum != 4 {
		t.Errorf("exptected 4. but it is %d", sum)
	}
}

func TestUnit4(t *testing.T) {
	datTest, _ := os.ReadFile("inputTest.txt")
	inputTestLine := strings.Split(string(datTest[:]), "\n")
	multiplyLagestTop3BasinSize(&inputTestLine)
}

func TestUnit5(t *testing.T) {
	datTest, _ := os.ReadFile("inputTest2.txt")
	inputTestLine := strings.Split(string(datTest[:]), "\n")
	multiplyLagestTop3BasinSize(&inputTestLine)
}

func TestUnit6(t *testing.T) {
	datTest, _ := os.ReadFile("inputTest3.txt")
	inputTestLine := strings.Split(string(datTest[:]), "\n")
	multiplyLagestTop3BasinSize(&inputTestLine)
}
