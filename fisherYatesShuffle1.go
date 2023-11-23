package main

import (
	"fmt"
	"math/rand"
	"time"
)

func knuthShuffle(arr []int) {
	n := len(arr)
	rand.Seed(time.Now().UnixNano())

	for i := n - 1; i > 0; i-- {
		// Генерируем случайное число от 0 до i (включительно)
		j := rand.Intn(i + 1)

		// Меняем элементы местами
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	// Пример использования
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Исходный массив:", arr)

	knuthShuffle(arr)
	fmt.Println("Перемешанный массив:", arr)
}
