package main

import "fmt"

func main() {
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	printDays(days)

	days[2] = "Woden's Day"

	printDays(days)
}

func printDays(days [7]string) {
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
}
