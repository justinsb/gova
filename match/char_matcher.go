package match

type CharMatcher struct {
	matcher singleCharMatcher
}

type singleCharMatcher interface {
	Matches(c rune) bool
}

type simpleSingleCharMatcher struct {
	matches map[rune]bool
}

func Is(c rune) *CharMatcher {
	s := string([]rune{c})
	return AnyOf(s)
}

func AnyOf(s string) *CharMatcher {
	singleMatcher := &simpleSingleCharMatcher{}
	singleMatcher.matches = map[rune]bool{}
	for _, c := range s {
		singleMatcher.matches[c] = true
	}

	matcher := &CharMatcher{}
	matcher.matcher = singleMatcher
	return matcher
}

func (self *CharMatcher) MatchesAllOf(s string) bool {
	for _, c := range s {
		match := self.matcher.Matches(c)
		if !match {
			return false
		}
	}
	return true
}

func (self *simpleSingleCharMatcher) Matches(c rune) bool {
	match, found := self.matches[c]
	return found && match
}
