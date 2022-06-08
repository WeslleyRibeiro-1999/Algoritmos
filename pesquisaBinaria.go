package main

import "fmt"

func PesquisaBinaria(lista []int, item int) interface{} {
	var (
		inicio = 0
		fim    = len(lista) - 1
	)

	for inicio <= fim {
		metade := (inicio + fim) / 2
		chute := lista[metade]

		if chute == item {
			return metade
		}

		if chute > item {
			fim = metade - 1
		} else {
			inicio = metade + 1
		}
	}
	return nil
}

func main() {
	minha_lista := []int{1, 3, 5, 7, 9}

	fmt.Println(PesquisaBinaria(minha_lista, 2))
}
