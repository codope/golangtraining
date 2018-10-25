package main

import (
	"os"
	"text/template"
)

type Person struct {
	//exported field since it begins with a capital letter
	Name
	Age int
}

type Name struct {
	Firstname string
	Lastname string
}

func main() {
	//create a new template with some name
	t := template.New("hello template")

	//parse some content and generate a template,
	//which is an internal representation
	t, _ = t.Parse("hello {{.Name.Lastname}} of age {{.Age}} years!")

	//define an instance with required field
	n := Name{Firstname: "Sagar", Lastname: "Sumit"}
	p := Person{Name: n, Age: 29}

	//merge template ‘t’ with content of ‘p’
	t.Execute(os.Stdout, p)
}
