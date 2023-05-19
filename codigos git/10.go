package main

import "fmt"

func main() {
	n := []int{1, 2, 2, 3, 4, 66, 66, 34, 3, 22, 89}
	r := parMap(n)
	for par, frequencia := range r {
		fmt.Printf("Par: %d - Frequência: %d\n", par, frequencia)
	}
}

func parMap(n []int) map[int]int {
	mapaPar := make(map[int]int)

	for i := 0; i < len(n)-1; i++ {
		if n[i]%2 == 0 {
			for j := i + 1; j < len(n); j++ {
				if n[j]%2 == 0 {
					par := n[i]*10 + n[j]
					mapaPar[par]++
				}
			}
		}
	}

	return mapaPar
}
