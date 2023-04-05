package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var testData, _ = os.ReadFile("testInput.txt")
var testLine = strings.Split(string(testData), "\n")

func TestUnit1(t *testing.T) {
	//rawHex := "60A100"


	var hexnum uint64
	hexnum = 0x38006F45291200
	fmt.Printf("%v\n", hexnum)
	fmt.Printf("%x\n", hexnum)
	fmt.Printf("%056b\n", hexnum)
	num := byteTo16BaseNumber(string(testData))
	fmt.Printf("%08b\n", num)


}

func byteTo16BaseNumber(input string)(num []uint8) {
	num = make([]uint8,0)
	inputLen := len(input)
	for i:=0; i < inputLen ; i+=2{
		a,_ := strconv.ParseInt(input[i:i+2],16,8)
		num = append(num,uint8(a))
	}
	if inputLen % 2 == 1 {
		a,_ := strconv.ParseInt(input[inputLen-1:inputLen]+"0",16,8)
		num = append(num,uint8(a))
	}
	return
}

func TestUnit2(t *testing.T) {

	polymer, pairInsertion := parseInput(&testLine)
	length, mostCnt, leastCnt := polymerization2(polymer, pairInsertion, 10)
	fmt.Printf("length : %d, most common count : %d, least common count : %d\n", length, mostCnt, leastCnt)
	if mostCnt-leastCnt != 1588 {
		t.Errorf("expected 1588 , but it is %d \n", mostCnt-leastCnt)
	}
}

func TestUnit3(t *testing.T) {

	polymer, pairInsertion := parseInput(&testLine)
	length, mostCnt, leastCnt := polymerization2(polymer, pairInsertion, 40)
	fmt.Printf("length : %d, most common count : %d, least common count : %d\n", length, mostCnt, leastCnt)
	if mostCnt-leastCnt != 2188189693529 {
		t.Errorf("expected 2188189693529 , but it is %d \n", mostCnt-leastCnt)
	}
}
