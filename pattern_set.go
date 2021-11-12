package commander

type patternSet struct {
	PreCommandPattern  string
	PostCommandPattern string
	InputPattern       string
	LazyInputPattern   string
	SpacePattern       string
	MatchOffset        int
	ExactMatch         bool
}

var defaultCommandPatternSet = patternSet{
	PreCommandPattern:  "(\\s|^)",
	PostCommandPattern: "(\\s|$)",
	InputPattern:       "(.+)",
	LazyInputPattern:   "(.+?)",
	SpacePattern:       "\\s+",
	MatchOffset:        2,
	ExactMatch:         false,
}

var exactMatchPatternSet = patternSet{
	PreCommandPattern:  "(^(\\s+)?)",
	PostCommandPattern: "((\\s+)?$)",
	InputPattern:       "(\\S.*)",
	LazyInputPattern:   "(\\S.*?)",
	SpacePattern:       "\\s+",
	MatchOffset:        3,
	ExactMatch:         true,
}
