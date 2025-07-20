package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10 // mengirim nilai 10 ke channel
	}()

	hasil := <-ch
	fmt.Println("Diterima dari channel:", hasil)
}
