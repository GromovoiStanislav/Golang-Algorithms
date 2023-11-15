package main

import "fmt"

func main() {
  // Создаем массив
  arr := []int{10, 5, 2, 7, 9}
  fmt.Printf("Исходный массив: %v\n", arr)

  // Циклически перебираем массив
  for i := 0; i < len(arr)-1; i++ {
    
    // Ищем минимальный элемент после текущего элемента
    minIndex := i
    for j := i + 1; j < len(arr); j++ {
      if arr[j] < arr[minIndex] {
        minIndex = j
      }
    }

    // Если найден минимальный элемент, то подменяем его с текущим элементом
    if minIndex != i {
      temp := arr[i]
      arr[i] = arr[minIndex]
      arr[minIndex] = temp
    }
  }

  // Выводим отсортированный массив
  fmt.Printf("Отсортированный массив: %v\n", arr)
}
