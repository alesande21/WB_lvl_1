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

// broadcaster получает сообщение из главного канала и распространяет его по всем работникам
func broadcaster(chs []chan interface{}, mainCh chan interface{}) {
	// читаются данных из канала, пока он не закрыт
	for msg := range mainCh {
		// при получении сообщения пересылается всем работникам
		for _, chWorker := range chs {
			chWorker <- msg
		}
	}

	// при закрытии главного канала закрываются и каналы работников
	for _, chWorker := range chs {
		close(chWorker)
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
	chMain := make(chan interface{}, n)
	chs := make([]chan interface{}, 0)

	// запускаем воркеров
	for i := 0; i < n; i++ {
		chWorker := make(chan interface{})
		chs = append(chs, chWorker)
		go worker(chWorker, i)
	}

	// запускаем вещатель для ретрансляции сообщения
	go broadcaster(chs, chMain)

	// создаем контекст с возможностью отмены для управления завершением работы sender
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// запускаем сендер для трансляции сообщений в главный канал
	go sender(ctx, chMain)

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
			close(chMain)

			fmt.Printf("Обработчик завершил работу работу\n")
			return
		}
	}

}
