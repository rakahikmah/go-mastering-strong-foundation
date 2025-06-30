package main

import "fmt"


func main() {

	angka := [5]int{10,20,30,40,50}
	slice1 := angka[1:4]

	fmt.Println("isi slice1: ", slice1)
	fmt.Println("length slice1 : " , len(slice1))
	fmt.Println("Capacity slice1:", cap(slice1))


		
	slice2 := []string{"apel", "pisang", "mangga"}
	fmt.Println("Isi slice2:", slice2)

	
	slice2 = append(slice2, "jeruk")
	fmt.Println("Setelah ditambah:", slice2)

	
	slice3 := make([]string, len(slice2))
	copy(slice3, slice2)
	fmt.Println("Slice hasil copy:", slice3)

}