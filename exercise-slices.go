package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for i := 0; i < dy; i++ {
		var col []uint8
		for j := 0; j < dx; j++ {
			col = append(col, uint8(i*j))
		}
		pic = append(pic, col)
	}
	return pic
}

func main()  {
	pic.Show(Pic)
}