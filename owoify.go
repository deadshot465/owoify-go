// Package owoify-go is a Go port of mohan-cao's owoify-js library (https://github.com/mohan-cao/owoify-js).
//
// Example and README can be found on owoify-go's GitHub repo: https://github.com/deadshot465/owoify-go
//
// Basically, you would and should only call owoify-go's Owoify function, and pass in the source string to be owoified and the desired owoness.
//
//		owoify_go.Owoify("Hello, World!", "uvu")
//
// The returned string will be the owoified source string.
package owoify_go

import (
	"github.com/deadshot465/owoify-go/structures/word"
	"github.com/deadshot465/owoify-go/util"
	"github.com/deadshot465/owoify-go/util/mappings"
	"github.com/deadshot465/owoify-go/util/presets"
	"regexp"
	"strings"
	"sync"
)

var (
	wordRegex *regexp.Regexp
	spaceRegex *regexp.Regexp
	once sync.Once
)

func initialize() {
	once.Do(func() {
		wordRegex = regexp.MustCompile("[^\\s]+")
		spaceRegex = regexp.MustCompile("\\s+")
		mappings.Initialize()
		presets.Initialize()
	})
}

// The main entry point of Owoify.
func Owoify(source string, level string) string {
	_level := strings.TrimSpace(strings.ToLower(level))
	initialize()
	wordMatches := wordRegex.FindAllStringSubmatch(source, -1)
	spaceMatches := spaceRegex.FindAllStringSubmatch(source, -1)
	var words []*word.Word
	var spaces []*word.Word
	for _, v := range wordMatches {
		words = append(words, word.New(v[0]))
	}
	for _, v := range spaceMatches {
		spaces = append(spaces, word.New(v[0]))
	}

	for _, w := range words {
		for _, v := range presets.SPECIFIC_WORD_MAPPING_LIST {
			w = v(w)
		}
		switch _level {
		case "owo":
			for _, v := range presets.OWO_WORD_MAPPING_LIST {
				w = v(w)
			}
		case "uwu":
			for _, v := range presets.UWU_WORD_MAPPING_LIST {
				w = v(w)
			}
			for _, v := range presets.OWO_WORD_MAPPING_LIST {
				w = v(w)
			}
		case "uvu":
			for _, v := range presets.UVU_WORD_MAPPING_LIST {
				w = v(w)
			}
			for _, v := range presets.UWU_WORD_MAPPING_LIST {
				w = v(w)
			}
			for _, v := range presets.OWO_WORD_MAPPING_LIST {
				w = v(w)
			}
		}
	}

	result := util.InterleaveArrays(words, spaces)
	var resultStrings []string
	for _, v := range result {
		resultStrings = append(resultStrings, v.ToString())
	}
	return strings.Join(resultStrings, "")
}