package main

import "fmt"

/*
	Поменять местами два числа без создания временной переменной
*/

func main() {
	var a, b int
	fmt.Print("Веедите первое число a=")
	fmt.Scan(&a)
	fmt.Print("Веедите второе число b=")
	fmt.Scan(&b)

	// Выполняем операцию параллельного присваивания
	a, b = b, a

	fmt.Printf("a=%d b=%d\n", a, b)
}
