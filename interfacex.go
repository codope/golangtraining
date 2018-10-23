package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Human struct {
	name string
}

type Animal struct {
	name string
}

type Alien struct {
	name string
}

func (h *Human) Speak() (string) {
	return h.name + " says hello!"
}

func (a *Animal) Speak() (string) {
	return a.name + " says meow!"
}

func (al *Alien) Speak() (string) {
	return al.name + " says dhooooop!"
}

func main() {
	h := &Human{"Sagar"}
	a := &Animal{"Cat"}
	al := &Alien{"Jadoo"}
	speaker := []Speaker{h, a, al}

	for _, obj := range speaker {
		fmt.Println(obj.Speak())
		f(obj)
	}

	fmt.Println("after")
	f(h)
}

func f(i interface{}) {
	switch v := i.(type) {
	case *Human:
		fmt.Println("Hello from ")
	case *Animal:
		fmt.Println("Meow from ", v.name)
	case *Alien:
		fmt.Println("Dhoooop from ", v.name)
	case Speaker:
		fmt.Println("Was speaker ")
	default:
		fmt.Println("Other")
	}
}
