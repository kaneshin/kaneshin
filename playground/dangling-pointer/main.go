package main

import "fmt"

type Rect struct {
	X float32
	Y float32
}

func NewRect(x, y float32) *Rect {
	r := Rect{X: x, Y: y}
	return &r
}

func main() {
	rp := NewRect(1.2, 3.4)
	fmt.Printf("address = %p\n", rp)
	fmt.Printf(" (x, y) = (%.2f, %.2f)\n", rp.X, rp.Y)
}
