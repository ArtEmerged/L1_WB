package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
	Реализовать конкурентную запись данных в map.
*/

// Структура cache содержащит map данных и мьютекс для синхронизации доступа к нему.
type cache struct {
	data map[int]int
	mu   sync.Mutex
}

func main() {
	// Инициализация струкуры cache
	newCache := &cache{
		// Создаем пустой map для данных
		data: make(map[int]int),
		// мьютекс инициализуется автоматически
	}

	// Канал для завершения работы горутин
	done := make(chan struct{})

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запуск нескольких горутин-воркеров для записи данных
	for i := 0; i < 5; i++ {
		// Добавляем в WaitGroup 1 для каждого воркера
		wg.Add(1)
		// Используем номер воркера для логирования
		numW := i + 1
		// Запуск горутины-воркера
		go newCache.worker(done, &wg, numW)
	}

	// work...
	time.Sleep(time.Second * 5)

	// Закрываем канал для завершения работы всех горутин
	close(done)

	// Ожидание завершения всех горутин
	wg.Wait()

	fmt.Println(newCache.data)
}

// Метод worker представляет собой воркера, который записывает данные в map.
func (c *cache) worker(done chan struct{}, wg *sync.WaitGroup, numW int) {
	// Выполняем отложенный вызов для уменьшаем счетчик WaitGroup на 1
	defer wg.Done()

	log.Printf("worker[%d] запущен\n", numW)

	// Запускаем бесконечный цикл
	for {
		// select будет рандомно выбирать тот case который не блокирует горутину при чтении из канала или записи
		// если таких case нет то выполняется default
		select {
		// Если канал done закрыт, завершаем работу горутины
		case <-done:
			log.Printf("worker[%d] завершился\n", numW)
			return
		default:
			// Генерация случайного ключа и значения для записи в кэш
			key, value := rand.Intn(100), rand.Intn(100)
			// Блокируем мьютекс для безопасного доступа к данным
			c.mu.Lock()
			// Запиываем данные в map
			c.data[key] = value
			// Разблокировка мьютекса
			c.mu.Unlock()

			fmt.Printf("worker[%d] записал %d:%d\n", numW, key, value)
			// Пауза для имитации работы
			time.Sleep(time.Second)
		}
	}
}
