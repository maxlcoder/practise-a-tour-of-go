package main

import "fmt"

func main()  {
	i := 1
	p := &i
	i++
	fmt.Println(*p)

	a := *p
	i++
	fmt.Println(a)
}