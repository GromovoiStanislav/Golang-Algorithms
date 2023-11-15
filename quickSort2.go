package main

import "fmt"

func main() {
  // Создаем массив
  arr := []int{10, 5, 2, 7, 9}
  fmt.Printf("Исходный массив: %v\n", arr)

  // Сортируем массив
  quickSort(arr, 0, len(arr)-1)

  // Выводим отсортированный массив
  fmt.Printf("Отсортированный массив: %v\n", arr)
}

func quickSort(arr []int, left int, right int) {
  // Если массив пуст или содержит один элемент, то мы возвращаемся
  if left >= right {
    return
  }

  // Выбираем опорный элемент
  pivot := arr[(left + right) / 2]

  // Разделяем массив на две части
  i := left
  j := right
  for i <= j {
    // Пока текущий элемент меньше опорного, мы увеличиваем индекс
    for i <= j && arr[i] < pivot {
      i++
    }

    // Пока текущий элемент больше опорного, мы уменьшаем индекс
    for i <= j && arr[j] > pivot {
      j--
    }

    // Если индексы не встретились, то мы меняем их местами
    if i <= j {
      arr[i], arr[j] = arr[j], arr[i]
      i++
      j--
    }
  }

  // Рекурсивно сортируем левое и правое подмассивы
  quickSort(arr, left, j)
  quickSort(arr, i, right)
}
