package main

import "fmt"

/*
	Реализовать быструю сортировку массива (quicksort) встроенными методами языка
*/

func main() {
	arr := []int{99994, 1, 9, 0, 7, 8, 1, 54, 34, 3, 1, 9, 88, 15, 90, 18, 22, 22, 22, 15, 4, 1, 87, 3, 454}
	quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) {
	// Выполняем сортировку, если длина среза больше 1
	if len(arr) > 1 {
		// Объявляем левый и правый индексы среза
		left, right := 0, len(arr)-1
		/*
			Определяем значение pivot. Это может быть случайное значение в нашем срезе.
			Обычно используют right, left или middle. Я привык брать middle.
		*/
		pivot := arr[(left+right)/2]
		/*
			Мы используем два указателя, left и right, которые двигаются к центру,
			пока левый меньше правого.
		*/
		for left <= right {
			// В этом цикле мы ищем значение, которое БОЛЬШЕ pivot.
			for arr[left] < pivot {
				left++
			}
			// В этом цикле мы ищем значение, которое МЕНЬШЕ pivot.
			for arr[right] > pivot {
				right--
			}
			// Если мы закончили два предыдущих цикла, то это значит, что мы нашли два значения, которые нужно поменять местами.
			// Важно поставить знак <=, так как в противном случае первый цикл будет бесконечным.
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				// Инкрементируем left и декрементируем right для перехода к следующим элементам.
				left++
				right--
			}
		}
		// Рекурсивно вызываем quickSort для левой и правой части среза.
		quickSort(arr[:right+1])
		quickSort(arr[left:])
	}
}
