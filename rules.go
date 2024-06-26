package filet

import ()

// Rules Type are assembling some of other structs, where Target and Set are being used. Another slice is used here to define where all the targets will be at in the grid, if you decide to go out of bounds, it will loop througth the sides of the grid to prevent errors.
type Rules struct {
	RuleSet              []Set
	TargetCellsLocations []Coordinates
	TargetValues         Target
}

// NewRules function returns an empty Rules Type
func NewRules() Rules {
	return Rules{
		RuleSet:              make([]Set, 0),
		TargetCellsLocations: make([]Coordinates, 0),
		TargetValues:         NewTarget(),
	}
}

// ruleCheck function checks if a rule is applicable from the context of the actual cell and it's targetted value
func ruleCheck(ruleSet []Set, cell Cell, target Cell) (int, bool) {

	for index, rule := range ruleSet {

		// checks if the Set cell and target values are identical and check if the rule should be processed from the isIn value and the shouldBeTargeted
		if cell.Equal(rule, target) {
			return index, true
		}
	}
	return 0, false
}
