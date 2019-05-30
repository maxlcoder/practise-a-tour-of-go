package main

import "fmt"

func fibonaccis(n int, c chan int)  {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		if i == 5 {
			close(c)
		}
	}

}

func main()  {
	c := make(chan int, 10)
	//go fibonaccis(cap(c), c)
	c <- 1
	for i := range c {
		fmt.Println(i)
	}
}
