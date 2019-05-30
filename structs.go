package main

import (
	"fmt"
)

type Vertex struct {
	X int
	y int
}

func main()  {
	a := Vertex{1,2}
	fmt.Println(Vertex{1,2})
	fmt.Println(a.X)
	fmt.Println(a.y)
	a.X = 3
	fmt.Println(a)

	p := &a
	fmt.Println(p.X)
	fmt.Println(p)
	fmt.Println(*p)
}

