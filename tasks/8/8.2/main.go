package main

import "fmt"

func main() {
	fmt.Printf("Введите через пробел чисто(int64) и индек бита(i) от 0 до 63\nПример:42 53 \nВвод:")
	var numb, index int64
	fmt.Scanf("%d %d", &numb, &index)

	if index > 63 || index < 0 {
		fmt.Println("Невалидный индекс")
		return
	}
	
	fmt.Printf("%064b\n", uint64(numb))
	numb = inverseBit(numb, index)
	fmt.Printf("%064b\n", uint64(numb))
}

/*
	^ - это логическая операция XOR (исключающее ИЛИ) - равен 1, когда один их битов a или b равен 1. В остальных ситуация равен 0
	Пример:
		01100
		11010
		_____
	   ^10110
*/

// инверсия бита
func inverseBit(number, i int64) int64 {
	/*
		Выполняем алгоритмический сдвиг числа 1 (00...01) влево на i-й бит
		Пример:
		00000001 << 3 = 00001000
		Применяем логическую операцию XOR к number и bitPosition
		Пример:
		Мы поменяли 1 в 3 индексе на 0
			01001001
			00001000
		   ^01000001

		Мы поменяли 1 в 3 индексе на 1
			01110111
			00001000
		   ^01111111
	*/
	var bitPosition int64 = 1 << i
	return number ^ bitPosition
}
