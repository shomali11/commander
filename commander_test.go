package commander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch(t *testing.T) {
	properties, isMatch := NewCommand("").Match("ping")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("").Match("")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("ping").Match("ping")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("ping").Match("pong")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("echo <word>").Match("echo")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("echo <word>").Match("echo hey")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("word", ""), "hey")

	properties, isMatch = NewCommand("repeat <word> <number>").Match("repeat hey 5")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("word", ""), "hey")
	assert.Equal(t, properties.IntegerParam("number", 0), 5)

	properties, isMatch = NewCommand("repeat <word> <number>").Match("repeat hey")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("word", ""), "hey")
	assert.Equal(t, properties.IntegerParam("number", 0), 0)
}

func TestIsParameter(t *testing.T) {
	assert.True(t, IsParameter("<value>"))
	assert.True(t, IsParameter("<123>"))
	assert.True(t, IsParameter("<value123>"))
	assert.False(t, IsParameter("value>"))
	assert.False(t, IsParameter("<value"))
	assert.False(t, IsParameter("value"))
	assert.False(t, IsParameter(""))
}
