package render

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

// Render builds a string
//
// rawTemplate should contain a valid template according to https://golang.org/pkg/text/template/
// data should contain any struct that is needed to render the template
func Render(rawTemplate string, data interface{}) (string, error) {
	t, err := template.New("").Parse(rawTemplate)
	if err != nil {
		return "", errors.WithMessage(err, "parse template")
	}

	var out bytes.Buffer
	err = t.Execute(&out, data)
	if err != nil {
		return "", errors.WithMessage(err, "render template")
	}

	return out.String(), nil
}
