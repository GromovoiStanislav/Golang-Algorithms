package main

import "fmt"

func reverse(array []int) []int {
    // Создаём новый массив
    newArray := make([]int, len(array))

    // Перебираем элементы массива в обратном порядке
    for i := len(array) - 1; i >= 0; i-- {
        // Копируем элемент в новый массив
        newArray[len(newArray)-i-1] = array[i]
    }

    return newArray
}

func main(){
	array := []int{1, 2, 3, 4, 5}

    newArray := reverse(array)

    fmt.Println(newArray)
}