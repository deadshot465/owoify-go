package util

import "github.com/deadshot465/owoify-go/structures/word"

func InterleaveArrays(a []*word.Word, b []*word.Word) []*word.Word {
	var arr []*word.Word
	observed := a
	other := b
	var temp []*word.Word
	
	for len(observed) > 0 {
		arr = append(arr, observed[0])
		observed = append(observed[:0], observed[1:]...)
		temp = observed
		observed = other
		other = temp
	}

	if len(other) > 0 {
		arr = append(arr, other...)
	}
	return arr
}