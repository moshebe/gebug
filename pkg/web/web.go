package web

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/moshebe/gebug/pkg/render"
)

const (
	dockerComposeTemplate = `version: '3'
services:
  gebug-webui:
    image: {{.ImageName}}
    environment:
      - PORT={{.Port}}
      - VUE_APP_GEBUG_PROJECT_LOCATION={{.Location}}
    ports:
      - {{.Port}}:{{.Port}}
    volumes:
      - {{.Location}}:{{.Location}}`
)

// Opts represents the docker-compose configuration file rendering options
type Opts struct {
	ImageName string
	Port      int
	Location  string
}

// RenderDockerCompose renders the docker-compose template and outputs the result into the given writer
func RenderDockerCompose(options *Opts, writer io.Writer) error {
	if options.Port == 0 {
		options.Port = 3030
	}

	if options.ImageName == "" {
		return fmt.Errorf("invalid image name")
	}

	if options.Location == "" {
		return fmt.Errorf("invalid project location")
	}

	out, err := render.Render(dockerComposeTemplate, options)
	if err != nil {
		return fmt.Errorf("render docker-compose template: %w", err)
	}

	_, err = writer.Write([]byte(out))
	if err != nil {
		return fmt.Errorf("write generated docker-compose: %w", err)
	}

	return nil
}

type dummyLogger struct {
}

func (dummyLogger) Printf(string, ...interface{}) {
}

// ReadinessProbe monitors a given URL until it received status OK (200) or gets to the timeout
func ReadinessProbe(url string, verbose bool) error {
	retryClient := retryablehttp.NewClient()
	if !verbose {
		retryClient.Logger = dummyLogger{}
	}
	retryClient.RetryMax = 25
	resp, err := retryClient.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: %d %s", resp.StatusCode, resp.Status)
	}

	return nil
}
