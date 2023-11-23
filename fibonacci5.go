package main

import (
	"fmt"
	"math/big"
)

func fibonacciSequenceBig(n int) []*big.Int {
	if n <= 0 {
		return []*big.Int{big.NewInt(0)}
	} else if n == 1 {
		return []*big.Int{big.NewInt(0), big.NewInt(1)}
	}

	sequence := []*big.Int{big.NewInt(0), big.NewInt(1)}
	for i := 2; i < n; i++ {
		nextNumber := new(big.Int)
		nextNumber.Add(sequence[i-1], sequence[i-2])
		sequence = append(sequence, nextNumber)
	}

	return sequence
}

func main() {
	n := 100
	result := fibonacciSequenceBig(n)
	fmt.Printf("Последовательность Фибоначчи из %d чисел: %v\n", n, result)
	
	//fmt.Printf("Последовательность Фибоначчи из %d чисел:\n", n)
	// for _, num := range result {
	// 	fmt.Println(num)
	// }
}
