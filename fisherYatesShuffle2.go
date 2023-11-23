package main

import (
	"fmt"
	"math/rand"
)

func knuthShuffle(arr []int) {
	
	// Определяем длину массива
    n := len(arr)

    // Цикл сортировки
    for i := 0; i < n; i++ {
        // Выбираем случайный индекс от i до конца массива, исключая i
        randomIndex := rand.Intn(i + 1)

        // Если randomIndex не равен i, меняем местами элементы на позициях i и randomIndex
        if randomIndex != i {
            arr[i], arr[randomIndex] = arr[randomIndex], arr[i]
        }
    }
}

func main() {
	// Пример использования
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Исходный массив:", arr)

	knuthShuffle(arr)
	fmt.Println("Перемешанный массив:", arr)
}
