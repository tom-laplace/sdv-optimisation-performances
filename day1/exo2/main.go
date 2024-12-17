package main

import "fmt"

func main() {
	sliceMaker()
}

func sliceMaker() {
	original := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbers := original[:]

	numbers = append(numbers, 11, 12)

	firstNumbers := numbers[:5]

	firstNumbers[1] = 99

	fmt.Printf("%v", original)
	fmt.Println()
	fmt.Printf("%v", numbers)
	fmt.Println()
	fmt.Printf("%v", firstNumbers)
}
