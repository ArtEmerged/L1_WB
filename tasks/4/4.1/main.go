package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
	Программа реализует непрерывную запись данных в канал (основной поток) и набор из N воркеров, 
	которые читают произвольные данные из канала и выводят их в стандартный вывод. 
	Пользователю предоставляется возможность выбрать количество воркеров при старте программы.
*/

func main() {
	var N int
	fmt.Println("укажите количество воркеров: 0 < N < 100")
	fmt.Scan(&N)

	// Проверяем валидность введенных данных о количестве воркеров
	if N < 1 || N > 99 {
		fmt.Printf("вы указали %d воркеров. количество воркеров: 0 < N < 100 ", N)
		return
	}

	// Создаем канал для передачи данных от основного потока к воркерам
	ch := make(chan int)

	// WaitGroup для ожидания завершения всех горутин.
	// Можно обойтись без WaitGroup,
	// использую его, чтобы горутина успела завершить работу до завершения работы main()
	var wg sync.WaitGroup

	// Создаем контекст для уведомления воркеров о завершении работы
	quit, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// 
	defer stop()

	// Запускаем воркеры
	for i := 0; i < N; i++ {
		// Добавляем в WaitGroup 1 для каждого воркера
		wg.Add(1)
		go worker(quit, &wg, ch, i+1)
	}

	log.Println("sender запущен")

	// 	Запускаем бесконечный цикл для непрерывно записывает данные в канал
	for {
		// select будет рандомно выбирать тот case, который не блокирует горутину при чтении из канала или записи.
		// Если таких case нет, то выполняется default.
		select {
		case <-quit.Done():
			// После получения сигнала завершения работы основного потока (Ctrl + C),
			//то закрываем канал и ожидаем завершения работы воркеров
			close(ch)
			log.Println("sender завершился")
			wg.Wait()
			return
		default:
			// Слип был добавлен для удобства чтения вывода в stdout 
			time.Sleep(time.Second)
			// Отправляем радомные значения в канал 
			numSend := rand.Intn(100)
			ch <- numSend
			log.Printf("sender передал: %d\n", numSend)
		}
	}
}

// worker является функцией-воркером, которая читает данные из канала и выводит их
func worker(ctx context.Context, wg *sync.WaitGroup, ch <-chan int, numW int) {
	// Выполняем отложенный вызов для уменьшения счетчика WaitGroup на 1
	defer wg.Done()
	log.Printf("worker[%d] запущен\n", numW)
	// Запускаем бесконечный цикл
	for {
		select {
		case <-ctx.Done():
			// По завершению работы контекста воркер завершает работу
			log.Printf("worker[%d] завершился\n", numW)
			return
		default:
			// Получаем данные из канала и выводим из в stdout 
			// Если ok == false то канал закрыт 
			out, ok := <-ch
			if ok {
				fmt.Printf("worker[%d] выводит %d\n", numW, out)
			}
		}
	}
}
