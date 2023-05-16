package main

func main() {
	c := []map[string]int{
		{
			"Je sui a garçon": 4,
		},
		{
			"harry potter": 2,
		},
	}
}
func contarListaPalavras(c []map[string]int) map[string]int {
	soma := 0
	for _, lista := range c {
		for _, num := range lista {
			soma += num
		}
	}
	return
}