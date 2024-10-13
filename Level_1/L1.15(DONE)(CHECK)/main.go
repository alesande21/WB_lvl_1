package main

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"
)

var justString string

func createHugeString(n int) string {
	// используем strings.Repeat для генерации строки длиной n.
	str := strings.Repeat("a", n)
	return str
}

func deepCopy(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}

func someFunc() {
	// мы не знаем какая строка приходит из метода
	v := createHugeString(1 << 10)
	// проверяем длину строки
	fmt.Printf("Длина v: %d\n", len(v))
	if len(v) >= 100 {
		// если не проверять длину строки и она будет меньше 100 то вызывается ошибка
		// panic: runtime error: slice bounds out of range [:100] with length 99

		//var sb strings.Builder
		//var err error
		//_, err = sb.WriteString(v[:100])
		//if err != nil {
		//	fmt.Printf("не удалось прочитать строку: %s", err)
		//	return
		//}
		//justString = sb.String()
		justString = v[:100]

	}
	fmt.Printf("Длина среза: %d\n", len(justString))

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Используемая память до GC: %d байт\n", memStats.Alloc)

	v = ""

	runtime.GC()

	runtime.ReadMemStats(&memStats)
	fmt.Printf("Используемая память после GC: %d байт\n", memStats.Alloc)
}

func main() {
	someFunc()
}
