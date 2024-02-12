package main

/*
	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
		var justString string

		func someFunc() {
			v := createHugeString(1 << 10)
			justString = v[:100]
		}

		func main() {
			someFunc()
		}
*/

var justString string

/*
	Я выявил три проблемы:
		1. justString является срезом от v. Происходит утечка памяти, так как исходная большая
		строка останется в памяти из-за того, что justString ссылается на неё.
		2. Если в переменной v есть символы, представленные не 1 байтом, то есть вероятность потери 1 символа.
		3. Можно получить паники, если в переменной будет меньше 100 символов, так как происходит обращение за пределы массива (out of range).
*/

func createHugeString(n int) string {
	return string(make([]byte, n))
}

func someFunc() {
	v := []rune(createHugeString(1 << 10))
	// Проверка количества рун, защита от паники (out of range).
	if len(v) > 99 {
		// Создаем буфер для записи 100 символов.
		buffer := make([]rune, 100)
		// Выполняем копирование в новую область памяти
		copy(buffer, v[:100])
		// Выполняем преобразование buffer из []rune в string и присваиваем значение переменной justString.
		justString = string(buffer)
	}
}

func main() {
	someFunc()
}
