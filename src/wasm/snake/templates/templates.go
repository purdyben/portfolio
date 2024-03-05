package templates

import _ "embed"

//go:embed board.html
var Boardhtml string

//go:embed row.html
var Rowhtml string

//go:embed square.html
var Squarehtml string


<div id="gridContainer"></div>

// <script>
// // Number of rows and columns
// const rows = 3; // Adjust the number of rows as needed
// const columns = 4; // Adjust the number of columns as needed

// // Create the grid cells
// const gridContainer = document.getElementById('gridContainer');
// for (let i = 0; i < rows * columns; i++) {
//   const cell = document.createElement('div');
//   cell.className = 'grid-cell';
//   cell.addEventListener('click', toggleCellColor);
//   gridContainer.appendChild(cell);
// }

