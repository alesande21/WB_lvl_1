package main

import (
	"fmt"
	"sync"
	"time"
)

func getNum(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Получено число: ", msg)
	}
}

func setNum(timeout <-chan time.Time, ch chan<- int) {
	count := 0
	for {
		select {
		case <-timeout:
			fmt.Println("Время вышло!")
		default:
			ch <- count
			count++
		}
	}

}

func main() {
	ch := make(chan int)
	timeout := time.After(6 * time.Second)
	var wg sync.WaitGroup

	go setNum(timeout, ch)
	go getNum(&wg, ch)

	wg.Wait()

}
