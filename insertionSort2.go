package main

import "fmt"

func insertionSort(arr []int) {
    n := len(arr)
    for i := 1; i < n; i++ {
        current := arr[i]

        j := i - 1
        for j >= 0 && arr[j] > current {
            arr[j+1] = arr[j]
            j--
        }

        arr[j+1] = current
    }
}

func main() {
    // Пример неотсортированного массива
    array := []int{12, 11, 13, 5, 6}

    fmt.Println("Исходный массив:", array)

    // Вызов функции сортировки вставкой
    insertionSort(array)

    fmt.Println("Отсортированный массив:", array)
}
