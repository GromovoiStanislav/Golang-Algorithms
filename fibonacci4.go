package main

import (
	"fmt"
)

//Только до 93 !!!
func fibonacciSequence(n int) []int {
	if n <= 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	}

	sequence := []int{0, 1}
	for i := 2; i < n; i++ {
		sequence = append(sequence, sequence[i-1]+sequence[i-2])
	}
	return sequence
}

func main() {
	n := 93
	result := fibonacciSequence(n)
	fmt.Printf("Последовательность Фибоначчи из %d чисел: %v\n", n, result)

	//fmt.Printf("Последовательность Фибоначчи из %d чисел:\n", n)
	// for _, num := range result {
	// 	fmt.Println(num)
	// }
}
