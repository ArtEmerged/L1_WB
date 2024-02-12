package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

// senderWithTimeout отправляет случайные значения в канал ch в течение n секунд
// По истечению времени программа завершает отправку и закрывает канал
func senderWithTimeout(n int) <-chan int {

	// Инициализируем канал для отправки значений
	ch := make(chan int)

	// Создаем таймер с заданным временем
	timeout := time.After(time.Duration(n) * time.Second)


	// Вызываем анонимную функцию в отдельной горутине
	go func() {
		log.Println("sender запущен")
		// Запускаем бесконечный цикл
		for {
			// select будет рандомно выбирать тот case, который не блокирует горутину при чтении из канала или записи.
			// Если таких case нет, то выполняется default
			select {
			// Отправляем случайное значение в канал
			case ch <- rand.Intn(100):
			// Если таймер сработал, завершаем работу и закрываем канал
			case <-timeout:
				log.Println("senderWithTimeout завершился")
				close(ch)
				return
			}
		}

	}()
	return ch
}

func main() {
	fmt.Println("укажите время выполнения программы в секундах")
	var n int
	fmt.Scan(&n)

	// Валидируем полученное значение
	if n < 1 {
		fmt.Println("вы указали время работы программы меньше 1 секунды")
		return
	}

	// Получаем канал с отправляемыми значениями
	out := senderWithTimeout(n)

	log.Println("worker запущен")

	// Читаем значения из канала и выводим их на экран
	for value := range out {
		fmt.Println(value)
	}
	log.Println("worker завершился")

}
