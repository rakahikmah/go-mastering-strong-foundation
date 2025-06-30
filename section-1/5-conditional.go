package main

import "fmt"

func main() {

	var nilai int
	var grade string

	nilai = 80

	if nilai >= 80 {
		fmt.Println(" Selamat anda Anda Lulus")
	}else {
		fmt.Println("Anda belum lulus harus lebih semangat lagi belajarnya")
	}

	switch {
	case nilai >= 80:
		grade = "A"
	case nilai >= 75:
		grade = "B"
	case nilai >= 65:
		grade = "C"
	case nilai >= 50:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("Nilai kamu: %d, Grade: %s\n", nilai, grade)

}