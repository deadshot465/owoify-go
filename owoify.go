package owoify_go

import (
	"owoify-go/structures/word"
	"owoify-go/util"
	"owoify-go/util/mappings"
	"owoify-go/util/presets"
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

func Owoify(source string, level string) string {
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
		switch level {
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