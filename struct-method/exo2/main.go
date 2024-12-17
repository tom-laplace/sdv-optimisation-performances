package main

import "fmt"

func main() {
	var mollat Library
	a := Book{
		Title:           "Le loup des steppes",
		Author:          "Herman Hesse",
		PublicationYear: 1989,
	}
	mollat.AddBook(a)

	b := Book{
		Title:           "Demian",
		Author:          "Herman Hesse",
		PublicationYear: 1989,
	}

	mollat.AddBook(b)

	c := Book{
		Title:           "Siddartha",
		Author:          "Herman Hesse",
		PublicationYear: 1989,
	}

	mollat.AddBook(c)

	mollat.ListBooks()
}

type Book struct {
	Title           string
	Author          string
	PublicationYear int
}

type Library struct {
	collection []Book
}

func (l *Library) AddBook(b Book) {
	l.collection = append(l.collection, b)
}

func (l *Library) ListBooks() {
	for i := range l.collection {
		printBook(l.collection[i])
	}
}

func printBook(b Book) {
	fmt.Println(b.Title, b.Author, b.PublicationYear)
}
