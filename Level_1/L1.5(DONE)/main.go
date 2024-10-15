package main

import (
	"fmt"
	"sync"
	"time"
)

// функция будет читать канал до тех пор пока он не будет закрыт
func getNum(wg *sync.WaitGroup, ch <-chan int) {
	// уменьшаем счетчик WaitGroup, когда функция завершается
	defer wg.Done()
	for {
		// читаем значение из канала ch
		msg, ok := <-ch
		// если канал закрыт и больше нет значений для чтения, выходим из цикла
		if !ok {
			break
		}
		fmt.Println("Получено число: ", msg)
	}
}

// функция setNum последовательно отправляет значения в канал до истечения времени
func setNum(timeout <-chan time.Time, ch chan<- int) {
	count := 0
	for {
		select {
		// если получили сигнал от таймера, завершаем отправку в канал
		case <-timeout:
			fmt.Println("Время вышло!")
			// закрываем канал чтобы сообщить что время вышло
			close(ch)
			return
		default:
			// отправляем текущее значение в канал
			ch <- count
			count++
			time.Sleep(time.Second)
		}
	}

}

func main() {
	var sec int
	fmt.Printf("Введите через сколько секунд завершить программу: ")
	_, err := fmt.Scanf("%d", &sec)
	if err != nil {
		fmt.Printf("\nОшибка при вводе данных: %s\n", err)
		return
	}
	ch := make(chan int)
	duration := time.Duration(sec) * time.Second
	timeout := time.After(duration)
	var wg sync.WaitGroup
	wg.Add(1)

	// запускаем горутину setNum для отправки данных в канал
	go setNum(timeout, ch)

	// запускаем горутину getNum для чтения данных из канала
	go getNum(&wg, ch)

	// Ожидаем завершения всех горутин
	wg.Wait()

}
