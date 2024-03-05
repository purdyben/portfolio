package snake

import (
	"fmt"
)

type Snake struct {
	body           *node
	length         int
	rowVelocity    int
	columnVelocity int
}

type node struct {
	Point point
	Next  *node
}
type point struct {
	X, Y int
}

func NewSnake(start point) *Snake {
	return &Snake{
		body:           &node{Point: start},
		length:         1,
		rowVelocity:    1,
		columnVelocity: 1,
	}
}
func (s *Snake) Head() *node {
	return s.body
}
func (s *Snake) Length() int {
	return s.length
}

func tick() {
	fmt.Println("hello")
}

// snake’s rowVelocity to -1 and columnVelocity to 0
// - If user input if down arrow key and snake is moving horizontally then set snake’s rowVelocity to 1 and columnVelocity to 0
// - If user input if left arrow key and snake is moving vertically then set snake’s rowVelocity to 0 and columnVelocity to -1
// - If user input if right arrow key and snake is moving vertically then set snake’s rowVelocity to 0 and columnVelocity to 1
