package filet

import (
	"fmt"
)

// Opcodes Type defines a Code map with contain each Code to match on with the associated function to operate with
type Opcodes struct {
	Code map[uint8]func(grid [][]Cell, pos []Coordinates) [][]Cell
}

// ProcessRule function executes rules (as functions) from the given opcode
func (opcd Opcodes) ProcessRule(opcode uint8, grid [][]Cell, cellPosition Coordinates, targetPosition Coordinates) ([][]Cell, error) {

	coo := []Coordinates{
		cellPosition,
		targetPosition,
	}

	rule, ok := opcd.Code[opcode]

	if ok {
		newGrid := rule(grid, coo)
		return newGrid, nil
	} else {
		return [][]Cell{}, fmt.Errorf("grid.go line 26 -> opcode with value '%d' isn't an instruction", opcode)
	}
}
