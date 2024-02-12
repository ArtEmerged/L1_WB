package main

import (
	"log"
	"sync"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	// Канал для завершения работы горутин
	done := make(chan struct{})

	// WaitGroup для ожидания завершения всех горутин.
	// Можно обойтись без WaitGroup,
	// использую его, чтобы горутина успела завершить работу до завершения работы main().
	var wg sync.WaitGroup

	// Добавляем в WaitGroup 1 для синхронизации
	wg.Add(1)
	// Запуск горутины-воркера.
	go worker(done, &wg)

	//work ...
	time.Sleep(time.Second * 5)

	// Закрываем канал для завершения работы горутин.
	close(done)

	// Ожидание завершения работы горутины.
	wg.Wait()

}

func worker(done chan struct{}, wg *sync.WaitGroup) {
	// Выполняем отложенный вызов для уменьшения счетчика WaitGroup на 1.
	defer wg.Done()
	log.Println("начинаю работу")

	// Запускаем бесконечный цикл.
	for {
		// select будет рандомно выбирать тот case, который не блокирует горутину при чтении из канала или записи.
		// Если таких case нет, то выполняется default.
		select {
		// Если канал done закрыт, завершаем работу горутины.
		case <-done:
			log.Println("заканчиваю работу")
			return
		default:
			//work ...
			time.Sleep(time.Second)
		}

	}
}
