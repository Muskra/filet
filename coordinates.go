package filet

import (
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

// FindRealTargetLocations function checks if any targetAddresses Coordinates Type from a given Cell Coordinates Type would be unreachable or out of bound of the Grid Type. For example your grid is of size 4x4 but you want to found the location 'x = -46' along the value 'y = 82', the function cicle througth the grid to find the actual position where we want to target.
func (xyLen Coordinates) FindRealTargetLocations(targetAddresses []Coordinates) ([]Coordinates, error) {

	ret := make([]Coordinates, 0)

	for _, coordinate := range targetAddresses {

		tempValueOfX, err := findRealLocation(coordinate.X, xyLen.X)
		if err != nil {
			return []Coordinates{}, fmt.Errorf("coordinates.go line 20 -> %s", err)
		}

		tempValueOfY, err := findRealLocation(coordinate.Y, xyLen.Y)
		if err != nil {
			return []Coordinates{}, fmt.Errorf("coordinates.go line 25 -> %s", err)
		}
		ret = append(ret, Coordinates{X: tempValueOfX, Y: tempValueOfY})
	}
	return ret, nil

}
