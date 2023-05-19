package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	palavras := []string{"amor", "roma", "carro", "mora", "arco"}
	mapaAnagramas := agruparAnagramas(palavras)

	for _, grupoAnagramas := range mapaAnagramas {
		fmt.Println(grupoAnagramas)
	}
}

func agruparAnagramas(palavras []string) map[string][]string {
	mapaAnagramas := make(map[string][]string)

	for _, palavra := range palavras {
		palavraOrdenada := ordenarPalavra(palavra)
		mapaAnagramas[palavraOrdenada] = append(mapaAnagramas[palavraOrdenada], palavra)
	}

	return mapaAnagramas
}

func ordenarPalavra(palavra string) string {
	palavra = strings.ToLower(palavra)
	palavraChars := strings.Split(palavra, "")
	sort.Strings(palavraChars)
	return strings.Join(palavraChars, "")
}
