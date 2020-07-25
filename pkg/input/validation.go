package input

import (
	"strings"

	valid "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type validatorIface interface {
	validate(string) error
}

type nonEmptyValidator struct{}

func (v nonEmptyValidator) validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty command")
	}

	return nil
}

type regexValidator struct {
	pattern string
}

func (v regexValidator) validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty input")
	}

	if !valid.Matches(input, v.pattern) {
		return errors.New("input does not matches pattern")
	}

	return nil
}

type numericRangeValidator struct {
	min int
	max int
}

func (v numericRangeValidator) validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty input")
	}

	if !valid.InRange(input, v.min, v.max) {
		return errors.Errorf("input is not in range (%d|%d)", v.min, v.max)
	}

	return nil
}
