package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Сounter struct {
	count int32
}

// Increase метод счетчика увеличивает значение счетчика count на 1, используя атомарную операцию
func (c *Сounter) Increase() {
	atomic.AddInt32(&c.count, 1)
}

// Get метод счетчика используя атомарный метод возвращает значение счетчика count
func (c *Сounter) Get() int32 {
	return atomic.LoadInt32(&c.count)
}

var counter Сounter

func main() {
	// создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// запускаем 1000 горутин, каждая из которых увеличивает значение счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increase()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Get())
}
