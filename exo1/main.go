package main

import "fmt"

func main() {
	printDays()
}

func printDays() {
	days := [7]string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
}
