// package main

// import "fmt"

// func main(){
// 	// fn := func (msg string)(u int){
// 	// 	u = 4
// 	// 	return
// 	// }

// 	//unbuffered
// 	counter := make(chan int)

// 	myMap := make(chan map[string]string)
// 	// variableMap = make(map[string]string)
// 	// append(variableMap, "hello":"shit")
// 	// myMap <- map[string]string{"hello":"shit"}

// 	//buffered channel
// 	num := make(chan int, 3)

// 	go func (){
// 		counter <- 3
// 		myMap <- map[string]string{"hello":"shit"}
// 		close(counter)
// 		close(myMap)
// 	}()

// 	go func (){
// 		num <- 10
// 		num <- 30
// 		num <- 50
// 	}()

// 	// read from the channel
// 	fmt.Println(<-counter)
	
// 	for k, v := range <-myMap{
// 		fmt.Printf("The Key is %s, And the Value is %s\n", k, v)
// 	}

// 	if value, ok := <-counter; ok{
// 		fmt.Println(value)
// 	}

// 	fmt.Println(<-num)
// 	fmt.Println(<-num)
// 	fmt.Println(<-num)
// 	close(num)
// }

// func jstTest(fn func(msg string)(u int)){}