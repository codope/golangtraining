package main

import (
	"errors"
	"fmt"
)

func main() {
	m := make(map[int]map[string]string)
	fmt.Println(m)

	// country to currency
	cc := map[string]string{"US":"USD", "UK":"GBP", "SG":"SGD", "TH":"BHT"}
	val, ok := cc["IN"]
	if ok {
		fmt.Println(val, ok)
	} else {
		fmt.Println("Key absent!")
	}
	for k, v := range cc {
		fmt.Printf("%s currency is %s", k, v)
		fmt.Println()
	}

	// vector dot product
	var v1 = [][]int{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	var v2 = [][]int{
		{1,2,3,4,5},
	}

	dotprod, err := dot(v1, v2)
	if err == nil {
		fmt.Println("Dot Product:", dotprod)
	}
}

func dot(x, y [][]int) ([][]int, error) {
	if len(x[0]) != len(y) {
		return nil, errors.New("Can't do matrix multiplication.")
	}

	out := make([][]int, len(x))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(y); j += 1 {
			if len(out[i]) < 1 {
				out[i] = make([]int, len(y))
			}
			out[i][j] += x[i][j] * y[j][i]
		}
	}
	return out, nil
}
