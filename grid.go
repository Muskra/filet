package filet

import (
	"fmt"
    "slices"
)

type Grid struct {
	State [][]Cell
}

// ActualCellState function returns the actual state of a given cell in comparison with a given rule
func (gd Grid) ActualCellState(cellPosition Coordinates, rules Rules) (bool, error) {

	grid := gd.State

	b, err := grid[cellPosition.X][cellPosition.Y].IsDeadOrAlive(rules)
	if err != nil {
		return false, fmt.Errorf("grid.go line 18 -> %s", err)
	}

	return b, nil
}

// NextGeneration function processes the next generation of a given cell
func (grid Grid) NextGeneration(cellPosition Coordinates, targetLocations []Coordinates, targetValues []int, ruleSet []Set) (Grid, error) {

	state := grid.State
	var err error

	for _, targetPosition := range targetLocations {

		cell := state[cellPosition.X][cellPosition.Y]
		target := state[targetPosition.X][targetPosition.Y]

		// check if the cell and the target are in the targetted values respectively
		cell.IsIn = isTargetIn(cell.Value, targetValues[:])
		target.IsIn = isTargetIn(target.Value, targetValues[:])

		index, ok := ruleCheck(ruleSet, cell, target)
		if ok {
			grid.State, err = OP.ProcessRule(
				ruleSet[index].Opcode, state, cellPosition, targetPosition)
			if err != nil {
				return Grid{}, fmt.Errorf("grid.go line 45 -> %s", err)
			}
		}
	}
	return grid, nil
}

// revTwoDimSlice function simply reverse a whole 2D slice
func (grid Grid) Reverse() Grid {

    cells := grid.State

    for xIndex := range cells {
        slices.Reverse(cells[xIndex])
    }
    
    slices.Reverse(cells)
    
    return Grid{
        State: cells,
    }
}

// FormatState function appends in a string every Values of a given Grid.State Cells Type appended as positive integers. This function is more of a test tool to see if the Grid.State has been altered by Rules Type
func (grid Grid) FormatState() string {
	formatted := ""

	for _, e := range grid.State {
		for _, v := range e {

			if v.Value < 0 {
				formatted = fmt.Sprintf("%s%d", formatted, -v.Value)

			} else {
				formatted = fmt.Sprintf("%s%d", formatted, v.Value)
			}
		}
	}

	return formatted
}

// PrintState function as it's named prints all the values of a given Grid Type. The format is the same as in FormatState function
func (grid Grid) PrintState() {
	for _, e := range grid.State {
		for _, v := range e {
			if v.Value < 0 {
				fmt.Printf("%d", -v.Value)
			} else {
				fmt.Printf("%d", v.Value)
			}
		}
	}
	fmt.Println()
}

// PrintDetailedState function prints all the details about a given Grid Type State
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

// GenerateTwoDimSlice function generates a 2D slice from given size
func GenerateTwoDimSlice(lines int, cols int) [][]Cell {

	grid := make([][]Cell, 0)

	for i := 0; i < lines; i = i + 1 {
		grid = append(grid, make([]Cell, 0))

		for j := 0; j < cols; j = j + 1 {
			grid[i] = append(grid[i], Cell{})
		}
	}
	return grid
}
