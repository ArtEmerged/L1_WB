package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}
	set := toSet(s)
	fmt.Println(set)
}

func toSet(s []string) []string {
	setMap := make(map[string]struct{})
	setSlice := []string{}
	for _, value := range s {
		if _, ok := setMap[value]; !ok {
			setMap[value] = struct{}{}
			setSlice = append(setSlice, value)
		}
	}
	return setSlice
}
