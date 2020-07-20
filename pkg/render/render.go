package render

import (
	"bytes"
	"github.com/pkg/errors"
	"text/template"
)

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
