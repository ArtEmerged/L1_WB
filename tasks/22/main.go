package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("110000000000000000000", 10) // 11e+19
	b, _ := new(big.Int).SetString("10000000000000000000", 10)  // 10e+18
	var operator big.Int
	fmt.Println("a + b =", operator.Add(a, b))
	fmt.Println("a - b =", operator.Sub(a, b))
	fmt.Println("a / b =", operator.Div(a, b))
	fmt.Println("a * b =", operator.Mul(a, b))
}
