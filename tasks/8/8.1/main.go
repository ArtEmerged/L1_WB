package main

import "fmt"

func setBit(number, i int64) int64 {
	var bitPosition int64 = 1 << i
	return number | bitPosition
}

func resetBit(number, i int64) int64 {
	var bitPosition int64 = 1 << i
	return number &^ bitPosition

}

func main() {
	fmt.Printf("Введите через пробел чисто(int64) и индек бита(i) от 0 до 63 и цифру (0 или 1)\nПример:42 53 1\nВвод:")
	var number, index, bit int64
	_, err := fmt.Scanf("%d %d %d", &number, &index, &bit)
	if err != nil {
		fmt.Println("Невалидные данные")
		return
	}

	if index > 63 || index < 0 {
		fmt.Println("Невалидный индекс")
		return
	}

	fmt.Printf("%064b\n", uint64(number))

	if bit == 1 {
		number = setBit(number, index)
	}
	if bit == 0 {
		number = resetBit(number, index)
	}
	fmt.Printf("%064b\n%[1]d\n ", uint64(number))

}
