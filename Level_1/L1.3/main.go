package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var sum int
	for _, num := range arr {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			sum += num * num
			mu.Unlock()
		}(num)
	}

	wg.Wait()

	fmt.Println(sum)

}
