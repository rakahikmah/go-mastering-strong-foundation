package main

import (
	"errors"
	"fmt"
)

func validasiUmur(umur int) error {
	if umur < 18 {
		return errors.New("umur kurang dari 18 tahun")
	}
	
	if umur > 100 {
		return fmt.Errorf("umur lebih dari 100 tahun: %d", umur)
	}
	return nil 
}

func main() {
	fmt.Println("--- Uji Umur 17 ---")
	err1 := validasiUmur(17)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println("Validasi berhasil: Tidak ada error.")
	}

	fmt.Println("\n--- Uji Umur 25 ---")
	err2 := validasiUmur(25)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Validasi berhasil: Tidak ada error.")
	}

	fmt.Println("\n--- Uji Umur 101 ---")
	err3 := validasiUmur(101)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("Validasi berhasil: Tidak ada error.")
	}
}