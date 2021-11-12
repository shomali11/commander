package commander

import (
	"regexp"
	"strings"

	"github.com/shomali11/proper"
)

const (
	escapeCharacter      = "\\"
	ignoreCase           = "(?i)"
	parameterPattern     = "<\\S+>"
	lazyParameterPattern = "<\\S+\\?>"
)

const (
	notParameter = iota
	greedyParameter
	lazyParameter
)

var (
	regexCharacters = []string{"\\", "(", ")", "{", "}", "[", "]", "?", ".", "+", "|", "^", "$"}
)

type Options struct {
	// If exact match is set, matching requires the command to begin at the start of input, and contain no extra output
	PatternSet *patternSet
}

type Option func(opts *Options)

// WithExactMatch requires that the command contains all parameters, and has no additional starting/ending text.
// ex: NewCommand("help", WithExactMatch(true)) -> "help me" doesn't match, while "  help  " will.
func WithExactMatch(m bool) Option {
	return func(opts *Options) {
		if m {
			opts.PatternSet = &exactMatchPatternSet
		} else {
			opts.PatternSet = &defaultCommandPatternSet
		}
	}
}

// NewCommand creates a new Command object from the format passed in
func NewCommand(format string, options ...Option) *Command {
	opts := createsOptions(options)
	tokens := tokenize(format)
	expressions := generate(tokens, opts)
	return &Command{tokens: tokens, expressions: expressions, opts: opts}
}

// Token represents the Token object
type Token struct {
	Word string
	Type int
}

func (t Token) IsParameter() bool {
	return t.Type != notParameter
}

// Command represents the Command object
type Command struct {
	opts        *Options
	tokens      []*Token
	expressions []*regexp.Regexp
}

// Match takes in the command and the text received, attempts to find the pattern and extract the parameters
func (c *Command) Match(text string) (*proper.Properties, bool) {
	if len(c.expressions) == 0 {
		return nil, false
	}

	matchOffset := c.opts.PatternSet.MatchOffset
	for _, expression := range c.expressions {
		matches := expression.FindStringSubmatch(text)
		if len(matches) == 0 {
			continue
		}

		values := matches[matchOffset : len(matches)-1]

		valueIndex := 0
		parameters := make(map[string]string)
		for i := 0; i < len(c.tokens) && valueIndex < len(values); i++ {
			token := c.tokens[i]
			if !token.IsParameter() {
				continue
			}

			parameters[token.Word] = values[valueIndex]
			valueIndex++
		}
		return proper.NewProperties(parameters), true
	}
	return nil, false
}

// Tokenize returns Command info as tokens
func (c *Command) Tokenize() []*Token {
	return c.tokens
}

func escape(text string) string {
	for _, character := range regexCharacters {
		text = strings.Replace(text, character, escapeCharacter+character, -1)
	}
	return text
}

func tokenize(format string) []*Token {
	parameterRegex := regexp.MustCompile(parameterPattern)
	lazyParameterRegex := regexp.MustCompile(lazyParameterPattern)
	words := strings.Fields(format)
	tokens := make([]*Token, len(words))
	for i, word := range words {
		switch {
		case lazyParameterRegex.MatchString(word):
			tokens[i] = &Token{Word: word[1 : len(word)-2], Type: lazyParameter}
		case parameterRegex.MatchString(word):
			tokens[i] = &Token{Word: word[1 : len(word)-1], Type: greedyParameter}
		default:
			tokens[i] = &Token{Word: word, Type: notParameter}
		}
	}
	return tokens
}

func generate(tokens []*Token, opts *Options) []*regexp.Regexp {
	regexps := []*regexp.Regexp{}
	if len(tokens) == 0 {
		return regexps
	}

	// Don't generate variations on exact match
	if opts.PatternSet.ExactMatch {
		regexps = append(regexps, compile(tokens, opts))
		return regexps
	}

	for index := len(tokens) - 1; index >= -1; index-- {
		regex := compile(create(tokens, index), opts)
		regexps = append(regexps, regex)
	}

	return regexps
}

func create(tokens []*Token, boundary int) []*Token {
	newTokens := []*Token{}
	for i := 0; i < len(tokens); i++ {
		if !tokens[i].IsParameter() || i <= boundary {
			newTokens = append(newTokens, tokens[i])
		}
	}
	return newTokens
}

func compile(tokens []*Token, opts *Options) *regexp.Regexp {
	if len(tokens) == 0 {
		return nil
	}
	patterns := opts.PatternSet
	pattern := patterns.PreCommandPattern + getInputPattern(tokens[0], patterns)
	for index := 1; index < len(tokens); index++ {
		currentToken := tokens[index]
		pattern += patterns.SpacePattern + getInputPattern(currentToken, patterns)
	}
	pattern += patterns.PostCommandPattern

	return regexp.MustCompile(ignoreCase + pattern)
}

func getInputPattern(token *Token, patterns *patternSet) string {
	switch token.Type {
	case lazyParameter:
		return patterns.LazyInputPattern
	case greedyParameter:
		return patterns.InputPattern
	default:
		return escape(token.Word)
	}
}

func createsOptions(options []Option) *Options {
	opts := &Options{
		PatternSet: &defaultCommandPatternSet,
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}
