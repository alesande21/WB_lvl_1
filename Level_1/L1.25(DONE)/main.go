package main

import (
	"fmt"
	"time"
)

// mySleep приостанавливает выполнение программы на заданную длительность
func mySleep(d time.Duration) {
	// Используем канал из функции time.After, который блокирует выполнение до истечения указанного времени
	<-time.After(d)

}

func main() {

	fmt.Println("Старт")
	mySleep(time.Second * 4)
	fmt.Println("Финиш")

}
