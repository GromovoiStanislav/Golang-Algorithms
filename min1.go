package main

import "fmt"

func findMin(arr []int) int {
    if len(arr) == 0 {
        return 0 // Или другое значение по умолчанию в зависимости от контекста
    }

    min := arr[0]
    for _, num := range arr {
        if num < min {
            min = num
        }
    }

    return min
}

func main() {
    array := []int{4, 2, -9, 1, 7}
    minValue := findMin(array)
    fmt.Printf("Минимальное значение в массиве: %d\n", minValue)
}