package mappings

import (
	"github.com/deadshot465/owoify-go/v2/structures/word"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	OToOwo                                          *regexp.Regexp
	EwToUwu                                         *regexp.Regexp
	HeyToHay                                        *regexp.Regexp
	DeadToDedUpper                                  *regexp.Regexp
	DeadToDedLower                                  *regexp.Regexp
	NVowelTToNd                                     *regexp.Regexp
	ReadToWeadUpper                                 *regexp.Regexp
	ReadToWeadLower                                 *regexp.Regexp
	BracketsToStartrailsFore                        *regexp.Regexp
	BracketsToStartrailsRear                        *regexp.Regexp
	PeriodCommaExclamationSemicolonToKaomojisFirst  *regexp.Regexp
	PeriodCommaExclamationSemicolonToKaomojisSecond *regexp.Regexp
	ThatToDatUpper                                  *regexp.Regexp
	ThatToDatLower                                  *regexp.Regexp
	ThToFUpper                                      *regexp.Regexp
	ThToFLower                                      *regexp.Regexp
	LeToWal                                         *regexp.Regexp
	VeToWeUpper                                     *regexp.Regexp
	VeToWeLower                                     *regexp.Regexp
	RyToWwy                                         *regexp.Regexp
	RorlToWUpper                                    *regexp.Regexp
	RorlToWLower                                    *regexp.Regexp
	LlToWw                                          *regexp.Regexp
	VowelOrRExceptOLToWlUpper                       *regexp.Regexp
	VowelOrRExceptOLToWlLower                       *regexp.Regexp
	OldToOwldUpper                                  *regexp.Regexp
	OldToOwldLower                                  *regexp.Regexp
	OlToOwlUpper                                    *regexp.Regexp
	OlToOwlLower                                    *regexp.Regexp
	LorrOToWoUpper                                  *regexp.Regexp
	LorrOToWoLower                                  *regexp.Regexp
	SpecificConsonantsOToLetterAndWoUpper           *regexp.Regexp
	SpecificConsonantsOToLetterAndWoLower           *regexp.Regexp
	VorwLeToWal                                     *regexp.Regexp
	FiToFwiUpper                                    *regexp.Regexp
	FiToFwiLower                                    *regexp.Regexp
	VerToWer                                        *regexp.Regexp
	PoiToPwoi                                       *regexp.Regexp
	SpecificConsonantsLeToLetterAndWal              *regexp.Regexp
	ConsonantRToConsonantW                          *regexp.Regexp
	LyToWyUpper                                     *regexp.Regexp
	LyToWyLower                                     *regexp.Regexp
	PleToPwe                                        *regexp.Regexp
	NrToNwUpper                                     *regexp.Regexp
	NrToNwLower                                     *regexp.Regexp
	FucToFwuc                                       *regexp.Regexp
	MomToMwom                                       *regexp.Regexp
	MeToMwe                                         *regexp.Regexp
	NVowelToNyFirst                                 *regexp.Regexp
	NVowelToNySecond                                *regexp.Regexp
	NVowelToNyThird                                 *regexp.Regexp
	OveToUvUpper                                    *regexp.Regexp
	OveToUvLower                                    *regexp.Regexp
	HahaToHeheXd                                    *regexp.Regexp
	TheToTeh                                        *regexp.Regexp
	YouToUUpper                                     *regexp.Regexp
	YouToULower                                     *regexp.Regexp
	TimeToTim                                       *regexp.Regexp
	OverToOwor                                      *regexp.Regexp
	WorseToWose                                     *regexp.Regexp
	FACES                                           []string
	once                                            sync.Once
)

func Initialize() {
	once.Do(func() {
		OToOwo = regexp.MustCompile("o")
		EwToUwu = regexp.MustCompile("ew")
		HeyToHay = regexp.MustCompile("([Hh])ey")
		DeadToDedUpper = regexp.MustCompile("Dead")
		DeadToDedLower = regexp.MustCompile("dead")
		NVowelTToNd = regexp.MustCompile("n[aeiou]*t")
		ReadToWeadUpper = regexp.MustCompile("Read")
		ReadToWeadLower = regexp.MustCompile("read")
		BracketsToStartrailsFore = regexp.MustCompile("[({<]")
		BracketsToStartrailsRear = regexp.MustCompile("[)}>]")
		//PERIOD_COMMA_EXCLAMATION_SEMICOLON_TO_KAOMOJIS_FIRST = regexp.MustCompile("[.,](?![0-9])")
		PeriodCommaExclamationSemicolonToKaomojisFirst = regexp.MustCompile("[.,]")
		PeriodCommaExclamationSemicolonToKaomojisSecond = regexp.MustCompile("[!;]+")
		ThatToDatUpper = regexp.MustCompile("That")
		ThatToDatLower = regexp.MustCompile("that")
		//TH_TO_F_UPPER = regexp.MustCompile("TH(?!E)")
		ThToFUpper = regexp.MustCompile("TH")
		//TH_TO_F_LOWER = regexp.MustCompile("[Tt]h(?![Ee])")
		ThToFLower = regexp.MustCompile("[Tt]h")
		LeToWal = regexp.MustCompile("le$")
		VeToWeUpper = regexp.MustCompile("Ve")
		VeToWeLower = regexp.MustCompile("ve")
		RyToWwy = regexp.MustCompile("ry")
		RorlToWUpper = regexp.MustCompile("(?:R|L)")
		RorlToWLower = regexp.MustCompile("(?:r|l)")
		LlToWw = regexp.MustCompile("ll")
		VowelOrRExceptOLToWlUpper = regexp.MustCompile("[AEIUR]([lL])$")
		VowelOrRExceptOLToWlLower = regexp.MustCompile("[aeiur]l$")
		OldToOwldUpper = regexp.MustCompile("OLD")
		OldToOwldLower = regexp.MustCompile("([Oo])ld")
		OlToOwlUpper = regexp.MustCompile("OL")
		OlToOwlLower = regexp.MustCompile("([Oo])l")
		LorrOToWoUpper = regexp.MustCompile("[LR]([oO])")
		LorrOToWoLower = regexp.MustCompile("[lr]o")
		SpecificConsonantsOToLetterAndWoUpper = regexp.MustCompile("([BCDFGHJKMNPQSTXYZ])([oO])")
		SpecificConsonantsOToLetterAndWoLower = regexp.MustCompile("([bcdfghjkmnpqstxyz])o")
		VorwLeToWal = regexp.MustCompile("[vw]le")
		FiToFwiUpper = regexp.MustCompile("FI")
		FiToFwiLower = regexp.MustCompile("([Ff])i")
		VerToWer = regexp.MustCompile("([Vv])er")
		PoiToPwoi = regexp.MustCompile("([Pp])oi")
		SpecificConsonantsLeToLetterAndWal = regexp.MustCompile("([DdFfGgHhJjPpQqRrSsTtXxYyZz])le$")
		ConsonantRToConsonantW = regexp.MustCompile("([BbCcDdFfGgKkPpQqSsTtWwXxZz])r")
		LyToWyUpper = regexp.MustCompile("Ly")
		LyToWyLower = regexp.MustCompile("ly")
		PleToPwe = regexp.MustCompile("([Pp])le")
		NrToNwUpper = regexp.MustCompile("NR")
		NrToNwLower = regexp.MustCompile("nr")
		FucToFwuc = regexp.MustCompile("([Ff])uc")
		MomToMwom = regexp.MustCompile("([Mm])om")
		MeToMwe = regexp.MustCompile("([Mm])e")
		NVowelToNyFirst = regexp.MustCompile("n([aeiou])")
		NVowelToNySecond = regexp.MustCompile("N([aeiou])")
		NVowelToNyThird = regexp.MustCompile("N([AEIOU])")
		OveToUvUpper = regexp.MustCompile("OVE")
		OveToUvLower = regexp.MustCompile("ove")
		HahaToHeheXd = regexp.MustCompile("\\b(ha|hah|heh|hehe)+\\b")
		TheToTeh = regexp.MustCompile("\\b([Tt])he\\b")
		YouToUUpper = regexp.MustCompile("\\bYou\\b")
		YouToULower = regexp.MustCompile("\\byou\\b")
		TimeToTim = regexp.MustCompile("\\b([Tt])ime\\b")
		OverToOwor = regexp.MustCompile("([Oo])ver")
		WorseToWose = regexp.MustCompile("([Ww])orse")
		FACES = []string{"(・`ω´・)", ";;w;;", "owo", "UwU", ">w<", "^w^", "(* ^ ω ^)", "(⌒ω⌒)", "ヽ(*・ω・)ﾉ", "(o´∀`o)", "(o･ω･o)", "＼(＾▽＾)／", "(*^ω^)", "(◕‿◕✿)", "(◕ᴥ◕)", "ʕ•ᴥ•ʔ", "ʕ￫ᴥ￩ʔ", "(*^.^*)", "(｡♥‿♥｡)", "OwO", "uwu", "uvu", "UvU", "(*￣з￣)", "(つ✧ω✧)つ", "(/ =ω=)/", "(╯°□°）╯︵ ┻━┻", "┬─┬ ノ( ゜-゜ノ)", "¯\\_(ツ)_/¯"}
	})
}

func MapOToOwO(input *word.Word) *word.Word {
	var replacement string
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) > 0 {
		replacement = "owo"
	} else {
		replacement = "o"
	}

	return input.Replace(OToOwo, replacement)

}

func MapEwToUwu(input *word.Word) *word.Word {
	return input.Replace(EwToUwu, "uwu")
}

func MapHeyToHay(input *word.Word) *word.Word {
	return input.Replace(HeyToHay, "${1}ay")
}

func MapDeadToDed(input *word.Word) *word.Word {
	return input.Replace(DeadToDedUpper, "Ded").Replace(DeadToDedLower, "ded")
}

func MapNVowelTToNd(input *word.Word) *word.Word {
	return input.Replace(NVowelTToNd, "nd")
}

func MapReadToWead(input *word.Word) *word.Word {
	return input.Replace(ReadToWeadUpper, "Wead").Replace(ReadToWeadLower, "wead")
}

func MapBracketToStarTrails(input *word.Word) *word.Word {
	return input.Replace(BracketsToStartrailsFore, "｡･:*:･ﾟ★,｡･:*:･ﾟ☆").Replace(BracketsToStartrailsRear, "☆ﾟ･:*:･｡,★ﾟ･:*:･｡")
}

func MapPeriodCommaExclamationSemicolonToKaomojis(input *word.Word) *word.Word {
	return input.ReplaceWithFuncSingle(PeriodCommaExclamationSemicolonToKaomojisFirst, func() string {
		return FACES[rand.Intn(len(FACES))]
	}).ReplaceWithFuncSingle(PeriodCommaExclamationSemicolonToKaomojisSecond, func() string {
		return FACES[rand.Intn(len(FACES))]
	})
}

func MapThatToDat(input *word.Word) *word.Word {
	return input.Replace(ThatToDatLower, "dat").Replace(ThatToDatUpper, "Dat")
}

func MapThToF(input *word.Word) *word.Word {
	return input.Replace(ThToFLower, "f").Replace(ThToFUpper, "F")
}

func MapLeToWal(input *word.Word) *word.Word {
	return input.Replace(LeToWal, "wal")
}

func MapVeToWe(input *word.Word) *word.Word {
	return input.Replace(VeToWeLower, "we").Replace(VeToWeUpper, "We")
}

func MapRyToWwy(input *word.Word) *word.Word {
	return input.Replace(RyToWwy, "wwy")
}

func MapROrLToW(input *word.Word) *word.Word {
	return input.Replace(RorlToWLower, "w").Replace(RorlToWUpper, "W")
}

func MapLlToWw(input *word.Word) *word.Word {
	return input.Replace(LlToWw, "ww")
}

func MapVowelOrRExceptOLToWl(input *word.Word) *word.Word {
	return input.Replace(VowelOrRExceptOLToWlLower, "wl").Replace(VowelOrRExceptOLToWlUpper, "W${1}")
}

func MapOldToOwld(input *word.Word) *word.Word {
	return input.Replace(OldToOwldLower, "${1}wld").Replace(OldToOwldUpper, "OWLD")
}

func MapOlToOwl(input *word.Word) *word.Word {
	return input.Replace(OlToOwlLower, "${1}wl").Replace(OlToOwlUpper, "OWL")
}

func MapLOrROToWo(input *word.Word) *word.Word {
	return input.Replace(LorrOToWoLower, "wo").Replace(LorrOToWoUpper, "W${1}")
}

func MapSpecificConsonantsOToLetterAndWo(input *word.Word) *word.Word {
	return input.Replace(SpecificConsonantsOToLetterAndWoLower, "${1}wo").ReplaceWithFuncMultiple(SpecificConsonantsOToLetterAndWoUpper, func(str1 string, str2 string) string {
		msg := str1
		if strings.ToUpper(str2) == str2 {
			msg += "W"
		} else {
			msg += "w"
		}
		msg += str2
		return msg
	})
}

func MapVOrWLeToWal(input *word.Word) *word.Word {
	return input.Replace(VorwLeToWal, "wal")
}

func MapFiToFwi(input *word.Word) *word.Word {
	return input.Replace(FiToFwiLower, "${1}wi").Replace(FiToFwiUpper, "FWI")
}

func MapVerToWer(input *word.Word) *word.Word {
	return input.Replace(VerToWer, "wer")
}

func MapPoiToPwoi(input *word.Word) *word.Word {
	return input.Replace(PoiToPwoi, "${1}woi")
}

func MapSpecificConsonantsLeToLetterAndWal(input *word.Word) *word.Word {
	return input.Replace(SpecificConsonantsLeToLetterAndWal, "${1}wal")
}

func MapConsonantsRToConsonantW(input *word.Word) *word.Word {
	return input.Replace(ConsonantRToConsonantW, "${1}w")
}

func MapLyToWy(input *word.Word) *word.Word {
	return input.Replace(LyToWyLower, "wy").Replace(LyToWyUpper, "Wy")
}

func MapPleToPwe(input *word.Word) *word.Word {
	return input.Replace(PleToPwe, "${1}we")
}

func MapNrToNw(input *word.Word) *word.Word {
	return input.Replace(NrToNwLower, "nw").Replace(NrToNwUpper, "NW")
}

func MapFucToFwuc(input *word.Word) *word.Word {
	return input.Replace(FucToFwuc, "${1}wuc")
}

func MapMomToMwom(input *word.Word) *word.Word {
	return input.Replace(MomToMwom, "${1}wom")
}

func MapMeToMwe(input *word.Word) *word.Word {
	return input.Replace(MeToMwe, "${1}we")
}

func MapNVowelToNy(input *word.Word) *word.Word {
	return input.Replace(NVowelToNyFirst, "ny${1}").Replace(NVowelToNySecond, "Ny${1}").Replace(NVowelToNyThird, "NY${1}")
}

func MapOveToUv(input *word.Word) *word.Word {
	return input.Replace(OveToUvLower, "uv").Replace(OveToUvUpper, "UV")
}

func MapHahaToHeheXd(input *word.Word) *word.Word {
	return input.Replace(HahaToHeheXd, "hehe xD")
}

func MapTheToTeh(input *word.Word) *word.Word {
	return input.Replace(TheToTeh, "${1}eh")
}

func MapYouToU(input *word.Word) *word.Word {
	return input.Replace(YouToUUpper, "U").Replace(YouToULower, "u")
}

func MapTimeToTim(input *word.Word) *word.Word {
	return input.Replace(TimeToTim, "${1}im")
}

func MapOverToOwor(input *word.Word) *word.Word {
	return input.Replace(OverToOwor, "${1}wor")
}

func MapWorseToWose(input *word.Word) *word.Word {
	return input.Replace(WorseToWose, "${1}ose")
}
