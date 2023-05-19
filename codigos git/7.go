package main

import (
	"fmt"
	"strings"
)

// Escreva uma função que receba uma string contendo uma frase e retorne um mapa onde as chaves são as palavras
// encontradas na frase e os valores são mapas contendo a contagem de cada letra em cada palavra.
func main() {
	str := "Hoje os franceses chegaram"
	for chave, valor := range contagemDePalavras(str) {
		fmt.Printf("Palavra: '%v' | Caracteres: %v\n", chave, valor)
	}
}

func contagemDePalavras(str string) map[string]int {
	var p []string
	p = strings.Fields(str)
	countWords := make(map[string]int)
	//
	for _, w := range p {
		countWords[w] = len(w)
	}
	return countWords
}
