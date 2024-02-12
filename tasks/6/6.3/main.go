package main

import (
	"context"
	"log"
	"sync"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	// Создаем контекст с таймаутом выполнения 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// WaitGroup для ожидания завершения всех горутин.
	// Можно обойтись без WaitGroup,
	// использую его, чтобы горутина успела завершить работу до завершения работы main().
	var wg sync.WaitGroup

	// Добавляем в WaitGroup 1 для синхронизации
	wg.Add(1)

	// Запуск горутины-воркера
	go worker(ctx, &wg)

	// Ожидание завершения работы горутины
	wg.Wait()

}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	// Выполняем отложенный вызов для уменьшения счетчика WaitGroup на 1
	defer wg.Done()

	log.Println("начинаю работу")

	// Запускаем бесконечный цикл
	for {
		// select будет рандомно выбирать тот case, который не блокирует горутину при чтении из канала или записи.
		// Если таких case нет, то выполняется default
		select {
		// Если контекст отменен, завершаем работу горутины (прошло 5 секунд)
		case <-ctx.Done():
			log.Println("заканчиваю работу")
			return
		default:
			//work ...
			time.Sleep(time.Second)
		}

	}
}
