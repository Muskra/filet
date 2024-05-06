package filet

import (
    "fmt"
)

// Opcodes defines a Code map with contain each Code to match on with the associated function to operate with
type Opcodes struct {
	Code map[uint8]func(grid [][]Cell, pos []Coordinates)
}

// This function executes rules (as functions) from opcodes
func (opcd Opcodes) ProcessRule(opcode uint8, grid [][]Cell, cellPosition Coordinates, targetPosition Coordinates) (Grid, error) {

    coo := []Coordinates{
        cellPosition,
        targetPosition,
    }

    rule, ok := opcd.Code[opcode]

	if ok {
        return rule(grid, coo), nil
    } else {
		return Grid{}, fmt.Errorf("grid.go line 25 -> opcode with value '%d' isn't an instruction", opcode)
	}
}
