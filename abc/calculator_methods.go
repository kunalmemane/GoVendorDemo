package abc

import (
	"fmt"
	"math"
)

// Add adds a number to the current value
func (c *Calculator) Add(num float64) {
	oldValue := c.Value
	c.Value += num
	c.Operation = fmt.Sprintf("Add %.2f", num)
	c.addToHistory(fmt.Sprintf("%.2f + %.2f = %.2f", oldValue, num, c.Value))
}

// Subtract subtracts a number from the current value
func (c *Calculator) Subtract(num float64) {
	oldValue := c.Value
	c.Value -= num
	c.Operation = fmt.Sprintf("Subtract %.2f", num)
	c.addToHistory(fmt.Sprintf("%.2f - %.2f = %.2f", oldValue, num, c.Value))
}

// Multiply multiplies the current value by a number
func (c *Calculator) Multiply(num float64) {
	oldValue := c.Value
	c.Value *= num
	c.Operation = fmt.Sprintf("Multiply by %.2f", num)
	c.addToHistory(fmt.Sprintf("%.2f * %.2f = %.2f", oldValue, num, c.Value))
}

// Divide divides the current value by a number
func (c *Calculator) Divide(num float64) error {
	if num == 0 {
		return fmt.Errorf("division by zero is not allowed")
	}
	oldValue := c.Value
	c.Value /= num
	c.Operation = fmt.Sprintf("Divide by %.2f", num)
	c.addToHistory(fmt.Sprintf("%.2f / %.2f = %.2f", oldValue, num, c.Value))
	return nil
}

// Power raises the current value to a power
func (c *Calculator) Power(exponent float64) {
	oldValue := c.Value
	c.Value = math.Pow(c.Value, exponent)
	c.Operation = fmt.Sprintf("Power %.2f", exponent)
	c.addToHistory(fmt.Sprintf("%.2f ^ %.2f = %.2f", oldValue, exponent, c.Value))
}

// SquareRoot calculates the square root of the current value
func (c *Calculator) SquareRoot() error {
	if c.Value < 0 {
		return fmt.Errorf("cannot calculate square root of negative number")
	}
	oldValue := c.Value
	c.Value = math.Sqrt(c.Value)
	c.Operation = "Square Root"
	c.addToHistory(fmt.Sprintf("√%.2f = %.2f", oldValue, c.Value))
	return nil
}

// Reset sets the calculator value back to zero
func (c *Calculator) Reset() {
	c.Value = 0
	c.Operation = "Reset"
	c.addToHistory("Reset to 0.00")
}

// GetValue returns the current value
func (c *Calculator) GetValue() float64 {
	return c.Value
}

// GetHistory returns the operation history
func (c *Calculator) GetHistory() []string {
	return c.History
}

// ClearHistory clears the operation history
func (c *Calculator) ClearHistory() {
	c.History = make([]string, 0)
	c.Operation = "History Cleared"
}

// String returns a string representation of the Calculator
func (c *Calculator) String() string {
	return fmt.Sprintf("Calculator{Value: %.2f, LastOperation: %s, HistoryEntries: %d}",
		c.Value, c.Operation, len(c.History))
}

// addToHistory is a helper method to add entries to history
func (c *Calculator) addToHistory(entry string) {
	c.History = append(c.History, entry)
	// Keep only last 100 entries
	if len(c.History) > 100 {
		c.History = c.History[len(c.History)-100:]
	}
}
