package main

import (
    "fmt"
)

func main() {
    // Создаем массив случайных чисел
    numbers := []int{2, 5, 3, 1, 4}

    // Сортируем массив методом вставки
    insertionSort(numbers)

    // Выводим отсортированный массив
    fmt.Println(numbers)
}

func insertionSort(numbers []int) {
    // Начинаем с первого элемента массива
    n := len(numbers)
    for i := 1; i < n; i++ {
        // Берём текущий элемент
        current := numbers[i]

        // Перебираем элементы перед текущим
		j := i - 1
        for ; j >= 0; j-- {
            // Если текущий элемент меньше предыдущего, перемещаем предыдущий элемент направо
            if current < numbers[j] {
                numbers[j+1] = numbers[j]
            } else {
                // Если текущий элемент больше или равен предыдущему, выходим из цикла
                break
            }
        }

        // Вставляем текущий элемент на освободившееся место
        numbers[j+1] = current
    }
}
