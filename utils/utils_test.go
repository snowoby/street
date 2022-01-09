package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncDec(t *testing.T) {
	p, err := Encrypt("asdddsa")
	fmt.Println(p)
	assert.Nil(t, err)
	assert.True(t, Validate("asdddsa", p))
}
