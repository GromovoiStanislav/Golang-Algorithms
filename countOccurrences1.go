package main

import "fmt"

func countOccurrences(arr []int) map[int]int {
    occurrences := make(map[int]int)

    for _, num := range arr {
        occurrences[num]++
    }

    return occurrences
}

func main() {
    array := []int{4, 2, 9, 1, 7, 4, 2, 4}
    occurrences := countOccurrences(array)

    fmt.Println("Число встреч каждого элемента:")
    for num, count := range occurrences {
        fmt.Printf("%d: %d раз(а)\n", num, count)
    }
}
