package main 

import (
	"fmt"
)

func main(){
	var arr1 = [3]int{1,2,3}
	arr2 := [...]int{4,5,6,7,8,9,10}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var cars = [4]string{"Volvo", "BMW", "Mercedes", "Audi"}
	fmt.Println(cars[0])
	fmt.Println(cars[1])
	fmt.Println(cars[2])
	fmt.Println(cars[3])
}