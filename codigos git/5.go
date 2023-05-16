package main

import (
	"fmt"
	"strings"
)

func main() {
	c := "Abacaxi"
	fmt.Print(contagemDeCarac(c))
}
func contagemDeCarac(c string) map[string]int {
	ocorrencia := make(map[string]int)
	for _, car := range strings.ToLower(c) {
		ocorrencia[string(car)]++
	}
	return ocorrencia
}
