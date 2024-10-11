package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	// создаем WaitGroup для ожидания завержения всех запущеных горутин
	var wg sync.WaitGroup
	for _, num := range arr {
		// увеличиваем счетчик WaitGroup, так как будем запускать новую горутину
		wg.Add(1)

		// запускаем горутину
		go func(num int) {
			// уменьшаем счетчик WaitGroup после завершения выполнения горутины
			defer wg.Done()
			fmt.Println(num * num)
		}(num)
	}

	// ждем пока завершаться все горутины
	wg.Wait()

}
