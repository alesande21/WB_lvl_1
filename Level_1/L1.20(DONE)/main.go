package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "snow dog sun — sun dog snow"

	// разделяем строку на слова, используя пробел в качестве разделителя
	// функция split возвращает срез строк
	slcStr := strings.Split(str, " ")

	// перебираем срез строк, чтобы получить перевернутую последовательность слов
	for i, j := 0, len(slcStr)-1; i < j; i, j = i+1, j-1 {
		slcStr[i], slcStr[j] = slcStr[j], slcStr[i]
	}

	// Объединяем перевернутые слова обратно в строку
	resStr := strings.Join(slcStr, " ")

	fmt.Println(resStr)

}

/*

func main() {

	str := "snow dog sun — sun dog snow"

	// разделяем строку на слова, используя пробел в качестве разделителя
	// функция split возвращает срез строк
	slcStr := strings.Split(str, " ")

	// переменная для хранения результата
	resStr := ""

	// перебираем срез строк в обратном порядке, чтобы получить перевернутую последовательность слов
	for i := len(slcStr) - 1; i >= 0; i-- {
		resStr += slcStr[i]

		// добавляем пробел между словами кроме последнего слова
		if i != 0 {
			resStr += " "
		}
	}

	fmt.Println(resStr)

}

*/
