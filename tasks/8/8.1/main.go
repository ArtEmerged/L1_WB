package main

import "fmt"

/*
	Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
*/

// Функция setBit устанавливает i-й бит в 1.
func setBit(number, i int64) int64 {

	/*
		Выполняем алгоритмический сдвиг числа 1 (00...01) влево на i-й бит
		Пример:
		00000001 << 3 = 00001000
		Применяем логическую операцию OR к number и bitPosition
		Пример:
		Мы поменяли 1 в 3 индексе на 1
			01001001
			00001000
		|	01001001

		Мы поменяли 0 в 3 индексе на 1
			01110111
			00001000
		|	01111111
	*/
	bitPosition := int64(1) << i
	return number | bitPosition
}

// Функция resetBit устанавливает i-й бит в 0.
func resetBit(number, i int64) int64 {
	/*
		Выполняем алгоритмический сдвиг числа 1 (00...01) влево на i-й бит
		Пример:
		00000001 << 3 = 00001000
		Применяем логическую операцию NOT AND к number и bitPosition
		^ инверсирует биты
		Пример:
		Мы поменяли 0 в 3 индексе на 0
			01000001
		^	11110111
		&	01000001

		Мы поменяли 1 в 3 индексе на 0
			01111111
		^	11110111
		&	01110111
	*/
	bitPosition := int64(1) << i
	return number &^ bitPosition
}

func main() {

	fmt.Printf("Введите через пробел число (int64), индекс бита (i) от 0 до 63 и значение бита (0 или 1)\nПример: 42 53 1\nВвод: ")
	var number, index, bit int64
	_, err := fmt.Scanf("%d %d %d", &number, &index, &bit)

	// Проверяем ввод на валидность
	if err != nil {
		fmt.Println("Невалидные данные")
		return
	}

	// Проверяем индекс на валидность
	if index > 63 || index < 0 {
		fmt.Println("Невалидный индекс")
		return
	}

	fmt.Printf("%064b\n", uint64(number))

	// Если значение бита равно 1, устанавливаем бит в 1
	if bit == 1 {
		number = setBit(number, index)
	}
	// Если значение бита равно 0, устанавливаем бит в 0
	if bit == 0 {
		number = resetBit(number, index)
	}

	fmt.Printf("%064b\n%[1]d\n", uint64(number))
}
