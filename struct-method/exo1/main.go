package main

import "fmt"

func main() {
	suzuki := Vehicle{
		Brand: "Suzuki",
		Model: "Celerio",
		Year:  2015,
	}

	suzuki.Description()
}

type Vehicle struct {
	Brand string
	Model string
	Year  int
}

func (v *Vehicle) Description() {
	fmt.Print("La voiture est une ", v.Brand, " ", v.Model, " de ", v.Year)
}
