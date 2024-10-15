package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	// используем strings.Builder для выделения строки без использования дополнительных буферов
	var b strings.Builder
	// Вызов метода Grow выделяет срез емкостью n
	b.Grow(n)
	for i := 0; i < n; i++ {
		// Вызовы WriteByte заполняют срез до емкости
		//b.WriteByte(0)
		// добавляем многобайтовый символ в строку
		fmt.Fprintf(&b, "у")
	}
	// Метод String() использует unsafe для преобразования этого среза в строку.
	s := b.String()
	return s
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Printf("Длина v: %d\n", len(v))

	// преобразуем строку в срез рун, чтобы корректно работать с многобайтовыми символами
	r := []rune(v)

	//Обрезаем строку и сохраняем первые 100 рун
	justString = string(r[:4])

	fmt.Println(justString)
}

func main() {
	someFunc()
}
