package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	expected := "##### Addition #####"
	actual := WhoAmI()

	assert.Equal(t, expected, actual)
}