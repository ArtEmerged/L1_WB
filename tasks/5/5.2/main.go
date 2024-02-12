package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	fmt.Println("укажите время выполнения программы в секундах")

	var t time.Duration
	fmt.Scan(&t)

	// Переводим в секунды
	t *= time.Second

	// Валидируем полученное значение
	if t < time.Second {
		fmt.Println("вы указали время работы программы меньше 1 секунды")
		return
	}

	// Создаем контекст с таймаутом
	ctx, stop := context.WithTimeout(context.Background(), t)
	defer stop()

	// Создаем канал для передачи данных от горутин-отправителей к горутине-воркеру
	ch := make(chan int)
	var wg sync.WaitGroup

	// Добавляем 1 воркер в WaitGroup для синхронизации
	wg.Add(1)

	// Запускаем горутину-воркера
	go worker(ctx, &wg, ch)

	log.Println("sender запущен")

	// Отправляем значения в канал, пока не истечет время работы
	for {
		select {
		// select будет рандомно выбирать тот case, который не блокирует горутину при чтении из канала или записи.
		// Если таких case нет, то выполняется default
		case <-ctx.Done():
			// После завершения работы контекста закрываем канал и ожидаем завершения воркера
			close(ch)
			log.Println("sender завершился")
			wg.Wait()
			return
		default:
			// Отправляем случайное значение в канал
			numSend := rand.Intn(100)
			ch <- numSend
			log.Printf("sender передал: %d\n", numSend)
		}
	}
}

// worker является горутиной-воркером, которая читает значения из канала и выводит их
func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
	// Выполняем отложенный вызов для уменьшения счетчика WaitGroup на 1
	defer wg.Done()
	log.Println("worker[1] запущен")
	// Запускаем бесконечный цикл
	for {
		select {
		case <-ctx.Done():
			// По завершению работы контекста воркер завершает работу
			log.Println("worker[1] завершился")
			return
		default:
			// Получаем данные и канала и выводим их в stdout
			out, ok := <-ch
			if ok {
				fmt.Printf("worker[1] выводит %d\n", out)
			}
		}
	}
}
