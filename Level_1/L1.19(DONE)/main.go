package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var str string
	var reverse string

	// используем strings.Builder для выделения строки без использования дополнительных буферов
	var b strings.Builder

	// считываем строку
	fmt.Printf("Введите строку: ")
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		fmt.Printf("Ошибка при вводе: %s\n", err)
		return
	}

	// проходим по строке в обратном порядке чтобы перевернуть её
	for i := len(str); i > 0; {
		// используем utf8.DecodeLastRuneInString для декодирования последней руны
		// это позволяет правильно обрабатывать многобайтовые символы Unicode
		runeValue, width := utf8.DecodeLastRuneInString(str[:i])

		fmt.Printf("%#U starts at byte position %d\n", runeValue, i-width)

		b.WriteRune(runeValue)
		// сдвигаемся влево на ширину текущей руны
		i -= width

		// добавляем руну к перевернутой строке
		//reverse += string(runeValue)
	}

	reverse = b.String()

	fmt.Println(reverse)
}
