package main

import (
	"fmt"
	"math/big"
)

/*
	Разработать программу, которая перемножает, делит, складывает,
	вычитает две числовых переменных a, b, значение которых > 2^20.
*/

func main() {
	/*
		Инициализируем переменные a и b с помощью пакета big.Int для работы с большими числами.
		Значение a устанавливается как 11e+19, значение b устанавливается как 10e+18
	*/
	a, _ := new(big.Int).SetString("110000000000000000000", 10)

	b, _ := new(big.Int).SetString("10000000000000000000", 10)

	// Инициализация переменной operator типа big.Int для выполнения операций
	var operator big.Int

	fmt.Println("a + b =", operator.Add(a, b))

	fmt.Println("a - b =", operator.Sub(a, b))

	fmt.Println("a / b =", operator.Div(a, b))

	fmt.Println("a * b =", operator.Mul(a, b))
}
