package filet

import (
	"fmt"
)

/* this is the "main" function of this library, it's being used to process generations grom a grid to another generation */
func Catch(state [][]Cell, rules Rules) ([][]Cell, error) {

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
				return nil, fmt.Errorf("filet.go line 25 -> %s\n", err)
			}

			realTargetLocations, err :=  Coordinates{
                X: len(grid.State),
                Y: len(grid.State[0]),
            }.findRealTargetLocations(rules.TargetCellsLocations)

			if err != nil {
				return [][]Cell{}, fmt.Errorf("filet.go line 34 -> %s\n", err)
			}

			switch grid.State[i][j].State {

			case true:
				err := grid.NextGeneration(coo, realTargetLocations, rules.TargetValues.TargetIfAlive, rules.RuleSet[:])
				if err != nil {
					return nil, fmt.Errorf("filet.go line 42 -> %s\n", err)
				}

			case false:
				err := grid.NextGeneration(coo, realTargetLocations, rules.TargetValues.TargetIfDead, rules.RuleSet[:])
				if err != nil {
					return nil, fmt.Errorf("filet.go line 48 -> %s\n", err)
				}
			}
		}
	}
	return grid.State, nil
}
