package web

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenderDockerCompose(t *testing.T) {
	tests := []struct {
		input      *Opts
		goldenFile string
		wantErr    bool
	}{
		{
			input: &Opts{
				ImageName: "gebug/webui",
				Port:      3030,
				Location:  "/Users/me/Dev/awesome-app",
			},
			wantErr:    false,
			goldenFile: "render_0",
		},
		{
			input: &Opts{
				ImageName: "gebug/webui",
				Location:  "/Users/me/Dev/awesome-app",
			},
			wantErr:    false,
			goldenFile: "render_0",
		},
		{
			input: &Opts{
				Location: "/Users/me/Dev/awesome-app",
			},
			wantErr: true,
		},
		{
			input: &Opts{
				ImageName: "gebug/webui",
				Port:      3030,
				Location:  "",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		got := bytes.NewBufferString("")
		err := RenderDockerCompose(test.input, got)
		if test.wantErr {
			require.Error(t, err)
			return
		}

		require.NoError(t, err)

		goldenPath := filepath.Join("testdata", test.goldenFile+".golden")
		goldenData, err := ioutil.ReadFile(goldenPath)
		require.NoError(t, err)

		require.Equal(t, bytes.NewBuffer(goldenData).String(), got.String())
	}

}
