package main

import "fmt"

func main() {
	var name string
	name = "Raka"

	fmt.Println(name)

	var umur int8
	umur = 20
	
	fmt.Println(umur)

	fmt.Printf("Hello, %s. You are %d years old.\n", name, umur)
	fmt.Println("Hello, " + name + ". You are " + fmt.Sprint(umur) + " years old.")
}
