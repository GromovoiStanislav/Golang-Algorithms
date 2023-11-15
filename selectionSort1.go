package main

import "fmt"

func selectionSort(arr []int) {
    n := len(arr)
 	
	// Циклически перебираем массив
    for i := 0; i < n-1; i++ {
		
		// Ищем минимальный элемент после текущего элемента
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }

        // Если найден минимальный элемент, то подменяем его с текущим элементом
        if minIndex != i {
            arr[i], arr[minIndex] = arr[minIndex], arr[i]
        }
        }
}

func main() {
    array := []int{4, 2, 9, 1, 7}
    fmt.Printf("Исходный массив: %v\n", array)

    selectionSort(array)

    fmt.Printf("Отсортированный массив: %v\n", array)
}
