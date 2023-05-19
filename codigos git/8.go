package main

import "fmt"

//Escreva uma função que receba um mapa onde as chaves são nomes de pessoas e os valores são as quantias de dinheiro
//que cada pessoa gastou em uma conta compartilhada. A função deve calcular o valor que cada pessoa deve receber ou pagar para igualar as despesas.

func main() {
	n := map[string]float64{
		"arthur": 200,
		"julia":  300,
	}
	pag := despesas(n)
	for pessoa, m := range pag {
		fmt.Println(pessoa, ":", m)
	}
}

func despesas(n map[string]float64) map[string]float64 {
	desp := make(map[string]float64)
	var soma float64 = 0
	var numPessoas = len(n)
	//
	for _, valor := range n {
		soma += valor
	}
	media := soma / float64(numPessoas)

	for pessoa, num := range n {
		pagamento := media - num
		desp[pessoa] = pagamento
	}
	return desp

}
