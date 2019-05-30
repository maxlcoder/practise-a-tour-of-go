package main

import (
	"fmt"
	"reflect"
)

func main() {

	s := []int{1, 2, 4, 5, 5, 56, 7, 6, 8, 85, 23}

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)

	s[2] = 133
	fmt.Println(s)


	var z [2]int
	z[0] = 1
	fmt.Println(reflect.TypeOf(z))

	b := make([]int, 0, 4)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b)

	fmt.Println(cap(b))

	c := b[:cap(b)]

	fmt.Println(c)
	fmt.Println(cap(c))

	c[0] = 1333
	fmt.Println(c)
	c = append(c, 12)
	fmt.Println(c)
	fmt.Println(len(c), cap(c))

	for i, v := range c {
		fmt.Println(i)
		fmt.Println(v)

	}
}
