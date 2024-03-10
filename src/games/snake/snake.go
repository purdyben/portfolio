package snake

import (
	"sync"

	"wasm/internal/games"
)

type Snake struct {
	head           *games.Node
	mu             sync.RWMutex
	length         int
	rowVelocity    int
	columnVelocity int
}

func NewSnake(x int, y int) *Snake {
	return &Snake{
		head:           games.NewNode(x, y),
		length:         1,
		rowVelocity:    0,
		columnVelocity: 0,
	}
}

func (s *Snake) AddNode(x, y int) {
	curr := last(s.head)
	curr.Next = games.NewNode(x, y)
}

func (s *Snake) Head() *games.Node {
	return s.head
}

func (s *Snake) UpdateBasedOnVelocity() {
	if s.head == nil {
		panic("head is missing")
	}
	s.mu.Lock()
	s.mu.Unlock()

	yVel := s.rowVelocity
	xVel := s.columnVelocity
	y := s.head.Y()
	x := s.head.X()

	if y+yVel < 0 || y+yVel > game.Rows() {
		yVel = 0
	}
	if x+xVel < 0 || x+xVel > game.Cols() {
		yVel = 0
	}
	n := games.NewNode(x+xVel, y+yVel)

	updateSnake(s, n)
}

func updateSnake(s *Snake, newHead *games.Node) {
	if s.head == nil {
		s.head = newHead
		return
	}
	// Add New Node
	newHead.Next = s.Head()
	s.head = newHead

	var prev *games.Node
	curr := newHead
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}
	if prev != nil {
		prev.Next = nil
	}
}

func (s *Snake) SetHead(newHead *games.Node) {
	newHead.Next = s.head.Next
	s.head = newHead
}

func (s *Snake) Velocity() (int, int) {
	return s.rowVelocity, s.columnVelocity
}

func (s *Snake) SetVelocity(r, c int) {
	s.rowVelocity = r
	s.columnVelocity = c
}

func (s *Snake) SetRowVelocity(i int) {
	s.rowVelocity = i
}

func (s *Snake) SetColumnVelocity(i int) {
	s.rowVelocity = i
}

func last(head *games.Node) *games.Node {
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	return curr
}

// snake’s rowVelocity to -1 and columnVelocity to 0
// - If user input if down arrow key and snake is moving horizontally then set snake’s rowVelocity to 1 and columnVelocity to 0
// - If user input if left arrow key and snake is moving vertically then set snake’s rowVelocity to 0 and columnVelocity to -1
// - If user input if right arrow key and snake is moving vertically then set snake’s rowVelocity to 0 and columnVelocity to 1
