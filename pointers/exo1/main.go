package main

import "fmt"

func main() {
	a := 100
	b := 200

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

	inversePointerValues(&a, &b)

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}

func inversePointerValues(a, b *int) {
	*a, *b = *b, *a
}
