package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testScenario struct {
	input   string
	wantErr bool
}

func testValidator(t *testing.T, validator Validator, tests []testScenario) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("%s#%d", t.Name(), i), func(t *testing.T) {
			err := validator.Validate(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNonEmptyValidator(t *testing.T) {
	testValidator(t, NonEmptyValidator{}, []testScenario{
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
	testValidator(t, RegexValidator{Pattern: `^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`}, []testScenario{
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
	testValidator(t, NumericRangeValidator{Min: 1, Max: 10}, []testScenario{
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
		{input: "10", wantErr: false},
	})
}
