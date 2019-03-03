// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(2)

// 	// start the routines
// 	go addTable()
// 	go multiTable()

// 	// wait for the job to end
// 	fmt.Println("Waiting")
// 	wg.Wait()

// 	fmt.Println("Everything Done!")
// }

// func addTable() {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		sleep := rand.Int63n(1000)
// 		time.Sleep(time.Duration(sleep) * time.Millisecond)
// 		fmt.Println("Additon Time to Table for: ", i)
// 		fmt.Println("Addition Table for:", i)
// 		for j := 1; j <= 10; j++ {
// 			fmt.Printf("%d+%d=%d\t", i, j, i+j)
// 		}
// 		fmt.Print("\n")
// 	}
// }

// func multiTable() {
// 	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
// 	defer wg.Done()
// 	for i := 1; i <= 10; i++ {
// 		sleep := rand.Int63n(1000)
// 		time.Sleep(time.Duration(sleep) * time.Millisecond)
// 		fmt.Println("Multiplication Table for:", i)
// 		for j := 1; j <= 10; j++ {
// 			fmt.Printf("%d*%d=%d\t", i, j, i*j)
// 		}
// 		fmt.Print("\n")
// 	}
// }
