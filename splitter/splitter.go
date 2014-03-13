package splitter

import (
	"strings"
)

type Splitter struct {
	separator        string
	trimResults      bool
	omitEmptyStrings bool
}

func On(separator string) Splitter {
	self := Splitter{}
	self.separator = separator
	return self
}

func (self Splitter) TrimResults() Splitter {
	var ret Splitter
	ret = self
	ret.trimResults = true
	return ret
}

func (self Splitter) OmitEmptyStrings() Splitter {
	var ret Splitter = self
	ret.omitEmptyStrings = true
	return ret
}

func (self Splitter) Split(s string) []string {
	tokens := strings.Split(s, self.separator)

	ret := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if self.trimResults {
			token = strings.TrimSpace(token)
		}
		if self.omitEmptyStrings && token == "" {
			continue
		}
		ret = append(ret, token)
	}

	return ret
}
