package client

import (
	"design_patterns/creational/factory_method"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntFactory(t *testing.T) {
	nf := factory_method.Get("int")
	assert.NotNil(t, nf)

	parse, err := nf.Parse("666")
	assert.Nil(t, err)
	assert.Equal(t, 666, parse)
}

func TestFloatFactory(t *testing.T) {
	nf := factory_method.Get("float")
	assert.NotNil(t, nf)
	parse, err := nf.Parse("666.6")
	assert.Nil(t, err)
	assert.Equal(t, 666.6, parse)
}
