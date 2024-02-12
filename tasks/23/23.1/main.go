package main

import "fmt"

/*
	Удаление i-го элемента из среза.
*/

func main() {
	newSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = deleteElem(newSlice, 1)

	fmt.Println(newSlice)

	// Здесь мы можем увидеть, что длина среза уменьшилась на один, а вместимость осталась прежней.
	fmt.Println(len(newSlice), cap(newSlice))
}

func deleteElem(old []int, i int) []int {
	// Проверяем i, чтобы избежать выхода за границы диапазона.
	if i < 0 || i > len(old)-1 {
		return old
	}
	/*
		Возвращаем новый срез с помощью append.
		Мы передаем срез от начала до элемента с индексом i (не включая его)
		и добавляем элементы, следующие за i (i+1).
		Важно понимать, что при этой операции мы не выходим за границы вместимости (cap)
		и новый срез использует тот же блок памяти, что и старый срез.
		Это можно увидеть при вызове fmt.Println(cap(newSlice)) в функции main: вместимость не изменилась,
		а изменились только элементы в срезе и его длина (len).
	*/
	return append(old[:i], old[i+1:]...)
}
