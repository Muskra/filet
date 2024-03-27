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
	Position          Coordinates
	Value                int
	State                bool
	ValidatedLinkedCells []int
    IsIn                 bool
}

/* the actual grid */
type Grid struct {
	State [][]Cell
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

/* this is the "main" function of this library, it's being used to process generations grom a grid to another generation */
func Catch(grid [][]Cell, rules Rules) ([][]Cell, error) {

	for i, line := range grid {
		for j := range line {
			var err error = nil
			grid[i][j].State, err = actualCellState(Coordinates{X: i, Y: j}, rules, grid)
			if err != nil {
				return nil, fmt.Errorf("< ! > Filet, CatchError: %s\n", err)
			}
            
            realTargetLocations, err := findRealTargetLocations(rules.TargetCellsLocations, Coordinates{X: len(grid), Y: len(grid[0])})
			if err != nil {
                return [][]Cell{}, fmt.Errorf("< ! > Filet, CatchError: %s\n", err)
            }
            switch grid[i][j].State {

			case true:
                err := nextGeneration(grid, Coordinates{X: i, Y: j}, realTargetLocations, rules.TargetValues.TargetIfAlive, rules.RuleSet[:])
				if err != nil {
					return nil, fmt.Errorf("< ! > Filet, CatchError: %s\n", err)
				}
			case false:
                err := nextGeneration(grid, Coordinates{X: i, Y: j}, realTargetLocations, rules.TargetValues.TargetIfDead, rules.RuleSet[:])
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
		return false, nil//fmt.Errorf("isDeadOrAlive() -> a cell can't be both alive and dead, got alive -> %t and dead -> %t", alive, dead)
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

func findRealTargetLocations(targetAddresses []Coordinates, xyLen Coordinates) ([]Coordinates, error) {

    ret := make([]Coordinates, 0)

    for _, coordinate := range targetAddresses {

        tempValueOfX, err := findRealLocation(coordinate.X, xyLen.X)
        if err != nil {
            return []Coordinates{}, fmt.Errorf("findRealTargetLocations() -> %s", err)
        }

        tempValueOfY, err := findRealLocation(coordinate.Y, xyLen.Y)
        if err != nil {
            return []Coordinates{}, fmt.Errorf("findRealTargetLocations() -> %s", err)
        }
        ret = append(ret, Coordinates{X: tempValueOfX, Y: tempValueOfY})
    }
    return ret, nil

}

func findRealLocation(target int, sizeLimit int) (int, error) {

    if target < 0 {
        return negativeLoopThrougth(target, sizeLimit), nil
    
    } else if target >= sizeLimit-1 {
        return positiveLoopThrougth(target, sizeLimit-1), nil

    } else if target >= 0 && target < sizeLimit {
        return target, nil

    } else {
        return 0, fmt.Errorf("findRealLocation() -> Error when searching for new position, got 'target = %d', 'sizeLimit = %d'", target, sizeLimit)
    }
}

func negativeLoopThrougth(target int, sizeLimit int) int {
    temp := sizeLimit
    for i := target ; i <= 0; i = i + 1 {
        if temp == 0 {
            temp = sizeLimit
        }
        temp = temp - 1
    }
    return temp
}

func positiveLoopThrougth(target int, sizeLimit int) int {
    temp := 0
    for i := target ; i >= sizeLimit; i = i - 1 {
        if temp == sizeLimit {
            temp = 0
        }
        temp = temp + 1
    }
    return temp
}

/* this function processes the next generation of a cell */
func nextGeneration(grid [][]Cell, cellPosition Coordinates, targetLocations []Coordinates, targetValues []int, ruleSet []Set) error {

    //fmt.Println("\n\n", targetLocations)
	for _, targetPosition := range targetLocations {
        cell := grid[cellPosition.X][cellPosition.Y]
        target := grid[targetPosition.X][targetPosition.Y]
        /* check if the cell and the target are in the targetted values respectively */
        cell.IsIn = isTargetIn(cell.Value, targetValues[:])
        target.IsIn = isTargetIn(target.Value, targetValues[:])

        index, ok := ruleCheck(ruleSet, cell, target)
        if ok == true {
            err := processRule(ruleSet[index].Opcode, grid, cellPosition, targetPosition)
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
    grid[cellPosition.X][cellPosition.Y].Position.X, grid[targetPosition.X][targetPosition.Y].Position.X = grid[targetPosition.X][targetPosition.Y].Position.X, grid[cellPosition.X][cellPosition.Y].Position.X 
    grid[cellPosition.X][cellPosition.Y].Position.Y, grid[targetPosition.Y][targetPosition.Y].Position.Y = grid[targetPosition.Y][targetPosition.Y].Position.Y,  grid[cellPosition.X][cellPosition.Y].Position.Y
}


/* this function kills a cell value, channging it's state and value both to zero/false */
func killCell(grid [][]Cell, cellPosition Coordinates) {
    // to work as a game of life, this maybe required
	//grid[cellPosition.X][cellPosition.Y].Value = 0
	grid[cellPosition.X][cellPosition.Y].State = false
}

/* this function is an opcode used to opy a cell value to another */
func copyCellValue(grid [][]Cell, cellPosition Coordinates, otherCell Coordinates) {
    /* saving old coordinates of the overwritten cell */
    temp := grid[cellPosition.X][cellPosition.Y].Position
	/* copy */
    grid[cellPosition.X][cellPosition.Y] = grid[otherCell.X][otherCell.Y]
    /* re-assign the old coordinates */
    grid[cellPosition.X][cellPosition.Y].Position = temp
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
            killCell(grid, cellPosition)
        case 1:
            swapCell(grid, cellPosition, targetPosition)
        case 2:
            copyCellValue(grid, cellPosition, targetPosition)
        default:
            return fmt.Errorf("processRule() -> opcode with value '%d' isn't an instruction", opcode)
    }
    return fmt.Errorf("processRule() -> unexpected behavior encountered")
}
