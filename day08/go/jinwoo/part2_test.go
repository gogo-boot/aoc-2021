package main

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestPart2(t *testing.T) {
	var digits []string
	var sum int
	for _, line := range inputLine {
		digits = strings.Fields(line)
		numberPattern := getNumberPattern(digits)
		sevenSegment := getSevenSegment(digits, numberPattern)
		sum += interpreter(digits, sevenSegment)
	}
	println(sum)
}

func interpreter(digits []string, s map[string]string) (sum int) {
	a := s["a"]
	b := s["b"]
	c := s["c"]
	d := s["d"]
	e := s["e"]
	f := s["f"]
	g := s["g"]

	numberPattern := make(map[string]*regexp.Regexp)
	numberPattern["0"], _ = regexp.Compile("[" + a + b + c + e + f + g + "]{6}")
	numberPattern["1"], _ = regexp.Compile("[" + c + f + "]{2}")
	numberPattern["2"], _ = regexp.Compile("[" + a + c + d + e + g + "]{5}")
	numberPattern["3"], _ = regexp.Compile("[" + a + c + d + f + g + "]{5}")
	numberPattern["4"], _ = regexp.Compile("[" + b + c + d + f + "]{4}")
	numberPattern["5"], _ = regexp.Compile("[" + a + b + d + f + g + "]{5}")
	numberPattern["6"], _ = regexp.Compile("[" + a + b + d + e + f + g + "]{6}")
	numberPattern["7"], _ = regexp.Compile("[" + a + c + f + "]{3}")
	numberPattern["8"], _ = regexp.Compile("[" + a + b + c + d + e + f + g + "]{7}")
	numberPattern["9"], _ = regexp.Compile("[" + a + b + c + d + f + g + "]{6}")
	var sumstring string

	for i := 11; i < 15; i++ {
		for k, pattern := range numberPattern {
			if len(pattern.ReplaceAllString(digits[i], "")) == 0 {
				sumstring += k
			}
		}
	}
	sum, _ = strconv.Atoi(sumstring)
	return
}

func getNumberPattern(digit []string) (numberPattern map[int]*regexp.Regexp) {

	numberPattern = make(map[int]*regexp.Regexp)
	for i := 0; i < 10; i++ {
		if len(digit[i]) == 2 {
			numberPattern[1], _ = regexp.Compile("[" + digit[i] + "]")
		} else if len(digit[i]) == 4 {
			numberPattern[4], _ = regexp.Compile("[" + digit[i] + "]")
		} else if len(digit[i]) == 3 {
			numberPattern[7], _ = regexp.Compile("[" + digit[i] + "]")
		} else if len(digit[i]) == 7 {
			numberPattern[8], _ = regexp.Compile("[" + digit[i] + "]")
		}
	}
	return numberPattern
}

func getSevenSegment(digits []string, numberPattern map[int]*regexp.Regexp) (sevenSegment map[string]string) {

	digit := strings.Join(digits[0:10], "")
	sevenSegment = make(map[string]string)
	counter := make(map[string]int)
	counter["a"] = strings.Count(digit, "a")
	counter["b"] = strings.Count(digit, "b")
	counter["c"] = strings.Count(digit, "c")
	counter["d"] = strings.Count(digit, "d")
	counter["e"] = strings.Count(digit, "e")
	counter["f"] = strings.Count(digit, "f")
	counter["g"] = strings.Count(digit, "g")

	for k, v := range counter {
		if v == 6 {
			sevenSegment["b"] = k
		} else if v == 4 {
			sevenSegment["e"] = k
		} else if v == 9 {
			sevenSegment["f"] = k
		} else if v == 7 {
			// d or g
			// patter4 - b,c,f rest d
			if numberPattern[4].MatchString(k) {
				sevenSegment["d"] = k
			} else {
				// rest g
				sevenSegment["g"] = k
			}
		} else if v == 8 {
			// a or c
			// pattern1 match c
			if numberPattern[1].MatchString(k) {
				// rest == c
				sevenSegment["c"] = k
			} else {
				sevenSegment["a"] = k
			}
		}
	}
	return
}
