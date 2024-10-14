package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	// WaitGroup для отслеживангия завершения горутин
	var wg sync.WaitGroup

	// мьютекс для синхронизации доступа к общей переменной sum
	var mu sync.Mutex

	// переменная для хранения суммы квадратов
	var sum int

	// проходим по каждому числу в массиве
	for _, num := range arr {

		// увеличичваем счетчик
		wg.Add(1)

		// запускакем горутину для вычисления квадрата
		go func(num int) {
			defer wg.Done()

			// блокируем доступ к sum используя мьютекс
			mu.Lock()
			sum += num * num
			mu.Unlock()
		}(num)
	}

	// ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println(sum)

}
