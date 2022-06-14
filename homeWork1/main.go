package main

import "fmt"

func main() {

	const (
		applePrice = 5.99
		pearPrice  = 7.0
		appleCount = 9.0
		pearCount  = 8.0
		amount     = 23
	)

	firstPrice := applePrice*appleCount + pearPrice*pearCount
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\n", firstPrice)

	pearMaxCount := float64(amount) / pearPrice
	fmt.Println("Скільки груш ми можемо купити?\n", int(pearMaxCount))

	appleMaxCount := float64(amount) / applePrice
	fmt.Println("Скільки яблук ми можемо купити?\n", int(appleMaxCount))

	checkPrice := 2.0*applePrice + 2.0*pearPrice

	if int(checkPrice) <= amount {
		fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?\nТак")
	} else {
		fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?\nНі")
	}
}
