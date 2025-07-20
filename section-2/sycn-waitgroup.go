package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1 selesai")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2 selesai")
	}()

	fmt.Println("Menunggu semua goroutine selesai...")
	wg.Wait()
	fmt.Println("Semua goroutine selesai.")
}
