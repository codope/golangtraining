package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	b := 5
	sum, prd, subt, div, err := sumProdSubt(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sum: %d  Product: %d  Diff: %d  Divis: %f", sum, prd, subt, div)
	}
}

func sumProdSubt(a, b int) (int, int, int, float64, error) {
	if b == 0 {
		return a + b, a * b, a - b, -1, errors.New("division by 0!")
	}
	return a + b, a * b, a - b, float64(a) / float64(b), nil
}
