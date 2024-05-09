package filet

import (
	"fmt"
)

type Coordinates struct {
	X int
	Y int
}

func (xyLen Coordinates) findRealTargetLocations(targetAddresses []Coordinates) ([]Coordinates, error) {

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
