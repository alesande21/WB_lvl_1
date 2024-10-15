package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// checkUnique проверяет, что все символы в строке уникальны
func checkUnique(str string) bool {
	// создаем пусту карту для хранения символов
	m := make(map[rune]struct{})

	// преобразуем строку к нижнему регистру, чтобы сделать проверку регистронезависимой
	str = strings.ToLower(str)

	// проходим по каждому символу
	for i := 0; i < len(str); {

		// получаем символ в виде руны и его ширину в байтах
		runeValue, width := utf8.DecodeRuneInString(str[i:])

		// проверяем встречался ли этот символ ранее
		_, founded := m[runeValue]
		if founded {
			// если встречался возвращаем false
			return false
		} else {
			// если не встречался то добавляем в карту
			m[runeValue] = struct{}{}
		}
		//fmt.Printf("%#U starts at byte position %d\n", runeValue, i-width)

		i += width

	}
	// если все символы уникальны, возвращаем true
	return true
}

func main() {

	fmt.Println(checkUnique("abcd"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("aabcd"))

}

/*
abcd — true

abCdefAaf — false

aabcd — false
*/
