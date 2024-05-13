package filet

import ()

// Set Type is being used to define the actual rules and conditions for the automata
type Set struct {
	CellValue        int
	CellState        bool
	TargetValue      int
	TargetState      bool
	ShouldBeTargeted bool
	Opcode           uint8
}

// NewSet function returns an empty Set Type
func NewSet() Set {
	return Set{
		CellValue:        0,
		CellState:        false,
		TargetValue:      0,
		TargetState:      false,
		ShouldBeTargeted: false,
		Opcode:           0,
	}
}
