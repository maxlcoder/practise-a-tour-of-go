package main

import (
	"math"
)

type methodVertex struct {
	X, Y float64
}

func (v *methodVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//
//func main()  {
//	v := &methodVertex{3, 4}
//	fmt.Println(v.Abs())
//}