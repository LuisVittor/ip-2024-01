package main

import "fmt"

func buble(arr []int) {
	n := len(arr)
	for i:=0; i< n-1; i++{
		for j:=0; j<n-i-1; j++ {
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
}

func main() {
	var n int
	fmt.Print("Insira o tamanho do array:\n")
	fmt.Scan(&n)
	arr := make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&arr[i])
	}
    buble(arr)
	fmt.Print("O arrey organizado:\n", arr)
	}
