package presets

import (
	"github.com/deadshot465/owoify-go/v2/structures/word"
	"github.com/deadshot465/owoify-go/v2/util/mappings"
	"sync"
)

var (
	SpecificWordMappingList []func(*word.Word) *word.Word
	UvuWordMappingList      []func(*word.Word) *word.Word
	UwuWordMappingList      []func(*word.Word) *word.Word
	OwoWordMappingList      []func(*word.Word) *word.Word
	once                    sync.Once
)

func Initialize() {
	once.Do(func() {
		SpecificWordMappingList = []func(*word.Word) *word.Word{
			mappings.MapFucToFwuc, mappings.MapMomToMwom, mappings.MapTimeToTim,
			mappings.MapMeToMwe, mappings.MapOverToOwor,
			mappings.MapOveToUv, mappings.MapHahaToHeheXd, mappings.MapTheToTeh,
			mappings.MapYouToU, mappings.MapReadToWead, mappings.MapWorseToWose,
			mappings.MapGreatToGwate, mappings.MapAviatToAwiat, mappings.MapDedicatToDeditat,
			mappings.MapRememberToRember, mappings.MapWhenToWen, mappings.MapFrightenedToFrigten,
			mappings.MapMemeToMem, mappings.MapFeelToFell}

		UvuWordMappingList = []func(*word.Word) *word.Word{
			mappings.MapOToOwO, mappings.MapEwToUwu, mappings.MapHeyToHay,
			mappings.MapDeadToDed, mappings.MapNVowelTToNd,
		}

		UwuWordMappingList = []func(*word.Word) *word.Word{
			mappings.MapBracketToStarTrails, mappings.MapPeriodCommaExclamationSemicolonToKaomojis,
			mappings.MapThatToDat, mappings.MapThToF, mappings.MapLeToWal, mappings.MapVeToWe,
			mappings.MapRyToWwy, mappings.MapROrLToW,
		}

		OwoWordMappingList = []func(*word.Word) *word.Word{
			mappings.MapNVowelToNy,
			mappings.MapLlToWw, mappings.MapVowelOrRExceptOLToWl, mappings.MapOldToOwld,
			mappings.MapOlToOwl, mappings.MapLOrROToWo, mappings.MapSpecificConsonantsOToLetterAndWo,
			mappings.MapVOrWLeToWal, mappings.MapFiToFwi, mappings.MapVerToWer, mappings.MapPoiToPwoi,
			mappings.MapSpecificConsonantsLeToLetterAndWal, mappings.MapConsonantsRToConsonantW,
			mappings.MapLyToWy, mappings.MapPleToPwe, mappings.MapNrToNw, mappings.MapMemToMwem,
			mappings.UnmapNywoToNyo}
	})
}
