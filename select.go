package main

import "fmt"

func fibo(c, quit chan int)  {
	x, y := 0, 1
	i := 0
	for {
		i++
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println(i)
		}	
	}
}

func main()  {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibo(c, quit)
}