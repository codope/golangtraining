package main

import "fmt"

func main() {
	var a = [5]int{2,3,4,5,6}
	s := a[:3]
	for i := 0; i < 5; i++ {
		s = append(s, 1)
		fmt.Printf("Slice: %v   Len: %d   Cap: %d", s, len(s), cap(s))
		fmt.Println()
	}
}
