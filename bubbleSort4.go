package main

import "fmt"

func bubbleSort(arr []int) {
    n := len(arr)
    swapped := true

    for swapped {
        swapped = false

        for i := 0; i < n-1; i++ {
            if arr[i] > arr[i+1] {
                // Обмен значениями, если текущий элемент больше следующего
                arr[i], arr[i+1] = arr[i+1], arr[i]
                swapped = true
            }
        }
    }
}

func main() {
    // Пример массива
    array := []int{64, 34, 25, 12, 22, 11, 90}

    fmt.Println("Исходный массив:", array)

    // Вызов функции для сортировки пузырьком
    bubbleSort(array)

    fmt.Println("Массив после сортировки пузырьком:", array)
}