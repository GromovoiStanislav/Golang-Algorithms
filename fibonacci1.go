package main

import (
	"fmt"
)

//Только до 93 !!!
func fibonacci(n int) int {
	if n <= 0 {
		fmt.Println("Пожалуйста, введите положительное число для вычисления числа Фибоначчи.")
		return -1 // Возврат -1 как индикатор ошибки
	} else if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		prevPrev := 0
		prev := 1

		for i := 3; i <= n; i++ {
			current := prev + prevPrev
			prevPrev = prev
			prev = current
		}

		return prev
	}
}

func main() {
	n := 93
	result := fibonacci(n)
	if result != -1 {
		fmt.Printf("Фибоначчи %d = %d\n", n, result)
	}
}
