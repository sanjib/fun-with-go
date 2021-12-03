package main

import "errors"

func main() {

}

type Point struct {
	x int
	y int
}

func (p *Point) Move(dx, dy int) {
	p.x = dx
	p.y = dy
}

type Square struct {
	center Point
	length int
}

func NewSquare(x, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, errors.New("length must be greater than zero")
	}

	return &Square{
		center: Point{x, y},
		length: length,
	}, nil
}

func (s *Square) Move(dx, dy int) {
	s.center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.length * s.length
}
