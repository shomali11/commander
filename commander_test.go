package commander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenize(t *testing.T) {
	tokens := NewCommand("search <pattern>").Tokenize()
	assert.Equal(t, len(tokens), 2)
	assert.Equal(t, tokens[0].Word, "search")
	assert.False(t, tokens[0].IsParameter)
	assert.Equal(t, tokens[1].Word, "pattern")
	assert.True(t, tokens[1].IsParameter)
}

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

	properties, isMatch = NewCommand("help").Match("helpful")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("search all").Match("search for all")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("help").Match("Could you help me?")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("echo <word>").Match("echo")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("echo <word>").Match("echo hey")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("word", ""), "hey")

	properties, isMatch = NewCommand("search <pattern>").Match("search *")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("pattern", ""), "*")

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
	assert.True(t, isParameter("<value>"))
	assert.True(t, isParameter("<123>"))
	assert.True(t, isParameter("<value123>"))
	assert.False(t, isParameter("value>"))
	assert.False(t, isParameter("<value"))
	assert.False(t, isParameter("value"))
	assert.False(t, isParameter(""))
}
