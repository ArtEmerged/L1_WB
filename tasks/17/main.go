package main

import "fmt"

// !Важно: выполняется поиск только по отсортированому массиву
func main() {
	slice := []int{1, 2, 3, 7, 10, 90, 99, 100}
	ind := binarySearch(slice, 8)
	fmt.Println(ind)

}
func binarySearch(s []int, find int) int {
	return search(s, find, 0, len(s)-1)
}

func search(s []int, find, low, higt int) int {
	if low <= higt {
		mid := (low + higt) / 2
		if s[mid] == find {
			return mid
		}
		if find > s[mid] {
			return search(s, find, mid+1, higt)
		}
		return search(s, find, low, mid-1)
	}
	return -1
}
