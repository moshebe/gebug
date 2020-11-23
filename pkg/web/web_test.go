package web

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestRenderDockerCompose(t *testing.T) {
	tests := []struct {
		input      *Opts
		goldenFile string
		wantErr    bool
	}{
		{
			input: &Opts{
				ImageName: "gebug-ui",
				Port:      3030,
				Location:  "/Users/me/Dev/awesome-app",
			},
			wantErr:    false,
			goldenFile: "render_0",
		},
		{
			input: &Opts{
				ImageName: "gebug-ui",
				Location:  "/Users/me/Dev/awesome-app",
			},
			wantErr:    false,
			goldenFile: "render_0",
		},
		{
			input: &Opts{
				Location: "/Users/me/Dev/awesome-app",
			},
			wantErr:    true,
		},
		{
			input: &Opts{
				ImageName: "gebug-ui",
				Port:      3030,
				Location:  "",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		assertion := assert.New(t)
		got := bytes.NewBufferString("")
		err := RenderDockerCompose(test.input, got)
		if test.wantErr {
			assertion.Error(err)
			return
		}

		assertion.NoError(err)

		goldenPath := filepath.Join("testdata", test.goldenFile+".golden")
		goldenData, err := ioutil.ReadFile(goldenPath)
		assertion.NoError(err)

		assertion.Equal(bytes.NewBuffer(goldenData).String(), got.String())
	}

}
