package filet

import (
	"fmt"
)

/* define a cell, by default, isIn should be set to false, the program uses it to keep track of some conditional events. The validatedLinkedCells, is currently not used in the library, it can be used to check wether a cell have neighbors (with an opcode for example). the other elements are self explanatory. */
type Cell struct {
	Position             Coordinates
	Value                int
	State                bool
	ValidatedLinkedCells []int
	IsIn                 bool
}

// Equal function calculates if a given rule Set Type and a target Cell Type are equal. Returns a boolean
func (cell *Cell) Equal(rule Set, target Cell) bool {
	return (cell.Value == rule.CellValue) &&
		(cell.State == rule.CellState) &&
		(target.Value == rule.TargetValue) &&
		(target.State == rule.TargetState) &&
		(target.IsIn == rule.ShouldBeTargeted)
}

// IsDeadOrAlive function compares the states between Cell Type and Rules Type from the Alive/Dead point of view
func (cell *Cell) IsDeadOrAlive(rules Rules) (bool, error) {

	alive := IsAlive(cell.Value, rules.TargetValues.AliveValues)
	dead := IsDead(cell.Value, rules.TargetValues.DeadValues)

	if alive == dead {
        // at first i wanted that every cells can't be both dead and alive but it made the program unreliable. So i changed it to false, behavior was basically the same, everything was skipped so life prevail on death.
        return true, nil //fmt.Errorf("isDeadOrAlive() -> a cell can't be both alive and dead, got alive -> %t and dead -> %t", alive, dead)
	} else if alive && !dead {
		return true, nil
	} else if dead && !alive {
		return false, nil
	} else {
		return false, fmt.Errorf("cell.go line 37 -> unexpected event when checking cell state, got alive -> %t and dead -> %t", alive, dead)
	}
}

// IsAlive function checks if an int value is in a []int slice. Returns a boolean
func IsAlive(c int, t []int) bool {
	for i := range t {
		if c == t[i] {
			return true
		}
	}
	return false
}

// IsDead function checks if an int value is in a []int slice. Returns a boolean
func IsDead(c int, t []int) bool {
	for i := range t {
		if c == t[i] {
			return true
		}
	}
	return false
}

// PrintDetailedState function printout all the details about a 2D slice of Cell Type
func PrintDetailedState(state [][]Cell) {
	for _, e := range state {
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
