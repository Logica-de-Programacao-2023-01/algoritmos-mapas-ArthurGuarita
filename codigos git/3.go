package main

import "fmt"

func main() {
	c := map[string]int{
		"num1": 4,
		"num2": 5,
		"num3": 8,
		"num4": 2,
	}
	fmt.Println(somaMap(c))
}

func somaMap(c map[string]int) int {
	soma := 0
	for _, valor := range c {
		soma += valor
	}
	return soma
}
