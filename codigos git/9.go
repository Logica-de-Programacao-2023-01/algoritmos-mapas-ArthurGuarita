package main

import "fmt"

func main() {
	n := 7
	seq := sequenciaFibonnaci(n)
	for chave, valor := range seq {
		fmt.Println(chave, valor)
	}
}

func sequenciaFibonnaci(n int) map[int]int {
	fib := make(map[int]int)
	fib[0] = 1
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}
