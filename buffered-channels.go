package main

import "fmt"

func inputData(i int,c chan int) {
	c <- i
}

func main()  {
	c := make(chan int, 2)
	go inputData(1, c)
	fmt.Println(<-c)
	c <- 2

	c <- 3
	fmt.Println(<-c)
	c <- 4
	fmt.Println(<-c)
	fmt.Println(<-c)
}