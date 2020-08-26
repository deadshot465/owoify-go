package presets

import (
	"github.com/deadshot465/owoify-go/structures/word"
	"github.com/deadshot465/owoify-go/util/mappings"
	"sync"
)

var (
	SPECIFIC_WORD_MAPPING_LIST []func(*word.Word) *word.Word
	UVU_WORD_MAPPING_LIST      []func(*word.Word) *word.Word
	UWU_WORD_MAPPING_LIST      []func(*word.Word) *word.Word
	OWO_WORD_MAPPING_LIST      []func(*word.Word) *word.Word
	once                       sync.Once
)

func Initialize() {
	once.Do(func() {
		SPECIFIC_WORD_MAPPING_LIST = []func(*word.Word) *word.Word {
			mappings.MapFucToFwuc, mappings.MapMomToMwom, mappings.MapTimeToTim,
			mappings.MapMeToMwe, mappings.MapNVowelToNy, mappings.MapOverToOwor,
			mappings.MapOveToUv, mappings.MapHahaToHeheXd, mappings.MapTheToTeh,
			mappings.MapYouToU, mappings.MapReadToWead, mappings.MapWorseToWose }

		UVU_WORD_MAPPING_LIST = []func(*word.Word) *word.Word {
			mappings.MapOToOwO, mappings.MapEwToUwu, mappings.MapHeyToHay,
			mappings.MapDeadToDed, mappings.MapNVowelTToNd,
		}

		UWU_WORD_MAPPING_LIST = []func(*word.Word) *word.Word {
			mappings.MapBracketToStarTrails, mappings.MapPeriodCommaExclamationSemicolonToKaomojis,
			mappings.MapThatToDat, mappings.MapThToF, mappings.MapLeToWal, mappings.MapVeToWe,
			mappings.MapRyToWwy, mappings.MapROrLToW,
		}

		OWO_WORD_MAPPING_LIST = []func(*word.Word) *word.Word {
			mappings.MapLlToWw, mappings.MapVowelOrRExceptOLToWl, mappings.MapOldToOwld,
			mappings.MapOlToOwl, mappings.MapLOrROToWo, mappings.MapSpecificConsonantsOToLetterAndWo,
			mappings.MapVOrWLeToWal, mappings.MapFiToFwi, mappings.MapPoiToPwoi,
			mappings.MapSpecificConsonantsLeToLetterAndWal, mappings.MapConsonantsRToConsonantW,
			mappings.MapLyToWy, mappings.MapPleToPwe, mappings.MapNrToNw,
		}
	})
}