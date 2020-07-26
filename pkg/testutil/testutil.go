package testutil

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func RunTestData(t *testing.T, prefix string, check func(t *testing.T, input, golden *bytes.Buffer)) {
	baseDir := "./testdata"
	files, err := ioutil.ReadDir(baseDir)
	assert.NoError(t, err)

	for _, file := range files {
		name := file.Name()
		if (prefix+".in" != name && !strings.HasPrefix(name, prefix+"_")) || !strings.HasSuffix(name, ".in") {
			continue
		}

		t.Run(name, func(t *testing.T) {
			assertion := assert.New(t)
			filePath := path.Join(baseDir, name)
			input, err := ioutil.ReadFile(filePath)
			assertion.NoError(err)
			golden, err := ioutil.ReadFile(strings.TrimSuffix(filePath, ".in") + ".golden")
			assertion.NoError(err)
			check(t, bytes.NewBuffer(input), bytes.NewBuffer(golden))
		})
	}
}