package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	group := map[int][]float64{}

	for _, tem := range arr {
		key := int(tem) / 10 * 10
		group[key] = append(group[key], tem)
	}
	fmt.Println(group)
}
