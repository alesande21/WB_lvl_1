package main

import "fmt"

func main() {
	arr1 := []int{4, 2, 3, 6, 1}
	arr2 := []int{10, 1, 8, 2, 3, 4}
	var res []int

	// проходим по элементам первого слайса
	for _, e1 := range arr1 {
		// проходим по элементам второго слайса
		for _, e2 := range arr2 {
			// если элементы совпадают добавляем элемент в результирующий слайс
			if e1 == e2 {
				res = append(res, e1)
			}
		}
	}

	fmt.Println(res)
}
