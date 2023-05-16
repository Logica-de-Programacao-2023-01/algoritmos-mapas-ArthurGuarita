package main

import "fmt"

func main() {
	n := []int{5, 5, 6, 1, 3, 6, 8, 9}
	fmt.Println(valoresUnicos(n))
}

func valoresUnicos(n []int) []int {
	var u []int
	//
	contagem := make(map[int]int)
	//
	for _, num := range n {
		contagem[num]++
		if contagem[num] > 1 {
			delete(contagem, num)
		} else {
			u = append(u, num)
		}

	}

	return u
}
