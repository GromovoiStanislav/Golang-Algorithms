package main

import "fmt"

func shakerSort(array []int) {
    // Определяем начальные и конечные индексы массива
    left := 0
    right := len(array) - 1

    // Цикл сортировки
    for left < right {
        // Перебираем элементы массива слева направо
        for i := left; i < right; i++ {
            // Если текущий элемент больше следующего, меняем их местами
            if array[i] > array[i+1] {
                array[i], array[i+1] = array[i+1], array[i]
            }
        }

        // Уменьшаем правый индекс
        right--

        // Перебираем элементы массива справа налево
        for i := right; i > left; i-- {
            // Если текущий элемент меньше предыдущего, меняем их местами
            if array[i] < array[i-1] {
                array[i], array[i-1] = array[i-1], array[i]
            }
        }

        // Увеличиваем левый индекс
        left++
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
