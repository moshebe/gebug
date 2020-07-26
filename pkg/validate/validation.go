package validate

import (
	"strconv"
	"strings"

	valid "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type Validator interface {
	Validate(string) error
}

type NonEmptyValidator struct{}

func (v NonEmptyValidator) Validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty command")
	}

	return nil
}

type RegexValidator struct {
	Pattern string
}

func (v RegexValidator) Validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty input")
	}

	if !valid.Matches(input, v.Pattern) {
		return errors.New("input does not matches pattern")
	}

	return nil
}

type NumericRangeValidator struct {
	Min int
	Max int
}

func (v NumericRangeValidator) Validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty input")
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("convert input to a number")
	}

	if !valid.InRange(num, v.Min, v.Max) {
		return errors.Errorf("input is not in range (%d|%d)", v.Min, v.Max)
	}

	return nil
}
