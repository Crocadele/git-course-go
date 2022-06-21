package main

import (
	"fmt"
)

func main() {

	//var max, min int32
	input := "1 3 45 64 34 -45 1 1 -45"

	arrOfStrings := make([]string, 0)
	var (
		checkElement  string
		resultElement string
		max           string
		min           string
	)

	i := 0
	for i < len(input) {

		checkElement = string(input[i])

		if checkElement != " " {

			resultElement += checkElement

			if len(input)-1 == i {
				arrOfStrings = append(arrOfStrings, string(resultElement))
				break
			}

		} else {
			arrOfStrings = append(arrOfStrings, string(resultElement))
			resultElement = ""
		}

		i++
	}

	min = arrOfStrings[0]
	max = arrOfStrings[0]

	for _, val := range arrOfStrings {

		if min > val {
			min = val
		}

		if max < val {
			max = val
		}
	}

	if min == max {
		fmt.Println(min)
	} else {
		fmt.Println(max + " " + min)
	}

	//fmt.Println(string(max), string(min))
	//fmt.Println(arrOfStrings)

}
