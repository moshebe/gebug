package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
// TODO: Get endpoint test

// TODO: Post endpoint test


func TestServer_Start(t *testing.T) {
	t.Log("Starting...")
	s := &Server{
		port: 3030,
		location: ".",
	}
	err := s.Start()
	assert.NoError(t, err)
}
