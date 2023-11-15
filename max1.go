package main

import "fmt"

func findMax(arr []int) int {
    if len(arr) == 0 {
        return 0 // Или другое значение по умолчанию в зависимости от контекста
    }

    max := arr[0]
    for _, num := range arr {
        if num > max {
            max = num
        }
    }

    return max
}

func main() {
    array := []int{4, 2, 9, 1, 7}
    maxValue := findMax(array)
    fmt.Printf("Максимальное значение в массиве: %d\n", maxValue)
}
