package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Major string
}

func (s Student) Print() {
	fmt.Println("=== Data Mahasiswa ===")
	fmt.Println("Nama  :", s.Name)
	fmt.Println("Umur  :", s.Age)
	fmt.Println("Jurusan:", s.Major)
	fmt.Println()
}

func main() {
	// Cara 1: Inisialisasi struct satu per satu
	var student1 Student
	student1.Name = "Raka"
	student1.Age = 28
	student1.Major = "Sistem Informasi"

	// Cara 2: Inisialisasi dengan named fields
	student2 := Student{
		Name:  "Rafa",
		Age:   22,
		Major: "Arsitek",
	}

	// Cara 3: Inisialisasi dengan urutan field
	student3 := Student{"Raudhah", 22, "Public Relations"}

	// Cetak data mahasiswa
	student1.Print()
	student2.Print()
	student3.Print()
}
