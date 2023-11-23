package main

import "fmt"

func moveLastToFirst(arr []int) {
    n := len(arr)
    if n > 1 {
        // Сохраняем значение последнего элемента
        lastElement := arr[n-1]

        // Сдвигаем все элементы вправо
        for i := n - 1; i > 0; i-- {
            arr[i] = arr[i-1]
        }

        // Помещаем сохраненное значение в начало
        arr[0] = lastElement
    }
}

func main() {
    // Пример массива
    array := []int{1, 2, 3, 4, 5}

    fmt.Println("Исходный массив:", array)

    // Вызов функции для перемещения последнего элемента в начало
    moveLastToFirst(array)

    fmt.Println("Массив после перемещения последнего элемента в начало:", array)
}
