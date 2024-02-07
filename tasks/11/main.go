package main

import "fmt"

func main() {
	set1 := []string{"apple", "banana", "orange", "kiwi", "strawberry", "grapefruit", "pineapple"}
	set2 := []string{"banana", "grape", "watermelon", "apple", "pear", "peach", "orange"}

	set := setIntersection(set1, set2)

	//apple banana orange
	fmt.Println(set)
}

func setIntersection(s1, s2 []string) []string {
	setMap := make(map[string]struct{}, len(s1))
	setSlice := []string{}

	for _, value := range s1 {
		setMap[value] = struct{}{}
	}

	for _, value := range s2 {
		if _, ok := setMap[value]; ok {
			setSlice = append(setSlice, value)
		}
	}

	return setSlice
}
