package templates

import _ "embed"

//go:embed board.html
var Boardhtml string

//go:embed row.html
var Rowhtml string

//go:embed square.html
var Squarehtml string
