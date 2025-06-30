package main

import(
	"fmt"
	"time"
)


func halo(){
	fmt.Println("Hallo dari go routinr")
}

func main() {
	go halo()
	time.Sleep(1 * time.Second)
}