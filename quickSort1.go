package main

import "fmt"

func quickSort(arr []int, low, high int) {
    if low < high {
        partitionIndex := partition(arr, low, high)

        quickSort(arr, low, partitionIndex-1)
        quickSort(arr, partitionIndex+1, high)
    }
}

func partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1

    for j := low; j <= high-1; j++ {
        if arr[j] <= pivot {
            i++
            // Обмен значениями
            arr[i], arr[j] = arr[j], arr[i]
        }
    }

    // Обмен значениями
    arr[i+1], arr[high] = arr[high], arr[i+1]

    return i + 1
}

func main() {
    array := []int{4, 2, 9, 1, 7}
    fmt.Printf("Исходный массив: %v\n", array)

    quickSort(array, 0, len(array)-1)

    fmt.Printf("Отсортированный массив: %v\n", array)
}
