package selector

var (
	_ Choice         = (*choice)(nil)         // Assert that choice implements Choice
	_ WeightedChoice = (*weightedChoice)(nil) // Assert that weightedChoice implements WeightedChoice
)

type (
	// Choice wraps a selection target
	Choice interface {
		Item() interface{}
	}

	// WeightedChoice is an item selection target with a non negative weight
	WeightedChoice interface {
		Choice
		Weight() uint
	}

	choice struct {
		item interface{}
	}

	weightedChoice struct {
		choice
		weight uint
	}
)

func NewChoice(item interface{}) Choice {
	return &choice{
		item: item,
	}
}

func (c *choice) Item() interface{} {
	return c.item
}

// NewWeightedChoice returns a choice that has a probability weight
func NewWeightedChoice(item interface{}, weight uint) WeightedChoice {
	return &weightedChoice{
		choice: choice{
			item: item,
		},
		weight: weight,
	}
}

func (c *weightedChoice) Weight() uint {
	return c.weight
}
