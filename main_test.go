package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI")
	assert.Equal(t, convert("A", 1), "A")
	assert.Equal(t, convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR")
}
