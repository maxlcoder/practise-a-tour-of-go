package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	fmt.Println(sum)
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		res := a
		a = b
		b = res + b
		return res
	}
}

func main()  {
	pos, neg := adder(), adder()


	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}