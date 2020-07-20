package input

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testScenario struct {
	input   string
	wantErr bool
}

func testValidator(t *testing.T, validator Validator, tests []testScenario) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("%s#%d", t.Name(), i), func(t *testing.T) {
			err := validator.validate(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNonEmptyValidator(t *testing.T) {
	var dummy string
	testValidator(t, NonEmptyValidator{field: &dummy}, []testScenario{
		{input: "", wantErr: true},
		{input: "  ", wantErr: true},
		{input: " \t", wantErr: true},
		{input: " \t\n", wantErr: true},
		{input: "hello", wantErr: false},
		{input: "hello-world", wantErr: false},
		{input: "hello world", wantErr: false},
	})
}

func TestRegexValidator(t *testing.T) {
	var dummy string
	testValidator(t, RegexValidator{pattern: `^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`, field: &dummy}, []testScenario{
		{input: "", wantErr: true},
		{input: "  ", wantErr: true},
		{input: " \t", wantErr: true},
		{input: " \t\n", wantErr: true},
		{input: "!hello", wantErr: true},
		{input: "hello world", wantErr: true},
		{input: "0abc", wantErr: false},
		{input: " hello", wantErr: false},
		{input: "hello", wantErr: false},
		{input: "hello-world", wantErr: false},
	})
}

func TestNumericRangeValidator(t *testing.T) {
	var dummy int
	testValidator(t, NumericRangeValidator{min: 1, max: 10, field: &dummy}, []testScenario{
		{input: "", wantErr: true},
		{input: "  ", wantErr: true},
		{input: " \t", wantErr: true},
		{input: " \t\n", wantErr: true},
		{input: "!hello", wantErr: true},
		{input: "0", wantErr: true},
		{input: "a1", wantErr: true},
		{input: "11", wantErr: true},
		{input: "1", wantErr: false},
		{input: "5", wantErr: false},
	})
}
