package main

import "fmt"

func main() {
	// Membuat map
	person := map[string]string{
		"name":    "Raka",
		"job":     "Engineer",
		"country": "Indonesia",
	}

	fmt.Println("Nama:", person["name"])
	fmt.Println("Pekerjaan:", person["job"])

	person["hobby"] = "Coding"

	person["job"] = "Software Developer"

	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

	delete(person, "country")

	fmt.Println("Setelah hapus 'country':", person)
}
