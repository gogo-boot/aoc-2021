package main

import (
	"testing"
)

func TestUnit1(t *testing.T) {
	fishCnt := dayAfterDay([]int{3, 4, 3, 1, 2}, 18)
	if fishCnt != 26 {
		t.Errorf("Expected 26 fishes, but it was %d", fishCnt)
	}
	fishCnt = dayAfterDay([]int{3, 4, 3, 1, 2}, 80)
	if fishCnt != 5934 {
		t.Errorf("Expected 5934 fishes, but it was %d", fishCnt)
	}
}

func TestUnit2(t *testing.T) {
	fishCnt := dayAfterDay2([]int{3, 4, 3, 1, 2}, 18)
	if fishCnt != 26 {
		t.Errorf("Expected 26 fishes, but it was %f", fishCnt)
	}
	fishCnt = dayAfterDay2([]int{3, 4, 3, 1, 2}, 80)
	if fishCnt != 5934 {
		t.Errorf("Expected 5934 fishes, but it was %f", fishCnt)
	}
}
