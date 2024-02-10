package main

import (
	"fmt"
	"time"
)

/*
	Реализовать собственную функцию sleep.
*/

func main() {
	// Определяем текущее время для замера
	t := time.Now()

	// Вызываем функцию sleep на 2 секунды (2000 миллисекунд)
	sleep(time.Millisecond * 2000)

	// Выводим результат замера
	fmt.Println(time.Since(t))
}

func sleep(n time.Duration) {
	// Определяем текущее время
	start := time.Now()
	/*
		time.Since(start) возвращает, сколько прошло времени с момента start (time.Now.Sub(start)).
		Цикл будет продолжаться до тех пор, пока время, прошедшее с момента обновления переменной start, меньше n (переданное время для sleep)
	*/
	for time.Since(start) < n {
		// Пустой цикл ожидания
	}
}
