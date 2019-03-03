package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	counter := make(chan int)
	wg.Add(2)
	go printCounter("Printer Number Two", counter)
	go printCounter("Printer Number One", counter)
	counter <- 1
	fmt.Println("Waiting for jobs to finish")
	wg.Wait()
	fmt.Println("Success")
}

func printCounter(label string, counter chan int){
	defer wg.Done()
	for value := range counter{
		fmt.Printf("Job With %s --> Value: %d\n", label, value)
		if value == 10{
			fmt.Printf("Closed With %s\n", label)
			close(counter)
			return
		}
		value++
		counter <- value
	}
}