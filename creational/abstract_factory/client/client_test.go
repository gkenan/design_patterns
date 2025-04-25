package client

import (
	"design_patterns/creational/abstract_factory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFastFactoryHtml(t *testing.T) {
	factory := abstract_factory.CreateFactory("fast")
	assert.NotNil(t, factory)
	html := factory.CreateHtml()
	assert.NotNil(t, html)
	err := html.Save("test.html")
	assert.Nil(t, err)
	toHtml := html.ToHtml()
	assert.Equal(t, "Fast Html Document: test.html", toHtml)
}

func TestFastFactoryWord(t *testing.T) {
	factory := abstract_factory.CreateFactory("fast")
	assert.NotNil(t, factory)
	html := factory.CreateWord()
	assert.NotNil(t, html)
	err := html.Save("test.word")
	assert.Nil(t, err)
	toHtml := html.ToWord()
	assert.Equal(t, "Fast Word Document: test.word", toHtml)
}

func TestGoodFactoryHtml(t *testing.T) {
	factory := abstract_factory.CreateFactory("good")
	assert.NotNil(t, factory)
	html := factory.CreateHtml()
	assert.NotNil(t, html)
	err := html.Save("test.html")
	assert.Nil(t, err)
	toHtml := html.ToHtml()
	assert.Equal(t, "Good Html Document: test.html", toHtml)
}

func TestGoodFactoryWord(t *testing.T) {
	factory := abstract_factory.CreateFactory("good")
	assert.NotNil(t, factory)
	html := factory.CreateWord()
	assert.NotNil(t, html)
	err := html.Save("test.word")
	assert.Nil(t, err)
	toHtml := html.ToWord()
	assert.Equal(t, "Good Word Document: test.word", toHtml)
}
