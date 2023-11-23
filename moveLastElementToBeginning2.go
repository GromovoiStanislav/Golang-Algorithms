package main

import "fmt"

func moveLastToFirst(arr []int) {
   // Получаем последний элемент массива
   lastElement := arr[len(arr)-1]

   // Сдвигаем элементы массива влево
   for i := len(arr) - 2; i >= 0; i-- {
    arr[i+1] = arr[i]
   }

   // Вставляем последний элемент в начало
   arr[0] = lastElement
}

func main() {
    // Пример массива
    array := []int{1, 2, 3, 4, 5}

    fmt.Println("Исходный массив:", array)

    // Вызов функции для перемещения последнего элемента в начало
    moveLastToFirst(array)

    fmt.Println("Массив после перемещения последнего элемента в начало:", array)
}
