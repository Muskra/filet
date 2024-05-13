package filet

import (
	"fmt"
)

// Catch function is used to process a Grid.State Type to the next generation. It's actually the 'main' function or tool that the filet library provides
func CatchOne(state [][]Cell, rules Rules) ([][]Cell, error) {

	grid := Grid{
		State: state,
	}

	for i, line := range state {
		for j := range line {

			var err error = nil
			coo := Coordinates{
				X: i,
				Y: j,
			}

			grid.State[i][j].State, err = grid.ActualCellState(coo, rules)
			if err != nil {
				return nil, fmt.Errorf("filet.go line 25 -> %s", err)
			}

			realTargetLocations, err := Coordinates{
				X: len(grid.State),
				Y: len(grid.State[0]),
			}.FindRealTargetLocations(rules.TargetCellsLocations)

			if err != nil {
				return [][]Cell{}, fmt.Errorf("filet.go line 34 -> %s", err)
			}

			switch grid.State[i][j].State {

			case true:
				grid, err = grid.NextGeneration(coo, realTargetLocations, rules.TargetValues.TargetIfAlive, rules.RuleSet[:])
				if err != nil {
					return nil, fmt.Errorf("filet.go line 42 -> %s", err)
				}

			case false:
				grid, err = grid.NextGeneration(coo, realTargetLocations, rules.TargetValues.TargetIfDead, rules.RuleSet[:])
				if err != nil {
					return nil, fmt.Errorf("filet.go line 48 -> %s", err)
				}
			}
		}
	}
	return grid.State, nil
}

// CatchNthGen function process the next 'Nth' number of generations
func CatchNthGen(grid Grid, rules Rules, lifetime int) (Grid, error) {

    var err error

	for i := 0; i < lifetime; i = i + 1 {
		grid.State, err = CatchOne(grid.State, rules)
		if err != nil {
			return Grid{}, fmt.Errorf("filet.go line 64 -> Can't process generation '%d'.", i)
		}
	}

    return grid, nil
}

