package main

import "fmt"

func main(){
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	fmt.Println("Perulangan selesai")


	fruits := []string{"Apple","Banana","Cherry"}

	for index, value := range fruits {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}