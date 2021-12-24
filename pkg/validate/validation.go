package validate

import (
	"fmt"
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

// Validator defines the behaviour validation behaviour
type Validator interface {
	Validate(string) error
}

// NonEmptyValidator checks that the input is not empty after trimming
type NonEmptyValidator struct{}

// Validate checks the input and return an error if its invalid
func (v NonEmptyValidator) Validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return fmt.Errorf("empty command")
	}

	return nil
}

// RegexValidator checks that the input matches a pattern after trimming
type RegexValidator struct {
	Pattern string
}

// Validate checks the input and return an error if its invalid
func (v RegexValidator) Validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return fmt.Errorf("empty input")
	}

	if !valid.Matches(input, v.Pattern) {
		return fmt.Errorf("input does not matches pattern")
	}

	return nil
}

// NumericRangeValidator checks that the input is a valid number after trimming and its value is between a given range
type NumericRangeValidator struct {
	Min int
	Max int
}

// Validate checks the input and return an error if its invalid
func (v NumericRangeValidator) Validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return fmt.Errorf("empty input")
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("convert input to a number")
	}

	if !valid.InRange(num, v.Min, v.Max) {
		return fmt.Errorf("input is not in range (%d|%d)", v.Min, v.Max)
	}

	return nil
}
