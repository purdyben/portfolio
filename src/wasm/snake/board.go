package snake

type Board struct {
	apple Apple
}

type Apple struct {
	Point  point
	Symbol rune
}
