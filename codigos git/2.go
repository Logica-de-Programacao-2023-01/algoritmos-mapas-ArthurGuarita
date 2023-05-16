package main

import "fmt"

func main() {
	a := map[int]int{
		1: 3,
		3: 2,
		5: 1,
	}
	b := map[int]int{
		3: 6,
		6: 1,
		7: 8,
	}
	fmt.Print(juncaoDeMapas(a, b))
}
func juncaoDeMapas(a map[int]int, b map[int]int) map[int]int {
	c := make(map[int]int)
	//
	for chaveA, valorA := range a {
		c[chaveA] = valorA
	}
	for chaveB, valorB := range b {
		c[chaveB] = valorB
	}
	return c
}
