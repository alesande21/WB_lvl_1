package main

import "fmt"

func main() {
	// Использование struct{} позволяет нам хранить ключи без выделения лишней памяти
	m := make(map[string]struct{})
	str := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, s := range str {
		// если у нас уже есть строка в мнодестве то она не будет повторна добавлена
		m[s] = struct{}{}
	}

	for key := range m {
		fmt.Println(key)
	}
}
