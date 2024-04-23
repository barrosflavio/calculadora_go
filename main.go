package main

import (
	"fmt"
)

func main() {
	so := soma(5, 7, 15, 18)
	su := subtracao(10, 80)
	mu := multiplicacao(17, 25)
	di := divisao(2, 15)
	fmt.Println(so, su, mu, di)
}

func soma(i ... int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtracao(i ... int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}

func multiplicacao(i ... int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func divisao(i ... float64) float64 {
	total := 1.0
	for _, v := range i {
		total = v / total
	}
	return total
}