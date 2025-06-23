package main

import "fmt"

func tambahData(a, b int) int {
	return a + b
}

func kurangData(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(tambahData(1, 2))
	fmt.Println(kurangData(1, 2))
}
