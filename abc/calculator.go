package abc

// Calculator represents a calculator with a current value
type Calculator struct {
	Value     float64
	History   []string
	Operation string
}

// NewCalculator creates a new Calculator instance with initial value
func NewCalculator(initialValue float64) *Calculator {
	return &Calculator{
		Value:     initialValue,
		History:   make([]string, 0),
		Operation: "initialized",
	}
}
