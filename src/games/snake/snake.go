package snake

import (
	"fmt"
)

type Snake struct {
	head           *node
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

func (s *Snake) AddNode(x, y int) {
	curr := s.head
	for curr.Next != nil {
		curr = curr.Next
	}
	// Now Head is the last node

	curr.Next = &node{point{x, y}, nil}

}

func NewSnake(start point) *Snake {
	return &Snake{
		head:           &node{Point: start},
		length:         1,
		rowVelocity:    1,
		columnVelocity: 1,
	}
}
func (s *Snake) Head() *node {
	return s.head
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
