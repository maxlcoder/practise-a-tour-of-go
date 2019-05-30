package main

import "fmt"

type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func main()  {
	addrs := map[string]IPAddr{
		"look": {127, 0, 0, 1},
		"dns": {0, 0, 0, 0},
	}

	for _,i := range addrs {
		fmt.Println(i)
	}
}
