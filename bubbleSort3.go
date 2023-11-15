package main

import "fmt"

func main() {
  // Создаем массив
  arr := []int{10, 5, 2, 7, 9}
  fmt.Printf("Исходный массив: %v\n", arr)

  // Перебираем массив
  for i := len(arr) - 1; i > 0; i-- {
    // Если текущий элемент больше следующего, то меняем их местами
    for j := 0; j < i; j++ {
      if arr[j] > arr[j+1] {
        arr[j], arr[j+1] = arr[j+1], arr[j]
      }
    }
  }

  // Выводим отсортированный массив
  fmt.Printf("Отсортированный массив: %v\n", arr)
}
