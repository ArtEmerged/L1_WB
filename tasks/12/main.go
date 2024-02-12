package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	
	set := toSet(s)
	
	fmt.Println(set)
}

// Функция toSet преобразует последовательность строк в множество
func toSet(s []string) []string {
	/*
		Инициализруем map для хранения уникальных значений, в качестве значения используем пустую структуру,
		так как она занимает 0 байт
	*/
	setMap := map[string]struct{}{}

	// Инициализруем среза для хранения уникальных элементов
	setSlice := []string{}

	// Проходимя циклом по элементов исходного среза
	for _, value := range s {

		// Проверка, есть ли текущий элемент в множестве
		if _, ok := setMap[value]; !ok {

			// Если элемент отсутствует, добавляем его в множество и в срез
			setMap[value] = struct{}{}
			setSlice = append(setSlice, value)
		}
	}
	// Возвращаем срез с уникальными элементами
	return setSlice
}
