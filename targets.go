package filet

import (
	"fmt"
)

// Target Type contains values that are used in some conditional statements, like if you want to target only some specific values, they are in those slices. For the alive/dead values, those are being used to define the states of the values that are alive or dead. The targetIf ones are being used to decide wether a value can be a target from alive/dead states. */
type Target struct {
	AliveValues   []int
	DeadValues    []int
	TargetIfAlive []int
	TargetIfDead  []int
}

// NewTarget function return an empty Target Type
func NewTarget() Target {
	return Target{
		AliveValues:   make([]int, 0),
		DeadValues:    make([]int, 0),
		TargetIfAlive: make([]int, 0),
		TargetIfDead:  make([]int, 0),
	}
}

// FindRealLocation function prevent given coordinates to go out of bounds of the Grid. This function takes the targetted position as int with the wanted max/min size of the grid
func findRealLocation(target int, sizeLimit int) (int, error) {

	if target < 0 {
		return negativeLoopThrougth(target, sizeLimit), nil

	} else if target >= sizeLimit-1 {
		return positiveLoopThrougth(target, sizeLimit-1), nil

	} else if target >= 0 && target < sizeLimit {
		return target, nil

	} else {
		return 0, fmt.Errorf("target.go line 27 -> Error when searching for new position, got 'target = %d', 'sizeLimit = %d'", target, sizeLimit)
	}
}

// negativeLoopThrougth function iterate throught a sizeLimit int with a negative target integer
func negativeLoopThrougth(target int, sizeLimit int) int {

	temp := sizeLimit
	for i := target; i <= 0; i = i + 1 {

		if temp == 0 {
			temp = sizeLimit
		}
		temp = temp - 1
	}
	return temp
}

// positiveLoopThrougth function iterate throught a sizeLimit int with a positive target integer
func positiveLoopThrougth(target int, sizeLimit int) int {

	temp := 0
	for i := target; i >= sizeLimit; i = i - 1 {

		if temp == sizeLimit {
			temp = 0
		}
		temp = temp + 1
	}
	return temp
}

// isTargetIn function search from a slice of elements if a value is in the slice, here we use it to check wether the value is in a tagert list */
func isTargetIn(target int, values []int) bool {

	for _, v := range values {
		if target == v {
			return true
		}
	}
	return false
}
