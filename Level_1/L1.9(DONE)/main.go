package main

import "fmt"

func calcNum(in <-chan int, out chan<- int) {
	// Читаем из канала in пока он не будет закрыт
	for num := range in {
		//умножаем числа на 2 и передаем данные в канал out
		out <- num * 2
	}
	// после передачи закрываем канал out
	close(out)
}

func main() {

	in := make(chan int)
	out := make(chan int)

	arr := []int{2, 4, 6, 8, 10}

	// запускаем горутину которая будет передавать все числа из массива в канал in
	go func() {
		for _, num := range arr {
			// передаем в канал in числа из массива
			in <- num
		}
		// закрываем канал in
		close(in)
	}()

	// запускаем горутину которая будет принимать числа из массива делаить расчёт и передавать в канал out
	go calcNum(in, out)

	// читаем из канала out пока он будет закрыт
	for res := range out {
		fmt.Printf("%d ", res)

	}

}

//func main() {
//	in := make(chan int)
//	arr := []int{2, 4, 6, 8, 10}
//	out := make(chan int)
//
//	go func() {
//		for _, n := range arr {
//			in <- n
//		}
//		in = nil
//	}()
//
//	go func() {
//		for in != nil {
//			n := <-in
//			out <- n * 2
//		}
//		out = nil
//	}()
//
//	for out != nil {
//		num := <-out
//		println(num)
//	}
//}
