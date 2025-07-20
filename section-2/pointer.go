package main

import (
	"errors"
	"fmt"
)

func main() {
	tanpaPointer := 10
	var denganPointer *int = &tanpaPointer

	fmt.Println("Nilai x : ", tanpaPointer)
	fmt.Println("Alamat x : ", denganPointer)
	fmt.Println("Nilai dari pointer b : ", *denganPointer)
	fmt.Println("-------------------------")

	nilaiA := 200
	nilaiB := 300
	hasil, err := hitungPenjumlahan(&nilaiA, &nilaiB)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil perhitungan penjumlahan: ", hasil)
	}
	fmt.Println("-------------------------")

	alas := 10.0
	tinggi := 5.0
	luas, err := luasSegitiga(&alas, &tinggi)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Luas segitiga: ", luas)
	}
}

func hitungPenjumlahan(a, b *int) (int, error) {
	if a == nil || b == nil {
		return 0, errors.New("parameter tidak boleh nil")
	}
	return *a + *b, nil
}

func luasSegitiga(alas, tinggi *float64) (float64, error) {
	if alas == nil || tinggi == nil {
		return 0, errors.New("parameter tidak boleh nil")
	}
	return 0.5 * (*alas) * (*tinggi), nil
}
