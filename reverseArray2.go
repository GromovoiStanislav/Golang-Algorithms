package main

import "fmt"

func reverse(array []int)  {
    // Перебираем элементы массива
    for i := 0; i < len(array)/2; i++ {
        // Меняем местами элементы на позициях i и len(array)-i-1
        swap(array, i, len(array)-i-1)
    }

}

func swap(array []int, i, j int) {
    // Меняем местами элементы на позициях i и j
    temp := array[i]
    array[i] = array[j]
    array[j] = temp
}

func main(){
	array := []int{1, 2, 3, 4, 5}

	reverse(array)
	
	fmt.Println(array)
}