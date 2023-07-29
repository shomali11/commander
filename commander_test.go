package commander

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	tokens := NewCommand("say    <input>   ").Tokenize()
	assert.Equal(t, len(tokens), 2)
	assert.Equal(t, tokens[0].Word, "say")
	assert.False(t, tokens[0].IsParameter())
	assert.Equal(t, tokens[1].Word, "input")
	assert.True(t, tokens[1].IsParameter())

	tokens = NewCommand("search <pattern>").Tokenize()
	assert.Equal(t, len(tokens), 2)
	assert.Equal(t, tokens[0].Word, "search")
	assert.False(t, tokens[0].IsParameter())
	assert.Equal(t, tokens[1].Word, "pattern")
	assert.True(t, tokens[1].IsParameter())

	tokens = NewCommand("<a> <123> <a123> <a-123> <a.123> b> <c e f? g?> <h?").Tokenize()
	assert.Equal(t, len(tokens), 11)
	assert.Equal(t, tokens[0].Word, "a")
	assert.True(t, tokens[0].IsParameter())
	assert.Equal(t, tokens[1].Word, "123")
	assert.True(t, tokens[1].IsParameter())
	assert.Equal(t, tokens[2].Word, "a123")
	assert.True(t, tokens[2].IsParameter())
	assert.Equal(t, tokens[3].Word, "a-123")
	assert.True(t, tokens[3].IsParameter())
	assert.Equal(t, tokens[4].Word, "a.123")
	assert.True(t, tokens[4].IsParameter())
	assert.Equal(t, tokens[5].Word, "b>")
	assert.False(t, tokens[5].IsParameter())
	assert.Equal(t, tokens[6].Word, "<c")
	assert.False(t, tokens[6].IsParameter())
	assert.Equal(t, tokens[7].Word, "e")
	assert.False(t, tokens[7].IsParameter())
	assert.Equal(t, tokens[8].Word, "f?")
	assert.False(t, tokens[8].IsParameter())
	assert.Equal(t, tokens[9].Word, "g?>")
	assert.False(t, tokens[9].IsParameter())
	assert.Equal(t, tokens[10].Word, "<h?")
	assert.False(t, tokens[10].IsParameter())

	tokens = NewCommand("{a} {123} {a123} {a-123} {a.123} b} {c e f? g?} {h?").Tokenize()
	assert.Equal(t, len(tokens), 11)

	assert.Equal(t, tokens[0].Word, "a")
	assert.True(t, tokens[0].IsParameter())
	assert.Equal(t, tokens[1].Word, "123")
	assert.True(t, tokens[1].IsParameter())
	assert.Equal(t, tokens[2].Word, "a123")
	assert.True(t, tokens[2].IsParameter())
	assert.Equal(t, tokens[3].Word, "a-123")
	assert.True(t, tokens[3].IsParameter())
	assert.Equal(t, tokens[4].Word, "a.123")
	assert.True(t, tokens[4].IsParameter())
	assert.Equal(t, tokens[5].Word, "b}")
	assert.False(t, tokens[5].IsParameter())
	assert.Equal(t, tokens[6].Word, "{c")
	assert.False(t, tokens[6].IsParameter())
	assert.Equal(t, tokens[7].Word, "e")
	assert.False(t, tokens[7].IsParameter())
	assert.Equal(t, tokens[8].Word, "f?")
	assert.False(t, tokens[8].IsParameter())
	assert.Equal(t, tokens[9].Word, "g?}")
	assert.False(t, tokens[9].IsParameter())
	assert.Equal(t, tokens[10].Word, "{h?")
	assert.False(t, tokens[10].IsParameter())

	tokens = NewCommand("\\ ( ) { } [ ] ? . + | ^ $").Tokenize()
	assert.Equal(t, len(tokens), 13)
	assert.Equal(t, tokens[0].Word, "\\")
	assert.False(t, tokens[0].IsParameter())
	assert.Equal(t, tokens[1].Word, "(")
	assert.False(t, tokens[1].IsParameter())
	assert.Equal(t, tokens[2].Word, ")")
	assert.False(t, tokens[2].IsParameter())
	assert.Equal(t, tokens[3].Word, "{")
	assert.False(t, tokens[3].IsParameter())
	assert.Equal(t, tokens[4].Word, "}")
	assert.False(t, tokens[4].IsParameter())
	assert.Equal(t, tokens[5].Word, "[")
	assert.False(t, tokens[5].IsParameter())
	assert.Equal(t, tokens[6].Word, "]")
	assert.False(t, tokens[6].IsParameter())
	assert.Equal(t, tokens[7].Word, "?")
	assert.False(t, tokens[7].IsParameter())
	assert.Equal(t, tokens[8].Word, ".")
	assert.False(t, tokens[8].IsParameter())
	assert.Equal(t, tokens[9].Word, "+")
	assert.False(t, tokens[9].IsParameter())
	assert.Equal(t, tokens[10].Word, "|")
	assert.False(t, tokens[10].IsParameter())
	assert.Equal(t, tokens[11].Word, "^")
	assert.False(t, tokens[11].IsParameter())
	assert.Equal(t, tokens[12].Word, "$")
	assert.False(t, tokens[12].IsParameter())
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

	properties, isMatch = NewCommand("abc").Match(".abc.")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("help").Match("helpful")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("search all").Match("search all")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("search all").Match("search     all")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("search all").Match("search for all")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("search all").Match("searching all")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("search all").Match("searchall")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("help").Match("Could you help me?")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("help me").Match("Could you help me?")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("help me").Match("please help me")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("echo <text>").Match("echo")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)

	properties, isMatch = NewCommand("echo <text>").Match("echo.")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("echo <text>").Match("echoing")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

	properties, isMatch = NewCommand("echo <text>").Match("echo hey")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("text", ""), "hey")

	properties, isMatch = NewCommand("echo <text>").Match("echo hello world")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("text", ""), "hello world")

	properties, isMatch = NewCommand("echo <text>").Match("echoing hey")
	assert.False(t, isMatch)
	assert.Nil(t, properties)

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

	properties, isMatch = NewCommand("repeat <text> <number>").Match("repeat hello world 10")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("text", ""), "hello world")
	assert.Equal(t, properties.IntegerParam("number", 0), 10)

	properties, isMatch = NewCommand("math {operation} <numbers>").Match("math + 2 10 56")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("operation", ""), "+")
	assert.Equal(t, properties.StringParam("numbers", ""), "2 10 56")

	properties, isMatch = NewCommand("math {number} {operation} {number2}").Match("math 2 + 10 56")
	assert.True(t, isMatch)
	assert.Equal(t, properties.StringParam("operation", ""), "+")
	assert.Equal(t, properties.StringParam("number", ""), "2")
	assert.Equal(t, properties.StringParam("number2", ""), "10")

	properties, isMatch = NewCommand("calculate <number1> plus <number2>").Match("calculate 10 plus 5")
	assert.True(t, isMatch)
	assert.Equal(t, properties.IntegerParam("number1", 0), 10)
	assert.Equal(t, properties.IntegerParam("number2", 0), 5)

	properties, isMatch = NewCommand("<number1> + <number2>").Match("10 + 5")
	assert.True(t, isMatch)
	assert.Equal(t, properties.IntegerParam("number1", 0), 10)
	assert.Equal(t, properties.IntegerParam("number2", 0), 5)

	properties, isMatch = NewCommand("<number1> + <number2>").Match("+")
	assert.True(t, isMatch)
	assert.Equal(t, properties.IntegerParam("number1", 0), 0)
	assert.Equal(t, properties.IntegerParam("number2", 0), 0)

	properties, isMatch = NewCommand("\\ ( ) { } [ ] ? . + | ^ $").Match("\\ ( ) { } [ ] ? . + | ^ $")
	assert.True(t, isMatch)
	assert.NotNil(t, properties)
}

func TestNewCommand(t *testing.T) {
	tests := []struct {
		name           string
		in             string
		wantTokens     int
		wantExpresions int
	}{
		{"simple command", "ping", 1, 2},
		{"command and parameter", "say <input>", 2, 3},
		{"no command", "", 0, 0},
		{"all tokens are parameter", "<a> <123> <a123> <a-123> <a.123>", 5, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommand(tt.in)
			if got := len(cmd.tokens); got != tt.wantTokens {
				t.Errorf("got tokens %v, want %v", got, tt.wantTokens)
			}
			if got := len(cmd.expressions); got != tt.wantExpresions {
				t.Errorf("got expressions %v, want %v", got, tt.wantExpresions)
			}
			// Check no panic
			_, _ = cmd.Match("")
		})
	}
}
