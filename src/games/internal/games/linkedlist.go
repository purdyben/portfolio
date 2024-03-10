package games

type Node struct {
	C    Coordinate
	Next *Node
}

func NewNode(x, y int) *Node {
	return &Node{
		C: NewCoordinate(x, y),
	}
}

func (n *Node) Coord() Coordinate {
	return n.C
}

func (n *Node) SetCoord(c Coordinate) {
	n.C = c
}

func (n *Node) X() int {
	return n.C.X()
}

func (n *Node) Y() int {
	return n.C.Y()
}

func Last(head *Node) *Node {
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	return curr
}
