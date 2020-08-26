package word

import (
	"regexp"
	"strings"
)

// The basic unit of owoify function.
type Word struct {
	word string
	replacedWords []string
}

func New(str string) *Word {
	return &Word{ str, []string{} }
}

// Replace all matched items with a string.
func (word *Word) Replace(searchValue *regexp.Regexp, replaceValue string) *Word {
	if word.searchValueContainsReplacedWords(searchValue, replaceValue) {
		return word
	}
	replacingWord := word.word
	if searchValue.MatchString(word.word) {
		replacingWord = searchValue.ReplaceAllString(word.word, replaceValue)
	}
	collection := searchValue.FindAllStringSubmatch(word.word, -1)
	var replacedWords []string
	if len(collection) > 1 {
		for _, v := range collection {
			replacedWords = append(replacedWords, strings.ReplaceAll(v[0], v[0], replaceValue))
		}
	}
	if replacingWord != word.word {
		for _, v := range replacedWords {
			word.replacedWords = append(word.replacedWords, v)
		}
		word.word = replacingWord
	}
	return word
}

// Replace all matched items by repeatedly invoking a function.
func (word *Word) ReplaceWithFuncSingle(searchValue *regexp.Regexp, f func() string) *Word {
	replaceValue := f()
	if word.searchValueContainsReplacedWords(searchValue, replaceValue) {
		return word
	}
	replacingWord := word.word
	if searchValue.MatchString(word.word) {
		matchItem := searchValue.FindStringSubmatch(word.word)[0]
		replacingWord = strings.ReplaceAll(word.word, matchItem, replaceValue)
	}
	collection := searchValue.FindAllStringSubmatch(word.word, -1)
	var replacedWords []string
	if len(collection) > 1 {
		for _, v := range collection {
			replacedWords = append(replacedWords, strings.ReplaceAll(v[0], v[0], replaceValue))
		}
	}
	if replacingWord != word.word {
		for _, v := range replacedWords {
			word.replacedWords = append(word.replacedWords, v)
		}
		word.word = replacingWord
	}
	return word
}

// Replace all matched items by repeatedly invoking a function.
func (word *Word) ReplaceWithFuncMultiple(searchValue *regexp.Regexp, f func(string, string) string) *Word {
	if !searchValue.MatchString(word.word) {
		return word
	}
	_word := word.word
	captures := searchValue.FindStringSubmatch(_word)
	replaceValue := f(captures[1], captures[2])
	if word.searchValueContainsReplacedWords(searchValue, replaceValue) {
		return word
	}
	replacingWord := strings.ReplaceAll(word.word, captures[0], replaceValue)
	collection := searchValue.FindAllStringSubmatch(word.word, -1)
	var replacedWords []string
	if len(collection) > 1 {
		for _, v := range collection {
			replacedWords = append(replacedWords, strings.ReplaceAll(v[0], v[0], replaceValue))
		}
	}
	if replacingWord != word.word {
		for _, v := range replacedWords {
			word.replacedWords = append(word.replacedWords, v)
		}
		word.word = replacingWord
	}
	return word
}

func (word *Word) ToString() string {
	return word.word
}

func (word *Word) searchValueContainsReplacedWords(searchValue *regexp.Regexp, replaceValue string) bool {
	for _, v := range word.replacedWords {
		if searchValue.MatchString(v) {
			matchResult := searchValue.FindStringSubmatch(v)[0]
			return strings.ReplaceAll(v, matchResult, replaceValue) == replaceValue
		}
	}
	return false
}