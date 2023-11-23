package main

import "fmt"

func reverseArray(arr []int) {
    n := len(arr)
    for i := 0; i < n/2; i++ {
        // Обмен значений между элементами с обеих сторон
        arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
    }
}

func main() {
    // Пример массива
    array := []int{1, 2, 3, 4, 5}

    fmt.Println("Исходный массив:", array)

    // Вызов функции для реверса массива
    reverseArray(array)

    fmt.Println("Массив после реверса:", array)
}
