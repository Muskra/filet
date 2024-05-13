package filet

// for more convenience, the opcodes are generated automatically as constant
const (
	SWAP_CELLS uint8 = iota
	KILL_CELL
	COPY_VALUE
)

// here is where the default Codes that the program naturally provides. Those can be changed as user preference. by convention and more readability, you must add an entry in the OP value, and then call the custom made function.
var OP Opcodes = Opcodes{

	Code: map[uint8]func(grid [][]Cell, pos []Coordinates) [][]Cell{

		// this function swap cells positions
		SWAP_CELLS: func(grid [][]Cell, pos []Coordinates) [][]Cell {
			return swapCells(grid, pos)
		},

		// this function kills a cell value, channging it's state and value both to zero/false
		KILL_CELL: func(grid [][]Cell, pos []Coordinates) [][]Cell {
			return killCell(grid, pos)
		},

		// this function is an opcode used to opy a cell value to another
		COPY_VALUE: func(grid [][]Cell, pos []Coordinates) [][]Cell {
			return copyValue(grid, pos)
		},
	},
}

// swapCells function swap cells positions
func swapCells(grid [][]Cell, pos []Coordinates) [][]Cell {

	cellPosition := pos[0]
	targetPosition := pos[1]

	// swap cell with target may be hard to read but it's what it does
	grid[cellPosition.X][cellPosition.Y], grid[targetPosition.X][targetPosition.Y] = grid[targetPosition.X][targetPosition.Y], grid[cellPosition.X][cellPosition.Y]

	// swap coordinates to prevent misleading values
	grid[cellPosition.X][cellPosition.Y].Position.X, grid[targetPosition.X][targetPosition.Y].Position.X = grid[targetPosition.X][targetPosition.Y].Position.X, grid[cellPosition.X][cellPosition.Y].Position.X
	grid[cellPosition.X][cellPosition.Y].Position.Y, grid[targetPosition.Y][targetPosition.Y].Position.Y = grid[targetPosition.Y][targetPosition.Y].Position.Y, grid[cellPosition.X][cellPosition.Y].Position.Y

	return grid
}

// killCell function kills a Cell
func killCell(grid [][]Cell, pos []Coordinates) [][]Cell {

	posX := pos[0].X
	posY := pos[0].Y

	// to work as a game of life or any simulation that works with neighborhoods, this underneath value maybe reset. In spark demo we don't need to use this value
	//grid[cellPosition.X][cellPosition.Y].Value = 0
	grid[posX][posY].State = false

	return grid
}

// copyValue function copy a cell value to another
func copyValue(grid [][]Cell, pos []Coordinates) [][]Cell {

	cellPosition := pos[0]
	otherCell := pos[1]

	// saving old coordinates of the overwritten cell
	temp := grid[cellPosition.X][cellPosition.Y].Position

	// copy
	grid[cellPosition.X][cellPosition.Y] = grid[otherCell.X][otherCell.Y]

	// re-assign the old coordinates
	grid[cellPosition.X][cellPosition.Y].Position = temp

	return grid
}
