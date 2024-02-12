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
	// Инициализация флага завершения горутины.
	var done bool

	// WaitGroup для ожидания завершения всех горутин.
	// Можно обойтись без WaitGroup,
	// использую его, чтобы горутина успела завершить работу до завершения работы main().
	var wg sync.WaitGroup

	// Мьютекс для безопасного доступа к флагу завершения
	var mu sync.Mutex

	// Добавляем в WaitGroup 1 для синхронизации.
	wg.Add(1)

	// Запуск горутины-воркера.
	go worker(&done, &wg)

	//work ...
	time.Sleep(time.Second * 5)

	// Блокируем мьютекс перед изменением флага
	mu.Lock()
	done = true
	mu.Unlock()

	// Ожидание завершения работы горутины
	wg.Wait()
}

func worker(done *bool, wg *sync.WaitGroup) {
	// Выполняем отложенный вызов для уменьшения счетчика WaitGroup на 1
	defer wg.Done()

	log.Println("начинаю работу")

	// Запускаем бесконечный цикл
	for {
		// Проверяем значение флага завершения горутины
		if *done {
			log.Println("заканчиваю работу")
			return
		}
		//work ...
		time.Sleep(time.Second)
	}

}
