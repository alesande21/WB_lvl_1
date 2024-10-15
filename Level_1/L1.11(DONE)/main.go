package main

import "fmt"

func main() {
	arr1 := []int{4, 2, 3, 6, 1}
	arr2 := []int{10, 1, 8, 2, 3, 4}
	var res []int
	mySet := make(map[int]struct{})

	// проходим по элементам первого слайса и заполняем мап
	for _, elem := range arr1 {
		mySet[elem] = struct{}{}
	}

	// проходим по элементам второго слайса
	for _, elem := range arr2 {
		// ищем элемент в мапе
		_, founded := mySet[elem]
		// если он найден добавляем в результирующий слайс
		if founded {
			res = append(res, elem)
		}
	}

	fmt.Println(res)
}
