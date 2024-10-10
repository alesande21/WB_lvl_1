package main

import (
	"fmt"
	"sync"
)

type ConcurrentCache struct {
	*cache
}

type cache struct {
	orders       map[string]int
	sync.RWMutex // RWMutex позволяет управлять доступом для конкурентного чтения и записи.
}

func NewConcurrentCache() *ConcurrentCache {
	orders := make(map[string]int)
	c := cache{
		orders:  orders,
		RWMutex: sync.RWMutex{},
	}
	return &ConcurrentCache{&c}
}

// Используется мьютекс для блокировки записи, чтобы гарантировать безопасность данных. Только одна горутина может читать или писать.
func (c *ConcurrentCache) Set(k string, value int) {
	c.Lock()         // Блокируем доступ на запись.
	defer c.Unlock() // Разблокируем после завершения функции.
	c.orders[k] = value
}

// Используеся блокировка для чтения, чтобы предотвратить конкурентные записи во время чтения. Может несколько горутин читать, но не записывать.
func (c *ConcurrentCache) ItemCount() int {
	c.RLock()         // Блокируем доступ на чтение.
	defer c.RUnlock() // Разблокируем после завершения функции.
	n := len(c.orders)
	return n
}

func (c *ConcurrentCache) Get(k string) (int, bool) {
	c.RLock()
	defer c.RUnlock()
	val, found := c.orders[k]
	if !found {
		return 0, false
	}
	return val, true
}

func main() {

	myMap := NewConcurrentCache()
	var wg sync.WaitGroup

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("Значение равно %d ==", i)
			go myMap.Set(key, i)
		}(i)

	}

	wg.Wait()

	for i := 7; i >= 0; i-- {
		key := fmt.Sprintf("Значение равно %d ==", i)
		val, find := myMap.Get(key)
		if find == true {
			fmt.Println(key, val)
		} else {
			fmt.Printf("Значение с ключом %s не найдено\n", key)
		}
	}

}
