package main

import "fmt"

func main() {
	alunos := map[string][]float64{
		"João":   {6.7, 8.9, 9.4},
		"Júlia":  {9.7, 6.9, 8.3},
		"Josuel": {7.2, 6.1, 8.4},
	}
	fmt.Println(mediaDeNotas(alunos))

}
func mediaDeNotas(alunos map[string][]float64) map[string]float64 {
	media := make(map[string]float64)
	//

	//
	for aluno, notas := range alunos {
		var soma float64 = 0
		for _, nota := range notas {
			soma += nota
		}
		media[aluno] = soma / float64(len(notas))
	}
	return media
}
