package main

import (
	"fmt"
	"strings"
)

func main() {
	p := "Je sui a garçon et"
	fmt.Println(contarPalavras(p))
}
func contarPalavras(p string) map[string]int {
	var words []string
	words = strings.Fields(p)
	ocorrenciaDaPalavra := make(map[string]int)
	for _, word := range words {
		ocorrenciaDaPalavra[word]++
	}
	return ocorrenciaDaPalavra
}
