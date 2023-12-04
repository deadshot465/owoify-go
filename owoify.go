// Package owoify_go is a Go port of mohan-cao's owoify-js library (https://github.com/mohan-cao/owoify-js).
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
	"github.com/deadshot465/owoify-go/v2/structures/word"
	"github.com/deadshot465/owoify-go/v2/util"
	"github.com/deadshot465/owoify-go/v2/util/mappings"
	"github.com/deadshot465/owoify-go/v2/util/presets"
	"regexp"
	"strings"
	"sync"
)

var (
	wordRegex  *regexp.Regexp
	spaceRegex *regexp.Regexp
	once       sync.Once
)

type Owoness int

const (
	Owo Owoness = iota
	Uwu
	Uvu
)

func initialize() {
	once.Do(func() {
		wordRegex = regexp.MustCompile("[^\\s]+")
		spaceRegex = regexp.MustCompile("\\s+")
		mappings.Initialize()
		presets.Initialize()
	})
}

// Owoify is the main entry point of Owoify. Available owoness levels are (from lowest to highest): Owo, Uwu, Uvu.
// Panics when passing an invalid owoness level.
func Owoify(source string, level Owoness) string {
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
		for _, v := range presets.SpecificWordMappingList {
			w = v(w)
		}
		switch level {
		case Owo:
			for _, v := range presets.OwoWordMappingList {
				w = v(w)
			}
		case Uwu:
			for _, v := range presets.UwuWordMappingList {
				w = v(w)
			}
			for _, v := range presets.OwoWordMappingList {
				w = v(w)
			}
		case Uvu:
			for _, v := range presets.UvuWordMappingList {
				w = v(w)
			}
			for _, v := range presets.UwuWordMappingList {
				w = v(w)
			}
			for _, v := range presets.OwoWordMappingList {
				w = v(w)
			}
		default:
			panic("The provided owoness level is not correct.")
		}
	}

	result := util.InterleaveSlices(words, spaces)
	var resultStrings []string
	for _, v := range result {
		resultStrings = append(resultStrings, v.ToString())
	}
	return strings.Join(resultStrings, "")
}

// Uwuify owoifies the given source string using Uwu owoness level.
func Uwuify(source string) string {
	return Owoify(source, Uwu)
}

// Uvuify owoifies the given source string using Uvu owoness level.
func Uvuify(source string) string {
	return Owoify(source, Uvu)
}
