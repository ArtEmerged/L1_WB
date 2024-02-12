package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	set1 := []string{"apple", "banana", "orange", "kiwi", "strawberry", "grapefruit", "pineapple"}
	set2 := []string{"banana", "grape", "watermelon", "apple", "pear", "peach", "orange"}

	set := setIntersection(set1, set2)

	//вывод:apple banana orange
	fmt.Println(set)
}

// Функция setIntersection возвращает пересечение двух неупорядоченных множеств
func setIntersection(s1, s2 []string) []string {
	/*
		Инициализруем map для хранения уникальных значений, в качестве значения используем пустую структуру,
		так как она занимает 0 байт
	*/
	setMap := make(map[string]struct{}, len(s1))

	// Инициализация среза для хранения уникальных элементов пересечения множеств
	setSlice := []string{}

	//	Проходимя циклом по элементов первого среза и добавляем их в map
	for _, value := range s1 {
		setMap[value] = struct{}{}
	}

	//	Проходимя циклом по элементов второго среза
	for _, value := range s2 {
		// Если элемент присутствует в map, добавляем его в срез пересечений множеств
		if _, ok := setMap[value]; ok {
			setSlice = append(setSlice, value)
		}
	}

	return setSlice
}
