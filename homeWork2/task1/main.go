package main

import (
	"fmt"
)

func main() {

	arr := []int{4, 1, 4, -4, 3, 6, 8, 8}
	var result []int
	checkMap := make(map[int]int, 0)

	for i, val := range arr {

		if _, ok := checkMap[val]; !ok {
			checkMap[val] = i
			result = append(result, val)
		}
	}
	fmt.Println(result)
}
