package main

import "fmt"

func main() {
	paul := Person{"Paul", 17}
	fmt.Printf("%v", paul)
	paul.ModifyPerson()
	fmt.Printf("%v", paul)

	jean := Person{"Jean", 16}
	fmt.Printf("%v", jean)
	jean.ModifyPerson()
	fmt.Printf("%v", jean)
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ModifyPerson() {
	if p.Name == "Paul" {
		p.Name = "Lisan al Gaib"
	}
	p.Age++
}
