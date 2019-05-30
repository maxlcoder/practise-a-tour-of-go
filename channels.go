package main

import (
	"fmt"
)

func sum(a []int, c chan int)  {
	fmt.Println(len(a))
	sum := 0
	for _,v := range a {
		sum += v
	}
	c <- sum
}

func testM()  {
	fmt.Println("fsdf")
}

func main()  {
	go testM()
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println("开始")
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从 c 中获取
	fmt.Println(x, y, x+y)
}