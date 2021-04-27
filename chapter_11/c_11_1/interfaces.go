package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1

	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

func (sq Square) Area() float32 {
	return sq.side * sq.side
}