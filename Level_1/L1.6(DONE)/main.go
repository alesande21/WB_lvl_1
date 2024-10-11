package main

import (
	"context"
	"fmt"
	"time"
)

// Пример с каналом.
func worker1(stop <-chan bool) {
	for {
		select {
		case <-stop:
			// Остановка работника при получении сигнала из канала
			fmt.Println("Работник 1 остановился")
			return
		default:
			// Работа выполняется в горутине
			fmt.Println("Работник 1 рабает")
			time.Sleep(1 * time.Second)

		}
	}
}

func runExample1() {
	// Канал для передачи сигнала об остановке горутины
	stop := make(chan bool)

	go worker1(stop)

	time.Sleep(6 * time.Second)
	// Останавливаем горутиину передавая сигнал
	stop <- true
	defer close(stop)

	time.Sleep(1 * time.Second)
	fmt.Println("Работник 1 остановлен")
}

// Пример с контекстом
func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Остановка работника при завершении контекста
			fmt.Println("Работник 2 остановился")
			return
		default:
			// Работа выполняется в горутине
			fmt.Println("Работник 2 рабает")
			time.Sleep(1 * time.Second)

		}
	}
}

func runExample2() {
	// создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// запускаем горутину с этим контекстом
	go worker2(ctx)

	// ждём
	time.Sleep(6 * time.Second)

	// Останавливаем горутиину отменяя контекст
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Работник 2 остановлен")
}

// Пример с таймаутом
func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Остановка работника при завершении времени
			fmt.Println("Работник 3 остановился")
			return
		default:
			// Работа выполняется в горутине
			fmt.Println("Работник 3 рабает")
			time.Sleep(1 * time.Second)

		}
	}
}

func runExample3() {
	// создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)

	// запускаем горутину с этим контекстом
	go worker3(ctx)

	// ждём
	time.Sleep(6 * time.Second)

	// Горутина отсановилась раньше отмены
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Работник 3 остановлен")
}

func main() {
	runExample1()
	runExample2()
	runExample3()
}
