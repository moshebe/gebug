package config

import (
	"bytes"
	"io"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Load(t *testing.T) {
	tests := []struct {
		input    string
		expected *Config
		wantErr  bool
	}{
		{
			input: `
name: my-app
output_binary: /app
build_command: go build -o {{.output_binary}} -gcflags="all=-N-l"
runtime_image: golang:1.14
run_command: '{{.output_binary}}'
debugger_port: 40000
expose_ports:
  - 8080:8080
`,
			expected: &Config{
				Name:             "my-app",
				OutputBinaryPath: "/app",
				BuildCommand:     `go build -o /app -gcflags="all=-N-l"`,
				RunCommand:       "/app",
				RuntimeImage:     "golang:1.14",
				DebuggerPort:     40000,
				ExposePorts:      []string{"8080:8080"},
			},
			wantErr: false,
		},
		{
			input: `
name: my-app
output_binary: /app
build_command: go build -o {{.output_binary}} -gcflags="all=-N-l"
runtime_image: golang:1.14
run_command: '{{.output_binary}}'
debugger_port: 40000
expose_ports:
  - 8080:8080
`,
			expected: &Config{
				Name:             "my-app",
				OutputBinaryPath: "/app",
				BuildCommand:     `go build -o /app -gcflags="all=-N-l"`,
				RunCommand:       "/app",
				RuntimeImage:     "golang:1.14",
				DebuggerPort:     40000,
				ExposePorts:      []string{"8080:8080"},
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		c, err := Load([]byte(test.input))
		assert.NoError(t, err)
		assert.NotNil(t, c)
		assert.Equal(t, test.expected, c)
	}
}

func TestGenerateFiles(t *testing.T) {
	tests := []struct {
		prefix   string
		generate func(*Config, io.Writer) error
	}{
		{
			prefix: "generate_docker_compose",
			generate: func(config *Config, writer io.Writer) error {
				return config.RenderDockerComposeFile(writer)
			},
		},
		{
			prefix: "generate_dockerfile",
			generate: func(config *Config, writer io.Writer) error {
				return config.RenderDockerfile(writer)
			},
		},
	}
	for _, test := range tests {
		generateFile(t, test.prefix, test.generate)
	}
}

func generateFile(t *testing.T, prefix string, generate func(*Config, io.Writer) error) {
	baseDir := "./testdata"
	files, err := ioutil.ReadDir(baseDir)
	assert.NoError(t, err)

	for _, file := range files {
		name := file.Name()
		if !strings.HasPrefix(name, prefix+"_") || !strings.HasSuffix(name, ".in") {
			continue
		}

		t.Run(name, func(t *testing.T) {
			assertion := assert.New(t)
			filePath := path.Join(baseDir, name)
			input, err := ioutil.ReadFile(filePath)
			assertion.NoError(err)
			golden, err := ioutil.ReadFile(strings.TrimSuffix(filePath, ".in") + ".golden")
			assertion.NoError(err)

			c, err := Load(input)
			assertion.NoError(err)
			assertion.NotNil(c)

			out := bytes.NewBufferString("")
			err = generate(c, out)
			assertion.NoError(err)
			assertion.Equal(string(golden), out.String())
		})
	}
}

func TestConfig_Write(t *testing.T) {
	expected := `name: my-app
output_binary: /app
build_command: go build -o /app -gcflags="all=-N-l"
run_command: /app
runtime_image: golang:1.14
debugger_enabled: true
debugger_port: 40000
expose_ports:
- "8080"
- 8081:8081
networks: []
`
	c := Config{
		Name:             "my-app",
		OutputBinaryPath: "/app",
		BuildCommand:     "go build -o /app -gcflags=\"all=-N-l\"",
		RunCommand:       "/app",
		RuntimeImage:     "golang:1.14",
		DebuggerPort:     40000,
		DebuggerEnabled:  true,
		ExposePorts:      []string{"8080", "8081:8081"},
	}

	buf := bytes.NewBufferString("")
	err := c.Write(buf)
	assert.NoError(t, err)
	assert.Equal(t, expected, buf.String())
}

func TestConfig_updateBuildCommand(t *testing.T) {
	tests := []struct {
		command  string
		enabled  bool
		expected string
	}{
		{
			command:  "go build",
			enabled:  true,
			expected: `go build -gcflags="all=-N -l"`,
		},
		{
			command:  "go build -o out",
			enabled:  true,
			expected: `go build -gcflags="all=-N -l" -o out`,
		},
		{
			command:  "go build -o out cmd/main.go",
			enabled:  true,
			expected: `go build -gcflags="all=-N -l" -o out cmd/main.go`,
		},
		{
			command:  `go build -gcflags="all=-N -l" -o out cmd/main.go`,
			enabled:  true,
			expected: `go build -gcflags="all=-N -l" -o out cmd/main.go`,
		},
		{
			command:  "go build -o out cmd/main.go",
			enabled:  false,
			expected: `go build -o out cmd/main.go`,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := updateBuildCommand(test.command, test.enabled)
			assert.Equal(t, test.expected, got)
		})
	}
}
