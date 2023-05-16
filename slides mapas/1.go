package main

import (
	"fmt"
	"strings"
)

func main() {
	c := "Abacaxi"
	fmt.Print(contarCaracteres(c))
}

func contarCaracteres(c string) map[string]int {
	caracteres := make(map[string]int)

	for _, carac := range strings.ToLower(c) {
		caracteres[string(carac)]++
	}
	//
	return caracteres
}
