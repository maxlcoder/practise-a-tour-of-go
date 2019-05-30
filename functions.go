package main

import (
	"fmt"
	"runtime"
	"time"
)


func add(x int, y int) int {
	return x + y
}

func del(x, y int) int {
	return x - y
}

 

var c, python, java bool

func main() {
	fmt.Println(time.Saturday)
	fmt.Println(runtime.GOOS)
	fmt.Println(Sqrt(3))
	fmt.Println(add(23, 3))
	fmt.Println(del(23, 12))
	var i int
	fmt.Println(i, c, python, java)
}
