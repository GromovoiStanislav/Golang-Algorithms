package main

import "fmt"

func main() {
  // Создаем массив
  arr := []int{10, 5, 2, 7, 9}

  // Инициализируем переменную для хранения максимального значения
  max := arr[0]

  // Циклически перебираем массив
  for i := 1; i < len(arr); i++ {
    // Если текущее значение больше максимального, то обновляем максимальное значение
    if arr[i] > max {
      max = arr[i]
    }
  }

  // Выводим максимальное значение
  fmt.Println("Максимальное значение:", max)
}
