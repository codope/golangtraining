package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	book := Book{Title:"Hello", Author:"World"}
	bookB, _ := json.Marshal(book)
	ioutil.WriteFile("book.json", bookB, 0755)

	var b Book
	bookR, _ := ioutil.ReadFile("book.json")
	err := json.Unmarshal(bookR, &b)
	if err == nil {
		fmt.Println(b.Author, b.Title)
	}
	books := []Book{
		{"Learning Go", "Sagar Sumit"},
		{"Fun with Go", "Sampathu M"},
		{"Hello Gopher", "Prashanth Surya"},
	}
	booksB, _ := json.Marshal(books)
	ioutil.WriteFile("books.json", booksB, 0755)

	var bs []Book
	booksR, _ := ioutil.ReadFile("books.json")
	if err := json.Unmarshal(booksR, &bs); err == nil {
		for _, buk := range bs {
			fmt.Println(buk.Title, buk.Author)
		}
	}
}
