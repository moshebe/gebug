package render

import (
	"bytes"
	"fmt"
	"text/template"
)

// Render builds a string
//
// rawTemplate should contain a valid template according to https://golang.org/pkg/text/template/
// data should contain any struct that is needed to render the template
func Render(rawTemplate string, data interface{}) (string, error) {
	t, err := template.New("").Parse(rawTemplate)
	if err != nil {
		return "", fmt.Errorf("parse template: %w", err)
	}

	var out bytes.Buffer
	err = t.Execute(&out, data)
	if err != nil {
		return "", fmt.Errorf("render template: %w", err)
	}

	return out.String(), nil
}
