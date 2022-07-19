package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {

	var wg sync.WaitGroup

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	wg.Add(len(n))

	for i, val := range n {
		go sum(val, &wg, i+1)

	}
	wg.Wait()

}

func sum(subArr []int, wg *sync.WaitGroup, number int) {
	defer wg.Done()
	var sum = 0
	for _, val := range subArr {
		sum += val
	}
	fmt.Printf("slice %v: %v\n", number, sum)
}
