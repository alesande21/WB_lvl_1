package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// функция worker получает данные из канала и обрабатывает их
func worker(ch chan interface{}, worker int) {
	// читаются данных из канала, пока он не закрыт
	for msg := range ch {
		fmt.Printf("Работник %d работает работу: %d\n", worker, msg)
	}
}

// функция sender отправляет данные в канал до тех пор, пока не получит сигнал завершения
func sender(ctx context.Context, ch chan interface{}) {
	i := 0
	for ch != nil { // Пока канал открыт
		select {
		// если получили сигнал завершения от контекста завершается функция
		case <-ctx.Done():
			fmt.Println("Сендер закрывается...")
			return
		default:
			// отправляем текущие значения i в канал
			ch <- i
			time.Sleep(time.Second)
			i++
		}
	}
}

func main() {
	var n int
	fmt.Printf("Введите количесво работников: \n")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}

	// создаем канал для передачи данных, размер буфера равен количеству работников
	ch := make(chan interface{}, n)

	// запускаем воркеров
	for i := 0; i < n; i++ {
		go worker(ch, i)
	}

	// создаем контекст с возможностью отмены для управления завершением работы sender
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go sender(ctx, ch)

	// создаем канал для прерывания
	interrupt := make(chan os.Signal, 1)
	defer close(interrupt)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for {

		select {
		case sig := <-interrupt:
			fmt.Printf("Приложение прерывается: %s\n", sig)

			// отправляем сигнал отмены горутине sender
			cancel()

			//даем время воркерам завершить работу
			time.Sleep(3 * time.Second)
			// закрываем канал, чтобы воркеры завершили работу
			close(ch)

			fmt.Printf("Обработчик завершил работу работу\n")
			return
		}
	}

}
