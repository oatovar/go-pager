package pager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUintFromInt(t *testing.T) {
	var expect uint = 10
	got := UintFromInt(-10)
	assert.Equal(t, expect, got)
}
