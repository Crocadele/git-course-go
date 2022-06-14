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
	fmt.Println("Гроші які потрібно витратити, щоб купити 9 яблук та 8 груш:", firstPrice)

	pearMaxCount := float64(amount) / pearPrice
	fmt.Println("Кількість груш що ми можемо купити: ", int(pearMaxCount))

	appleMaxCount := float64(amount) / applePrice
	fmt.Println("Кількість яблук що ми можемо купити: ", int(appleMaxCount))

	checkPrice := 2.0*applePrice + 2.0*pearPrice
	fmt.Println("Ми не можемо купити 2 груші та 2 яблука, так як сума грошей щонам потрібна стоновить:", checkPrice)

}
