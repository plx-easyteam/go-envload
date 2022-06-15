package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T){ 
	expected := "##### This is a message found in .env file #####"
	actual := HelloMsg()

	assert.Equal(t, expected, actual)
}