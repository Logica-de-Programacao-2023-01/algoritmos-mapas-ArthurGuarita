package main

import "fmt"

//Escreva uma função que receba uma lista de mapas, onde cada mapa contém a contagem de palavras de um texto, e
//retorne um único mapa contendo a soma de todas as contagens.

func main() {
	c := []map[string]int{
		{
			"Je sui a garçon": 4,
		},
		{
			"harry potter": 2,
		},
	}
	fmt.Print(contarListaPalavras(c))
}
func contarListaPalavras(c []map[string]int) map[string]int {
	soma := 0
	for _, lista := range c {
		for _, num := range lista {
			soma += num
		}
	}
	mapa := map[string]int{
		"soma": soma,
	}

	return mapa
}