package main

import (
	"fmt"
)

// функция принимает интерфейс и возвращает строку
func getType(elem interface{}) string {
	// 	используем fmt.Sprintf с форматом %T для получения строкового представления типа переменной
	return fmt.Sprintf("%T", elem)
}

func main() {
	fmt.Println(getType(int(2)))
	fmt.Println(getType(string("aa")))
	fmt.Println(getType(bool(true)))
	fmt.Println(getType(make(chan int)))
}
