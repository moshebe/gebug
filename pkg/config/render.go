package config

import (
	"gebug/pkg/render"
	"github.com/pkg/errors"
	"io"
)

func (c *Config) renderedWrite(template string, writer io.Writer) error {
	out, err := render.Render(template, c)
	if err != nil {
		return errors.WithMessage(err, "render docker-compose template")
	}

	_, err = writer.Write([]byte(out))
	if err != nil {
		return errors.WithMessage(err, "write generated configuration")
	}

	return nil
}

func (c *Config) RenderDockerComposeFile(writer io.Writer) error {
	return c.renderedWrite(`version: '3'
services:
  gebug-{{.Name}}:
    build:
      context: ..
      dockerfile: .gebug/Dockerfile{{if .DebuggerEnabled}}
    cap_add:
      - SYS_PTRACE{{end}}
    volumes:
      - ../:/src:ro
    ports:{{ range $key, $value := .ExposePorts }}
      - {{ $value }}{{ end }}
{{if .DebuggerEnabled}}      - {{.DebuggerPort}}:{{.DebuggerPort}}{{end}}`, writer)
}

func (c *Config) RenderDockerfile(writer io.Writer) error {
	return c.renderedWrite(`FROM {{.RuntimeImage}}
RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /src
COPY . .

{{if .DebuggerEnabled}}RUN {{.BuildCommand}}
ENTRYPOINT dlv --listen=:{{.DebuggerPort}} --headless=true --api-version=2 --accept-multiclient exec {{.OutputBinaryPath}}{{else}}ENTRYPOINT CompileDaemon -log-prefix=false -build="{{.BuildCommand}}" -command="{{.RunCommand}}"{{end}}`, writer)
}
