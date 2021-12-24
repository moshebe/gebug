package validate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNonEmptyValidator(t *testing.T) {
	tests := []struct {
		input string
		err   require.ErrorAssertionFunc
	}{
		{input: "", err: require.Error},
		{input: "  ", err: require.Error},
		{input: " \t", err: require.Error},
		{input: " \t\n", err: require.Error},
		{input: "hello", err: require.NoError},
		{input: "hello-world", err: require.NoError},
		{input: "hello world", err: require.NoError},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%s#%d", t.Name(), i), func(t *testing.T) {
			tt.err(t, NonEmptyValidator{}.Validate(tt.input))
		})
	}
}

func TestRegexValidator(t *testing.T) {
	tests := []struct {
		input string
		err   require.ErrorAssertionFunc
	}{
		{input: "", err: require.Error},
		{input: "  ", err: require.Error},
		{input: " \t", err: require.Error},
		{input: " \t\n", err: require.Error},
		{input: "!hello", err: require.Error},
		{input: "hello world", err: require.Error},
		{input: "0abc", err: require.NoError},
		{input: " hello", err: require.NoError},
		{input: "hello", err: require.NoError},
		{input: "hello-world", err: require.NoError},
	}

	valiator := RegexValidator{Pattern: `^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%s#%d", t.Name(), i), func(t *testing.T) {
			tt.err(t, valiator.Validate(tt.input))
		})
	}
}

func TestNumericRangeValidator(t *testing.T) {
	tests := []struct {
		input string
		err   require.ErrorAssertionFunc
	}{
		{input: "", err: require.Error},
		{input: "  ", err: require.Error},
		{input: " \t", err: require.Error},
		{input: " \t\n", err: require.Error},
		{input: "!hello", err: require.Error},
		{input: "0", err: require.Error},
		{input: "a1", err: require.Error},
		{input: "11", err: require.Error},
		{input: "1", err: require.NoError},
		{input: "5", err: require.NoError},
		{input: "10", err: require.NoError},
	}

	valiator := NumericRangeValidator{Min: 1, Max: 10}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%s#%d", t.Name(), i), func(t *testing.T) {
			tt.err(t, valiator.Validate(tt.input))
		})
	}
}
