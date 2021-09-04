package testutil

import (
	"bytes"
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RunTestData runs all the tests resides in the `testdata` directory and their input/golden file has the following prefix
// using the `check` function the caller can easily consume the test file input and compare the actual result with the golden one.
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
