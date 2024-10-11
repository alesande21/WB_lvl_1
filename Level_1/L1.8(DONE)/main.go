package main

import (
	"fmt"
)

func main() {
	var (
		num int64
		pos int
		bit int
	)

	fmt.Printf("Введите число: ")
	_, err := fmt.Scanf("%d\n", &num)
	if err != nil {
		fmt.Printf("Не удалось обработать число: %s", err)
		return
	}
	fmt.Printf("Введите позицию которую необходимо изменить: ")
	_, err = fmt.Scanf("%d\n", &pos)
	if err != nil {
		fmt.Printf("Не удалось обработать число: %s", err)
		return
	}

	if pos < 0 || pos > 63 {
		fmt.Printf("Неверная позиция: %d", bit)
		return
	}

	fmt.Printf("Введите бит: ")
	_, err = fmt.Scanf("%d\n", &bit)
	if err != nil {
		fmt.Printf("Не удалось обработать число: %s", err)
		return
	}

	if bit < 0 || bit > 1 {
		fmt.Printf("Неверное значение бита должен быть 0 или 1: %d", bit)
		return
	}

	// Устанавливаем значение бита в указанной позиции
	// Если бит равен 1, используем побитовую операцию OR (|=) для установки бита
	// 1 | 1 = 1
	// 0 | 1 = 1
	if bit == 1 {
		num |= 1 << pos
	} else {
		// Если бит равен 0, используем операцию AND NOT (&^) для сброса бита
		// 1 &^ 1 = 0
		// 0 &^ 1 = 0
		num &^= 1 << pos
	}

	fmt.Printf("%064b\n", num)
	fmt.Printf("%d\n", num)

}
