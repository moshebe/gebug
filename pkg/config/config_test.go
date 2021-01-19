package config

import (
	"bytes"
	"github.com/moshebe/gebug/pkg/testutil"
	"io"
	"strconv"
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
runtime_image: golang:1.15.2
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
				RuntimeImage:     "golang:1.15.2",
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
runtime_image: golang:1.15.2
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
				RuntimeImage:     "golang:1.15.2",
				DebuggerPort:     40000,
				ExposePorts:      []string{"8080:8080"},
			},
			wantErr: false,
		},
		{
			input:   `invalid yaml`,
			wantErr: true,
		},
		{
			input:   `name: {{.NotExists}}`,
			wantErr: true,
		},
	}

	for _, test := range tests {
		c, err := Load([]byte(test.input))
		if test.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, c)
			assert.Equal(t, test.expected, c)
		}
	}
}

func testGenerateHelper(t *testing.T, input, golden *bytes.Buffer, generate func(config *Config, writer io.Writer) error) {
	assertion := assert.New(t)
	c, err := Load(input.Bytes())
	assertion.NoError(err)
	assertion.NotNil(c)

	got := bytes.NewBufferString("")
	err = generate(c, got)
	assertion.NoError(err)
	assertion.Equal(golden.String(), got.String())
}

func TestConfig_GenerateDockerfile(t *testing.T) {
	testutil.RunTestData(t, "generate_dockerfile", func(t *testing.T, input, golden *bytes.Buffer) {
		testGenerateHelper(t, input, golden, func(c *Config, writer io.Writer) error {
			return c.RenderDockerfile(writer)
		})
	})
}

func TestConfig_GenerateDockerCompose(t *testing.T) {
	testutil.RunTestData(t, "generate_docker_compose", func(t *testing.T, input, golden *bytes.Buffer) {
		testGenerateHelper(t, input, golden, func(c *Config, writer io.Writer) error {
			return c.RenderDockerComposeFile(writer)
		})
	})
}

func TestConfig_Write(t *testing.T) {
	tests := []struct {
		input    *Config
		expected string
		wantErr  bool
	}{
		{
			input: &Config{
				Name:             "my-app",
				OutputBinaryPath: "/app",
				BuildCommand:     "go build -o /app -gcflags=\"all=-N-l\"",
				BuildDir:         "/app",
				RunCommand:       "/app",
				RuntimeImage:     "golang:1.15.2",
				DebuggerPort:     40000,
				DebuggerEnabled:  true,
				ExposePorts:      []string{"8080", "8081:8081"},
			},
			expected: `name: my-app
output_binary: /app
build_command: go build -o /app -gcflags="all=-N-l"
build_dir: /app
run_command: /app
runtime_image: golang:1.15.2
debugger_enabled: true
debugger_port: 40000
expose_ports:
- "8080"
- 8081:8081
networks: []
environment: []
`,
			wantErr: false,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			buf := bytes.NewBufferString("")
			err := test.input.Write(buf)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, buf.String())
			}
		})
	}
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

func TestConfig_ResolvePath(t *testing.T) {
	expected := "/Users/me/Dev/project/.gebug/config.yaml"
	variations := []string{
		"/Users/me/Dev/project/.gebug/config.yaml",
		"/Users/me/Dev/project/.gebug/",
		"/Users/me/Dev/project/.gebug",
		"/Users/me/Dev/project/",
		"/Users/me/Dev/project",
	}

	for _, input := range variations {
		got := ResolvePath(input)
		assert.Equal(t, got, expected)
	}
}
