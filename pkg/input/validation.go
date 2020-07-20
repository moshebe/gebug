package input

import (
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

type Validator interface {
	validate(string) error
}

type NonEmptyValidator struct {
	field *string
}

func (v NonEmptyValidator) validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty command")
	}
	*v.field = input
	return nil
}

type RegexValidator struct {
	pattern string
	field   *string
}

func (v RegexValidator) validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty input")
	}
	pattern := regexp.MustCompile(v.pattern)
	if !pattern.MatchString(input) {
		return errors.New("invalid value")
	}
	*v.field = input
	return nil
}

type NumericRangeValidator struct {
	min   int
	max   int
	field *int
}

func (v NumericRangeValidator) validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty input")
	}
	port, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("invalid port")
	}
	if port < v.min || port > v.max {
		return errors.New("port not in valid range")
	}

	*v.field = port
	return nil
}
