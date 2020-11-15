package webui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer_Start(t *testing.T) {
	s := &Server{
		Port: 3030,
	}
	err := s.Start()
	assert.NoError(t, err)
}
