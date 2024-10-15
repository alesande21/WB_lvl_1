package main

import (
	"fmt"
	"math/big"
)

func main() {
	// инициализация числовых значений, которые больше чем 2^20
	a := new(big.Int)
	b := new(big.Int)

	// устанавливаем значения для переменных
	a.SetString("1048576", 10) // 2^20
	b.SetString("2097152", 10) // 2^21

	// переменные для хранения результатов операций
	sum := new(big.Int)
	diff := new(big.Int)
	product := new(big.Int)
	div := new(big.Int)

	// выполняем сложение
	sum.Add(a, b)
	fmt.Printf("Сумма: %s + %s = %s\n", a.String(), b.String(), sum.String())

	// выполняем вычитание
	diff.Sub(a, b)
	fmt.Printf("Разность: %s - %s = %s\n", a.String(), b.String(), diff.String())

	// выполняем умножение
	product.Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a.String(), b.String(), product.String())

	// выполняем деление
	// обрабатываем случай, когда b = 0
	if b.Cmp(big.NewInt(0)) != 0 {
		div.Div(a, b)
		fmt.Printf("Деление: %s / %s = %s\n", a.String(), b.String(), div.String())
	} else {
		fmt.Println("Ошибка деление на ноль")
	}
}
