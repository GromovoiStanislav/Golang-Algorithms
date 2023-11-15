package main

import "fmt"

func createIndexHash(arr []int) map[int][]int {
    indexHash := make(map[int][]int)

    for i, num := range arr {
        indexHash[num] = append(indexHash[num], i)
    }

    return indexHash
}

func main() {
    array := []int{4, 2, 9, 1, 7, 4, 2, 4}
    indexHash := createIndexHash(array)

    fmt.Println("Хэш-таблица с индексами для каждого элемента:")
    for num, indices := range indexHash {
        fmt.Printf("%d: %v\n", num, indices)
    }
}
