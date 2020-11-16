package webui

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer_Start(t *testing.T) {
	t.Log("Starting...")
	s := &Server{
		Port: 3030,
	}
	err := s.Start()
	assert.NoError(t, err)
}
