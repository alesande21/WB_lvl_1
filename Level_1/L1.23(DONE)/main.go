package main

import "fmt"

// определяем ограничение для обобщенного типа, который может быть использован в функции
type myType interface {
	int64 | float64 | int | string
}

func myCopy[T myType](to []T, from []T, index int) {
	// переменные i и j используются для итерации по исходному и новому слайсу
	for i, j := 0, 0; i < len(from); i++ {
		// пропускаем элемент с заданным индексом
		if i != index {
			to[j] = from[i]
			j++
		}
	}
}

func deleteElement[T myType](from []T, index int) []T {
	// проверка на корректность индекса. Индекс должен быть в пределах от 1 до len(from)
	if index < 0 || index > len(from)-1 {
		return from
	}

	// cоздаем новый слайс с размером на один элемент меньше исходного
	to := make([]T, len(from)-1)

	// копируем элементы из исходного слайса в новый, исключая элемент с заданным индексом
	myCopy(to, from, index)

	return to
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr = deleteElement(arr, 2)

	fmt.Println(arr)

}

/*
func deleteElement[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
*/
