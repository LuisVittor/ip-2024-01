package main

import (
	"fmt"
)

func troca(vetor []int, i, j int) {
	vetor[i], vetor[j] = vetor[j], vetor[i]
}

func trocaOpostosSeMenor(vetor []int, tamanho int) {
	for i := 0; i < tamanho/2; i++ {
		oposto := tamanho - 1 - i
		if vetor[i] < vetor[oposto] {
			troca(vetor, i, oposto)
		}
	}
}

func main() {
	var testes int

	fmt.Scan(&testes)

	for t := 0; t < testes; t++ {
		var n int

		fmt.Scan(&n)

		vetor := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&vetor[i])
		}

		trocaOpostosSeMenor(vetor, n)

		for i, val := range vetor {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}