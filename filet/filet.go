package filet

import (
	"fmt"
)

const (
    SWAP_CELLS uint8 = iota
    KILL_CELL
    COPY_VALUE
)

/* this structure is being used to define the actual rules and conditions for the automata */
type Set struct {
    CellValue        int
    CellState        bool
    TargetValue      int
    TargetState      bool
    ShouldBeTargeted bool
    Opcode           uint8
}

type Coordinates struct {
	X int
	Y int
}

/* those are values used in some conditional statements, like if you want to target only some specific values, they are in those slices. For the alive/dead values, those are being used to define the states of the values that are alive or dead. The targetIf ones are being used to decide wether a value can be a target from alive/dead states. */
type Target struct {
	AliveValues   []int
	DeadValues    []int
	TargetIfAlive []int
	TargetIfDead  []int
}

/* Rules are assembling some of other structs, where Target and Set are being used. Another slice is used here to define where all the targets will be at in the grid, if you decide to go out of bounds, it will loop througth the sides of the grid to prevent errors. */
type Rules struct {
    RuleSet              []Set
	TargetCellsLocations []Coordinates
	TargetValues         Target
}

/* define a cell, by default, isIn should be set to false, the program uses it to keep track of some conditional events. The validatedLinkedCells, is currently not used in the library, it can be used to check wether a cell have neighbors (with an opcode for example). the other elements are self explanatory. */
type Cell struct {
	Coordinates          Coordinates
	Value                int
	State                bool
	ValidatedLinkedCells []int
    IsIn                 bool
}

/* the actual grid */
type Grid struct {
	State [][]Cell
}

/* this is the "main" function of this library, it's being used to process generations grom a grid to another generation */
func Catch(grid [][]Cell, rules Rules) ([][]Cell, error) {

	for i, line := range grid {
		for j := range line {
			var err error = nil
			grid[i][j].State, err = actualCellState(Coordinates{X: i, Y: j}, rules, grid)
			if err != nil {
				return nil, fmt.Errorf("< ! > Filet, CatchError: %s\n", err)
			}
            realTargetLocations := findRealTargetLocations(Coordinates{X: grid[i][j].Coordinates.X, Y: grid[i][j].Coordinates.Y}, rules.TargetCellsLocations, Coordinates{X: len(grid), Y: len(grid[0])})
			switch grid[i][j].State {

			case true:
                err := nextGeneration(grid[:], Coordinates{X: i, Y: j}, realTargetLocations, rules.TargetValues.TargetIfAlive, rules.RuleSet[:])
				if err != nil {
					return nil, fmt.Errorf("< ! > Filet, CatchError: %s\n", err)
				}
			case false:
                err := nextGeneration(grid[:], Coordinates{X: i, Y: j}, realTargetLocations, rules.TargetValues.TargetIfDead, rules.RuleSet[:])
				if err != nil {
					return nil, fmt.Errorf("< ! > Filet, CatchError: %s\n", err)
				}
			}
		}
	}
	return grid, nil
}

/* return the actual state of a given cell in comparison with a given rule */
func actualCellState(cellPosition Coordinates, rules Rules, grid [][]Cell) (bool, error) {
	b, err := isDeadOrAlive(isAlive(grid[cellPosition.X][cellPosition.Y].Value, rules.TargetValues.AliveValues), isDead(grid[cellPosition.X][cellPosition.Y].Value, rules.TargetValues.DeadValues))
	if err != nil {
		return false, fmt.Errorf("actualCellState() -> %s", err)
	}
	return b, nil
}

/* comparing the states between cell and rule */
func isDeadOrAlive(alive bool, dead bool) (bool, error) {
	if alive == dead {
		return false, fmt.Errorf("isCellDeadOrAlive() -> a cell can't be both alive and dead, got alive -> %t and dead -> %t", alive, dead)
	} else if (alive == true) && (dead == false) {
		return true, nil
	} else if (dead == true) && (alive == false) {
		return false, nil
	} else {
		return false, fmt.Errorf("isDeadOrAlive() -> unexpected event when checking cell state, got alive -> %t and dead -> %t", alive, dead)
	}
}

func isAlive(c int, t []int) bool {
	for i := range t {
		if c == t[i] {
			return true
		}
	}
	return false
}

func isDead(c int, t []int) bool {
	for i := range t {
		if c == t[i] {
			return true
		}
	}
	return false
}

// update this function by adding the lower version and separating the lower and greater into other functiions
func findRealTargetLocations(actualPosition Coordinates, targetAdresses []Coordinates, xyLen Coordinates) []Coordinates {

	r := make([]Coordinates, len(targetAdresses))

	for index, tCoordinates := range targetAdresses {
        
        if tCoordinates.X < 0 {
            xUnderflow := checkPositionUnderflow(actualPosition.X, tCoordinates.X)
            if xUnderflow == true {
                r[index].X = xyLen.X - (actualPosition.X - tCoordinates.X)
            } else {
                r[index].X = actualPosition.X - tCoordinates.X
            }
        }
        if tCoordinates.Y < 0 {
            yUnderflow := checkPositionUnderflow(actualPosition.Y, tCoordinates.Y)
            if yUnderflow == true {
                r[index].Y = xyLen.Y - (actualPosition.Y - tCoordinates.Y)
            } else {
                r[index].Y = actualPosition.Y - tCoordinates.Y
            }
        }
        if tCoordinates.X > 0 {
            xOverflow := checkPositionOverflow(actualPosition.X, tCoordinates.X, xyLen.X)
            if xOverflow == true {
                r[index].X = xyLen.X - tCoordinates.X
            } else {
                r[index].X = actualPosition.X + tCoordinates.X
            }
        }
        if tCoordinates.Y > 0 {
            yOverflow := checkPositionOverflow(actualPosition.Y, tCoordinates.Y, xyLen.Y)
            if yOverflow == true {
                r[index].Y = xyLen.Y - tCoordinates.Y
            } else {
                r[index].Y = actualPosition.Y + tCoordinates.Y
            }
        }
    }
    return r
}

/* here we check if the targetPosition can be reached out from cell's position */
func checkPositionUnderflow(cellPosition int, targetPosition int) bool {
    ret := false
    if (targetPosition - cellPosition) < 0 {
        ret = true
    }
    return ret
}

/* this function checks wether a cell position is out of bound of the grid */
func checkPositionOverflow(cellPosition int, targetPosition int, length int) bool {
    ret := false
    if (targetPosition + cellPosition) > length {
        ret = true
    }
    return ret

}

/* this function processes the next generation of a cell */
func nextGeneration(grid [][]Cell, cellPosition Coordinates, targetLocations []Coordinates, targetValues []int, ruleSet []Set) error {

	for _, targetPosition := range targetLocations {

        cell := grid[cellPosition.X][cellPosition.Y]
        target := grid[targetPosition.X][targetPosition.Y]
        /* check if the cell and the target are in the targetted values respectively */
        cell.IsIn = isTargetIn(cell.Value, targetValues[:])
        target.IsIn = isTargetIn(target.Value, targetValues[:])

        index, ok := ruleCheck(ruleSet, cell, target)
        if ok == true {
            err := processRule(ruleSet[index].Opcode, grid[:], cellPosition, targetPosition)
            if err != nil {
                return fmt.Errorf("nextGeneration() -> %s", err)
            }
        }
	}
	return nil
}

/* this function swap cells positions */
func swapCell(grid [][]Cell, cellPosition Coordinates, targetPosition Coordinates) {
    /* swap cell with target may be hard to read but it's what it does */
	grid[cellPosition.X][cellPosition.Y], grid[targetPosition.X][targetPosition.Y] = grid[targetPosition.X][targetPosition.Y], grid[cellPosition.X][cellPosition.Y]
    /* swap coordinates to prevent misleading values */
    grid[cellPosition.X][cellPosition.Y].Coordinates.X, grid[targetPosition.X][targetPosition.Y].Coordinates.X = grid[targetPosition.X][targetPosition.Y].Coordinates.X, grid[cellPosition.X][cellPosition.Y].Coordinates.X 
    grid[cellPosition.X][cellPosition.Y].Coordinates.Y, grid[targetPosition.Y][targetPosition.Y].Coordinates.Y = grid[targetPosition.Y][targetPosition.Y].Coordinates.Y,  grid[cellPosition.X][cellPosition.Y].Coordinates.Y
}


/* this function kills a cell value, channging it's state and value both to zero/false */
func killCell(grid [][]Cell, cellPosition Coordinates) {
	grid[cellPosition.X][cellPosition.Y].Value = 0
	grid[cellPosition.X][cellPosition.Y].State = false
}

/* this function is an opcode used to opy a cell value to another */
func copyCellValue(grid [][]Cell, cellPosition Coordinates, otherCell Coordinates) {
    /* saving old coordinates of the overwritten cell */
    temp := grid[cellPosition.X][cellPosition.Y].Coordinates
	/* copy */
    grid[cellPosition.X][cellPosition.Y] = grid[otherCell.X][otherCell.Y]
    /* re-assign the old coordinates */
    grid[cellPosition.X][cellPosition.Y].Coordinates = temp
}

/* search from a slice of elements if a value is in the slice, here we use it to check wether the value is in a tagert list */
func isTargetIn(target int, values []int) bool {
	for _, v := range values {
		if target == v {
			return true
		}
	}
	return false
}

/* This function checks if a rule is applicable from the context of the actual cell and it's targetted value */
func ruleCheck(ruleSet []Set, cell Cell, target Cell) (int, bool) {
    for index, rule := range ruleSet {
        
        /* checks if the Set cell and target values are identical and check if the rule should be processed from the isIn value and the shouldBeTargeted */
        isContextApplicable := (cell.Value == rule.CellValue) && (cell.State == rule.CellState) && (target.Value == rule.TargetValue) && (target.State == rule.TargetState) && (target.IsIn == rule.ShouldBeTargeted)
        
        if isContextApplicable == true {
            return index, true
        } else if isContextApplicable == false {
            return 0, false
        }
    }
    return 0, false
}

/* This function executes rules from opcodes */
func processRule(opcode uint8, grid [][]Cell, cellPosition Coordinates, targetPosition Coordinates) (error) {
    switch opcode {
        case 0:
            killCell(grid[:], cellPosition)
        case 1:
            swapCell(grid[:], cellPosition, targetPosition)
        case 2:
            copyCellValue(grid[:], cellPosition, targetPosition)
        default:
            return fmt.Errorf("processRule() -> opcode with value '%d' isn't an instruction", opcode)
    }
    return fmt.Errorf("processRule() -> unexpected behavior encountered")
}
