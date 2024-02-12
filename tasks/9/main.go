package main

import (
	"fmt"
	"sync"
)

/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.
*/

// Функция worker создает конвейер для умножения чисел из входного канала на 2.
func worker(input chan int) <-chan int {
	output := make(chan int)

	go func() {
		// Получаем числам из входного канала (input)
		for value := range input {
			// Умножение числа на 2 и отправка результата в выходной канал (output)
			output <- value * 2
		}

		// Закрытие выходного канала после завершения работы
		close(output)
		fmt.Println("worker end")
	}()
	return output
}

// Функция writer выводит числа из канала в stdout.
func writer(output <-chan int, wg *sync.WaitGroup) {
	// Выполняем отложенный вызов для уменьшаем счетчик WaitGroup на 1
	defer wg.Done()

	// Получаем результат вычислений из канала output
	for result := range output {

		// Вывод числа в stdout
		fmt.Println(result)
	}
	fmt.Println("writer end")
}

func main() {
	// Инициализируем массив чисел
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 12, 14, 15, 16, 100, 20, 1, 23, 41, 41, 4, 11, 1, 41, 0}

	// Создание WaitGroup для синхронизации
	var wg sync.WaitGroup

	wg.Add(1)

	// Создание входного канала
	input := make(chan int)

	// Создание конвейера и запуск воркера для умножения чисел на 2
	output := worker(input)

	// Запуск функции writer для вывода чисел из выходного канала
	// и передаем WaitGroup для синхронизации
	go writer(output, &wg)

	// Проходимся по исходному массиву и отправка их во входной канал (input)
	for _, num := range arr {
		input <- num
	}

	// Закрываем входного канала после отправки всех чисел
	close(input)

	// Ожидаем завершения работы
	wg.Wait()

	fmt.Println("sender end")
}
