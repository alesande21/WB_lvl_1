package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Введите количесво работников: \n")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}

	ch := make(chan interface{}, n)
	defer close(ch)

	for i := 0; i < n; i++ {
		go func(ch chan interface{}, worker int) {
			ch <- fmt.Sprintf("Работник %d работает", worker)
		}(ch, i+1)
	}

	for {
		data, ok := <-ch
		if ok != true {
			break
		}
		fmt.Println(data)
	}

}
