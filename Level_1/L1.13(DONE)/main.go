package main

import "fmt"

func main() {
	var (
		elem1 = 1
		elem2 = 2
	)
	fmt.Printf("before: elem1: %d - elem2: %d\n", elem1, elem2)
	//значение elem2 присваивается переменной elem1, а значение elem1 — переменной elem2.
	elem1, elem2 = elem2, elem1

	fmt.Printf("after: elem1: %d - elem2: %d", elem1, elem2)

}
