package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main()  {
	a := &Person{"zhagnsan", 18}
	b := &Person{"lisi", 20}
	a.Name = "wangwu"

	fmt.Println(a, b)
}