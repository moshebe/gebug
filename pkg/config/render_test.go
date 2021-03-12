package config

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockConfig = &Config{
	Name:             "my-app",
	OutputBinaryPath: "/app",
	BuildCommand:     "go build",
	RunCommand:       "/app",
	RuntimeImage:     "golang:1.15.2",
	DebuggerEnabled:  false,
	DebuggerPort:     0,
	ExposePorts:      []string{"8080"},
}

func TestConfig_RenderDockerComposeFileCamlCaseAppName(t *testing.T) {
	cfg := *mockConfig
	cfg.Name = "myApp2"
	out := bytes.NewBufferString("")
	err := cfg.RenderDockerComposeFile(out)
	assert.NoError(t, err)
	assert.Equal(t,
		`version: '3'
services:
  gebug-my-app-2:
    build:
      context: ..
      dockerfile: .gebug/Dockerfile
    volumes:
      - ../:/src:ro
    ports:
      - 8080
`,
		out.String())
}

func TestConfig_RenderDockerComposeFile(t *testing.T) {
	out := bytes.NewBufferString("")
	err := mockConfig.RenderDockerComposeFile(out)
	assert.NoError(t, err)
	assert.Equal(t,
		`version: '3'
services:
  gebug-my-app:
    build:
      context: ..
      dockerfile: .gebug/Dockerfile
    volumes:
      - ../:/src:ro
    ports:
      - 8080
`,
		out.String())
}

func TestConfig_RenderDockerComposeFile_NoPorts(t *testing.T) {
	portsTable := [][]string{
		nil,
		make([]string, 0),
	}

	for _, ports := range portsTable {
		out := bytes.NewBufferString("")
		mockConfig.ExposePorts = ports
		err := mockConfig.RenderDockerComposeFile(out)
		assert.NoError(t, err)
		assert.Equal(t,
			`version: '3'
services:
  gebug-my-app:
    build:
      context: ..
      dockerfile: .gebug/Dockerfile
    volumes:
      - ../:/src:ro
`,
			out.String())
	}
}

func TestConfig_RenderDockerfile(t *testing.T) {
	out := bytes.NewBufferString("")
	err := mockConfig.RenderDockerfile(out)
	assert.NoError(t, err)
	assert.Equal(t,
		`FROM golang:1.15.2
RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /src
COPY . .

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build" -command="/app"`,
		out.String())
}
