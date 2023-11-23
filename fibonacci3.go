package main

import (
	"fmt"
	"math/big"
)

func fibonacciBig(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	} else if n == 1 {
		return big.NewInt(1)
	}

	a := big.NewInt(0)
	b := big.NewInt(1)
	result := new(big.Int)

	for i := 2; i < n; i++ {
		result.Add(a, b)
		a.Set(b)
		b.Set(result)
	}

	return result
}

func main() {
	n := 100
	result := fibonacciBig(n)
	fmt.Printf("Число Фибоначчи номер %d: %v\n", n, result)
}
