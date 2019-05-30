package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)


type mapVertex struct {
	Lat, Long float64
}

var m map[string]mapVertex
var m1 map[string]int

func WordCount(s string) map[string]int  {
	count := make(map[string]int)
	for _,v := range strings.Fields(s) {
		count[v] = count[v] + 1
	}
	return count
}


func main()  {
	wc.Test(WordCount)
	m = make(map[string]mapVertex)
	m["Bell Labs"] = mapVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	m1 = make(map[string]int)
	m1["One"] = 1
	fmt.Println(m1["One"])

	m := make(map[string]int)
	m["One"] = 1
	fmt.Println(m["One"])

	m["One"] = 2

	fmt.Println(m["One"])

	delete(m, "One")
	fmt.Println(m["One"])

	v, ok := m["One"]
	fmt.Println(v, ok)
}
