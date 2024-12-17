package main

import (
	"fmt"
)

var ROMAN_NUM = map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}

func main() {
	value := romainToInteger("XXI")
	fmt.Println(value)
}

func romainToInteger(s string) int {
	sum := 0
	i := 0

	for i < len(s) {
		if i+1 < len(s) {
			twoChar := s[i : i+2]
			if val, exists := ROMAN_NUM[twoChar]; exists {
				sum += val
				i += 2
				continue
			}
		}

		oneChar := string(s[i])
		if val, exists := ROMAN_NUM[oneChar]; exists {
			sum += val
		}
		i++
	}

	return sum
}
