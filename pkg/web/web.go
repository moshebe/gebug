package web

import (
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/moshebe/gebug/pkg/render"
	"github.com/pkg/errors"
	"io"
	"net/http"
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

	defaultPort      = 3030
	defaultImageName = "gebug-ui" // TODO: replace?? also in the tests and goldens
)

type Opts struct {
	ImageName string
	Port      int
	Location  string
}

func RenderDockerCompose(options *Opts, writer io.Writer) error {
	if options.Port == 0 {
		options.Port = 3030
	}

	if options.ImageName == "" {
		options.ImageName = defaultImageName
	}

	if options.Location == "" {
		return errors.New("invalid project location")
	}

	out, err := render.Render(dockerComposeTemplate, options)
	if err != nil {
		return errors.WithMessage(err, "render docker-compose template")
	}

	_, err = writer.Write([]byte(out))
	if err != nil {
		return errors.WithMessage(err, "write generated docker-compose")
	}

	return nil
}

type dummyLogger struct {
}

func (dummyLogger) Printf(string, ...interface{}) {
}
func ReadinessProbe(url string, verbose bool) error {
	retryClient := retryablehttp.NewClient()
	if !verbose {
		retryClient.Logger = dummyLogger{}
	}
	retryClient.RetryMax = 10
	resp, err := retryClient.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: %d %s", resp.StatusCode, resp.Status)
	}

	return nil
}
