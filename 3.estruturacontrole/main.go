package main

import "fmt"

func main() {
	idade := 18

	if idade >= 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	numeros := []int{1, 2, 3, 4, 5}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
