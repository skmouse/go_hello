package main

import "fmt"

const Width, Height = 640, 480
type Cursor struct {
	X, Y int
}

func (c *Cursor) Center(){
	c.X += Width / 2
	c.Y += Height / 2
}



func main()  {
	c := new(Cursor)
	c.Center()
	fmt.Println(c.X, c.Y)
}
1
2
3
4
5
6
9
8
