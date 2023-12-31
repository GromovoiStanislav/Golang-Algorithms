package main

import "fmt"

func main() {
  // Создаем массив
  arr := []int{10, 5, 2, 7, 9}

  // Инициализируем переменную для хранения минимального значения
  min := arr[0]

  // Циклически перебираем массив
  for i := 1; i < len(arr); i++ {
    // Если текущее значение меньше минимального, то обновляем минимальное значение
    if arr[i] < min {
      min = arr[i]
    }
  }

  // Выводим минимальное значение
  fmt.Println("Минимальное значение:", min)
}
