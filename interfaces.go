package main

import (
	"fmt"
)

type Abser interface {
	Abs() float64
}

func main()  {
	var a Abser
	v := &methodVertex{3, 4}

	a = v

	fmt.Println(a.Abs())

}

