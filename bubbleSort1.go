package main

import "fmt"

func bubbleSort(arr []int) {
    n := len(arr)

    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Обмен значениями
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    array := []int{4, 2, 9, 1, 7}
    fmt.Printf("Исходный массив: %v\n", array)

    bubbleSort(array)

    fmt.Printf("Отсортированный массив: %v\n", array)
}
