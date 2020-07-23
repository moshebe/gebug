package input

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type validatorIface interface {
	validate(string) error
}

type nonEmptyValidator struct {
	field *string
}

func (v nonEmptyValidator) validate(input string) error {
	input = strings.TrimSpace(input)
	if len(input) <= 0 {
		return errors.New("empty command")
	}
	*v.field = input
	return nil
}

type regexValidator struct {
	pattern string
	field   *string
}

func (v regexValidator) validate(input string) error {
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

type numericRangeValidator struct {
	min   int
	max   int
	field *int
}

func (v numericRangeValidator) validate(input string) error {
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
