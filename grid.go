package filet

import (
	"fmt"
)

type Grid struct {
	State [][]Cell
}

/* return the actual state of a given cell in comparison with a given rule */
func (gd Grid) ActualCellState(cellPosition Coordinates, rules Rules) (bool, error) {

	grid := gd.State

	b, err := grid[cellPosition.X][cellPosition.Y].IsDeadOrAlive(rules)
	if err != nil {
		return false, fmt.Errorf("grid.go line 16 -> %s", err)
	}

	return b, nil
}

/* this function processes the next generation of a cell */
func (grid Grid) NextGeneration(cellPosition Coordinates, targetLocations []Coordinates, targetValues []int, ruleSet []Set) error {

	state := grid.State

	//fmt.Println("\n\n", targetLocations)
	for _, targetPosition := range targetLocations {

		cell := state[cellPosition.X][cellPosition.Y]
		target := state[targetPosition.X][targetPosition.Y]

		/* check if the cell and the target are in the targetted values respectively */
		cell.IsIn = isTargetIn(cell.Value, targetValues[:])
		target.IsIn = isTargetIn(target.Value, targetValues[:])

		index, ok := ruleCheck(ruleSet, cell, target)
		if ok {
			err := OP.ProcessRule(ruleSet[index].Opcode, state, cellPosition, targetPosition)

			if err != nil {
				return fmt.Errorf("grid.go line 38 -> %s", err)
			}
		}
	}
	return nil
}

func (grid Grid) PrintState() {
    for _, e := range grid.State {
        for _, v := range e {
            fmt.Printf("%d", v.Value)
        }
    }
    fmt.Println()
}

func (grid Grid) PrintDetailedState() {
	for _, e := range grid.State {
		for _, v := range e {
			fmt.Printf("X: %d, Y: %d, V: %d, S: %v, VLC: %v, II: %v\n",
				v.Position.X,
				v.Position.Y,
				v.Value,
				v.State,
				v.ValidatedLinkedCells,
				v.IsIn,
			)
		}
		fmt.Println()
	}
}

func GenerateTwoDimArray(lines int, cols int) [][]Cell {

	grid := make([][]Cell, 0)

	for i := 0; i < lines; i = i + 1 {
		grid = append(grid, make([]Cell, 0))

		for j := 0; j < cols; j = j + 1 {
			grid[i] = append(grid[i], Cell{})
		}
	}
	return grid
}
