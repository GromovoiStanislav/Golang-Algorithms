package main

import (
	"fmt"
	"math/big"
)

func fibonacciBig(n int) *big.Int {
	if n <= 0 {
		fmt.Println("Пожалуйста, введите положительное число для вычисления числа Фибоначчи.")
		return nil
	} else if n == 1 {
		return big.NewInt(0)
	} else if n == 2 {
		return big.NewInt(1)
	} else {
		prevPrev := big.NewInt(0)
		prev := big.NewInt(1)

		for i := 3; i <= n; i++ {
			current := new(big.Int).Add(prev, prevPrev)
			prevPrev.Set(prev)
			prev.Set(current)
		}

		return prev
	}
}

func main() {
	n := 10
	result := fibonacciBig(n)
	if result != nil {
		fmt.Printf("Фибоначчи %d = %s\n", n, result.String())
	}
}
