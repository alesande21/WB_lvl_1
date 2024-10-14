package main

import "fmt"

func binarySearch(arr []int, target int) bool {
	// вызываем вспомогательную функцию передавая туда начальные значения правого и левого указателя
	return binarySearch2(arr, 0, len(arr)-1, target)
}

func binarySearch2(arr []int, l, r, target int) bool {
	// если левый указатель становится больше правого поиск заканчивается
	for l <= r {

		// находим середину текущего диапазона
		mid := l + (r-l)/2

		// если текущий элемент равен искомому выходим из цикла
		if arr[mid] == target {
			return true
			// если текущий элемент меньше искомого то левый указатель сдвигаем вправо
		} else if arr[mid] < target {
			l = mid + 1
			// если текущий элемент больше искомого то правый указатель сдвигаем влево
		} else {
			r = mid - 1
		}
	}

	// если элемент не найден возвращаем false
	return false
}

func main() {

	arr := []int{2, 3, 4, 7, 10, 17, 21, 40}

	fmt.Println(binarySearch(arr, 23))

}
