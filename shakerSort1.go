package main

import "fmt"

func shakerSort(arr []int) {
    n := len(arr)
    swapped := true

    for swapped {
        swapped = false

        // Проход слева направо
        for i := 0; i < n-1; i++ {
            if arr[i] > arr[i+1] {
                arr[i], arr[i+1] = arr[i+1], arr[i]
                swapped = true
            }
        }

        // Если не было обменов, сортировка завершена
        if !swapped {
            break
        }

        swapped = false

        // Проход справа налево
        for i := n - 1; i > 0; i-- {
            if arr[i-1] > arr[i] {
                arr[i-1], arr[i] = arr[i], arr[i-1]
                swapped = true
            }
        }
    }
}

func main() {
    // Пример массива
    array := []int{64, 34, 25, 12, 22, 11, 90}

    fmt.Println("Исходный массив:", array)

    // Вызов функции для шейкерной сортировки
    shakerSort(array)

    fmt.Println("Массив после шейкерной сортировки:", array)
}
