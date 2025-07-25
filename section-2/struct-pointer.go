package main

import (
	"errors"
	"fmt"
)

// Definisi struct
type Segitiga struct {
	Alas   float64
	Tinggi float64
}

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
	fmt.Println("-------------------------")

	// Contoh penggunaan struct dan pointer ke struct
	segitiga1 := Segitiga{Alas: 6.0, Tinggi: 4.0}
	fmt.Println("Segitiga 1 (by value):", segitiga1)
	fmt.Println("Luas segitiga 1:", hitungLuasSegitigaStruct(&segitiga1))

	// Mengubah nilai struct lewat pointer
	ubahAlas(&segitiga1, 12.0)
	fmt.Println("Segitiga 1 setelah alas diubah:", segitiga1)
	fmt.Println("Luas segitiga baru:", hitungLuasSegitigaStruct(&segitiga1))
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

func hitungLuasSegitigaStruct(s *Segitiga) float64 {
	return 0.5 * s.Alas * s.Tinggi
}

func ubahAlas(s *Segitiga, nilaiBaru float64) {
	s.Alas = nilaiBaru
}
