package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println(time.Saturday)

	today := time.Now().Weekday()

	fmt.Println(today)

	switch time.Friday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
		fmt.Println(today + 1)
	default:
		fmt.Println("不确定")

	}

	a := 1
	switch {
	case a == 1:
		fmt.Println("a is 1")
	case a == 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is nothing")
	}

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}