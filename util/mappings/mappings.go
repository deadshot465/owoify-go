package mappings

import (
	"github.com/deadshot465/owoify-go/structures/word"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	O_TO_OWO *regexp.Regexp
	EW_TO_UWU *regexp.Regexp
	HEY_TO_HAY *regexp.Regexp
	DEAD_TO_DED_UPPER *regexp.Regexp
	DEAD_TO_DED_LOWER *regexp.Regexp
	N_VOWEL_T_TO_ND *regexp.Regexp
	READ_TO_WEAD_UPPER *regexp.Regexp
	READ_TO_WEAD_LOWER *regexp.Regexp
	BRACKETS_TO_STARTRAILS_FORE *regexp.Regexp
	BRACKETS_TO_STARTRAILS_REAR *regexp.Regexp
	PERIOD_COMMA_EXCLAMATION_SEMICOLON_TO_KAOMOJIS_FIRST *regexp.Regexp
	PERIOD_COMMA_EXCLAMATION_SEMICOLON_TO_KAOMOJIS_SECOND *regexp.Regexp
	THAT_TO_DAT_UPPER *regexp.Regexp
	THAT_TO_DAT_LOWER *regexp.Regexp
	TH_TO_F_UPPER *regexp.Regexp
	TH_TO_F_LOWER *regexp.Regexp
	LE_TO_WAL *regexp.Regexp
	VE_TO_WE_UPPER *regexp.Regexp
	VE_TO_WE_LOWER *regexp.Regexp
	RY_TO_WWY *regexp.Regexp
	RORL_TO_W_UPPER *regexp.Regexp
	RORL_TO_W_LOWER *regexp.Regexp
	LL_TO_WW *regexp.Regexp
	VOWEL_OR_R_EXCEPT_O_L_TO_WL_UPPER *regexp.Regexp
	VOWEL_OR_R_EXCEPT_O_L_TO_WL_LOWER *regexp.Regexp
	OLD_TO_OWLD_UPPER *regexp.Regexp
	OLD_TO_OWLD_LOWER *regexp.Regexp
	OL_TO_OWL_UPPER *regexp.Regexp
	OL_TO_OWL_LOWER *regexp.Regexp
	LORR_O_TO_WO_UPPER *regexp.Regexp
	LORR_O_TO_WO_LOWER *regexp.Regexp
	SPECIFIC_CONSONANTS_O_TO_LETTER_AND_WO_UPPER *regexp.Regexp
	SPECIFIC_CONSONANTS_O_TO_LETTER_AND_WO_LOWER *regexp.Regexp
	VORW_LE_TO_WAL *regexp.Regexp
	FI_TO_FWI_UPPER *regexp.Regexp
	FI_TO_FWI_LOWER *regexp.Regexp
	VER_TO_WER *regexp.Regexp
	POI_TO_PWOI *regexp.Regexp
	SPECIFIC_CONSONANTS_LE_TO_LETTER_AND_WAL *regexp.Regexp
	CONSONANT_R_TO_CONSONANT_W *regexp.Regexp
	LY_TO_WY_UPPER *regexp.Regexp
	LY_TO_WY_LOWER *regexp.Regexp
	PLE_TO_PWE *regexp.Regexp
	NR_TO_NW_UPPER *regexp.Regexp
	NR_TO_NW_LOWER *regexp.Regexp
	FUC_TO_FWUC *regexp.Regexp
	MOM_TO_MWOM *regexp.Regexp
	ME_TO_MWE *regexp.Regexp
	N_VOWEL_TO_NY_FIRST *regexp.Regexp
	N_VOWEL_TO_NY_SECOND *regexp.Regexp
	N_VOWEL_TO_NY_THIRD *regexp.Regexp
	OVE_TO_UV_UPPER *regexp.Regexp
	OVE_TO_UV_LOWER *regexp.Regexp
	HAHA_TO_HEHE_XD *regexp.Regexp
	THE_TO_TEH *regexp.Regexp
	YOU_TO_U_UPPER *regexp.Regexp
	YOU_TO_U_LOWER *regexp.Regexp
	TIME_TO_TIM *regexp.Regexp
	OVER_TO_OWOR *regexp.Regexp
	WORSE_TO_WOSE *regexp.Regexp
	FACES []string
	once sync.Once
)

func Initialize() {
	once.Do(func() {
		O_TO_OWO = regexp.MustCompile("o")
		EW_TO_UWU = regexp.MustCompile("ew")
		HEY_TO_HAY = regexp.MustCompile("([Hh])ey")
		DEAD_TO_DED_UPPER = regexp.MustCompile("Dead")
		DEAD_TO_DED_LOWER = regexp.MustCompile("dead")
		N_VOWEL_T_TO_ND = regexp.MustCompile("n[aeiou]*t")
		READ_TO_WEAD_UPPER = regexp.MustCompile("Read")
		READ_TO_WEAD_LOWER = regexp.MustCompile("read")
		BRACKETS_TO_STARTRAILS_FORE = regexp.MustCompile("[({<]")
		BRACKETS_TO_STARTRAILS_REAR = regexp.MustCompile("[)}>]")
		//PERIOD_COMMA_EXCLAMATION_SEMICOLON_TO_KAOMOJIS_FIRST = regexp.MustCompile("[.,](?![0-9])")
		PERIOD_COMMA_EXCLAMATION_SEMICOLON_TO_KAOMOJIS_FIRST = regexp.MustCompile("[.,]")
		PERIOD_COMMA_EXCLAMATION_SEMICOLON_TO_KAOMOJIS_SECOND = regexp.MustCompile("[!;]+")
		THAT_TO_DAT_UPPER = regexp.MustCompile("That")
		THAT_TO_DAT_LOWER = regexp.MustCompile("that")
		//TH_TO_F_UPPER = regexp.MustCompile("TH(?!E)")
		TH_TO_F_UPPER = regexp.MustCompile("TH")
		//TH_TO_F_LOWER = regexp.MustCompile("[Tt]h(?![Ee])")
		TH_TO_F_LOWER = regexp.MustCompile("[Tt]h")
		LE_TO_WAL = regexp.MustCompile("le$")
		VE_TO_WE_UPPER = regexp.MustCompile("Ve")
		VE_TO_WE_LOWER = regexp.MustCompile("ve")
		RY_TO_WWY = regexp.MustCompile("ry")
		RORL_TO_W_UPPER = regexp.MustCompile("(?:R|L)")
		RORL_TO_W_LOWER = regexp.MustCompile("(?:r|l)")
		LL_TO_WW = regexp.MustCompile("ll")
		VOWEL_OR_R_EXCEPT_O_L_TO_WL_UPPER = regexp.MustCompile("[AEIUR]([lL])$")
		VOWEL_OR_R_EXCEPT_O_L_TO_WL_LOWER = regexp.MustCompile("[aeiur]l$")
		OLD_TO_OWLD_UPPER = regexp.MustCompile("OLD")
		OLD_TO_OWLD_LOWER = regexp.MustCompile("([Oo])ld")
		OL_TO_OWL_UPPER = regexp.MustCompile("OL")
		OL_TO_OWL_LOWER = regexp.MustCompile("([Oo])l")
		LORR_O_TO_WO_UPPER = regexp.MustCompile("[LR]([oO])")
		LORR_O_TO_WO_LOWER = regexp.MustCompile("[lr]o")
		SPECIFIC_CONSONANTS_O_TO_LETTER_AND_WO_UPPER = regexp.MustCompile("([BCDFGHJKMNPQSTXYZ])([oO])")
		SPECIFIC_CONSONANTS_O_TO_LETTER_AND_WO_LOWER = regexp.MustCompile("([bcdfghjkmnpqstxyz])o")
		VORW_LE_TO_WAL = regexp.MustCompile("[vw]le")
		FI_TO_FWI_UPPER = regexp.MustCompile("FI")
		FI_TO_FWI_LOWER = regexp.MustCompile("([Ff])i")
		VER_TO_WER = regexp.MustCompile("([Vv])er")
		POI_TO_PWOI = regexp.MustCompile("([Pp])oi")
		SPECIFIC_CONSONANTS_LE_TO_LETTER_AND_WAL = regexp.MustCompile("([DdFfGgHhJjPpQqRrSsTtXxYyZz])le$")
		CONSONANT_R_TO_CONSONANT_W = regexp.MustCompile("([BbCcDdFfGgKkPpQqSsTtWwXxZz])r")
		LY_TO_WY_UPPER = regexp.MustCompile("Ly")
		LY_TO_WY_LOWER = regexp.MustCompile("ly")
		PLE_TO_PWE = regexp.MustCompile("([Pp])le")
		NR_TO_NW_UPPER = regexp.MustCompile("NR")
		NR_TO_NW_LOWER = regexp.MustCompile("nr")
		FUC_TO_FWUC = regexp.MustCompile("([Ff])uc")
		MOM_TO_MWOM = regexp.MustCompile("([Mm])om")
		ME_TO_MWE = regexp.MustCompile("([Mm])e")
		N_VOWEL_TO_NY_FIRST = regexp.MustCompile("n([aeiou])")
		N_VOWEL_TO_NY_SECOND = regexp.MustCompile("N([aeiou])")
		N_VOWEL_TO_NY_THIRD = regexp.MustCompile("N([AEIOU])")
		OVE_TO_UV_UPPER = regexp.MustCompile("OVE")
		OVE_TO_UV_LOWER = regexp.MustCompile("ove")
		HAHA_TO_HEHE_XD = regexp.MustCompile("\\b(ha|hah|heh|hehe)+\\b")
		THE_TO_TEH = regexp.MustCompile("\\b([Tt])he\\b")
		YOU_TO_U_UPPER = regexp.MustCompile("\\bYou\\b")
		YOU_TO_U_LOWER = regexp.MustCompile("\\byou\\b")
		TIME_TO_TIM = regexp.MustCompile("\\b([Tt])ime\\b")
		OVER_TO_OWOR = regexp.MustCompile("([Oo])ver")
		WORSE_TO_WOSE = regexp.MustCompile("([Ww])orse")
		FACES = []string{ "(・`ω´・)", ";;w;;", "owo", "UwU", ">w<", "^w^", "(* ^ ω ^)", "(⌒ω⌒)", "ヽ(*・ω・)ﾉ", "(o´∀`o)", "(o･ω･o)", "＼(＾▽＾)／", "(*^ω^)", "(◕‿◕✿)", "(◕ᴥ◕)", "ʕ•ᴥ•ʔ", "ʕ￫ᴥ￩ʔ", "(*^.^*)", "(｡♥‿♥｡)", "OwO", "uwu", "uvu", "UvU", "(*￣з￣)", "(つ✧ω✧)つ", "(/ =ω=)/", "(╯°□°）╯︵ ┻━┻", "┬─┬ ノ( ゜-゜ノ)", "¯\\_(ツ)_/¯" }
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

	return input.Replace(O_TO_OWO, replacement)

}

func MapEwToUwu(input *word.Word) *word.Word {
	return input.Replace(EW_TO_UWU, "uwu")
}

func MapHeyToHay(input *word.Word) *word.Word {
	return input.Replace(HEY_TO_HAY, "${1}ay")
}

func MapDeadToDed(input *word.Word) *word.Word {
	return input.Replace(DEAD_TO_DED_UPPER, "Ded").Replace(DEAD_TO_DED_LOWER, "ded")
}

func MapNVowelTToNd(input *word.Word) *word.Word {
	return input.Replace(N_VOWEL_T_TO_ND, "nd")
}

func MapReadToWead(input *word.Word) *word.Word {
	return input.Replace(READ_TO_WEAD_UPPER, "Wead").Replace(READ_TO_WEAD_LOWER, "wead")
}

func MapBracketToStarTrails(input *word.Word) *word.Word {
	return input.Replace(BRACKETS_TO_STARTRAILS_FORE, "｡･:*:･ﾟ★,｡･:*:･ﾟ☆").Replace(BRACKETS_TO_STARTRAILS_REAR, "☆ﾟ･:*:･｡,★ﾟ･:*:･｡")
}

func MapPeriodCommaExclamationSemicolonToKaomojis(input *word.Word) *word.Word {
	return input.ReplaceWithFuncSingle(PERIOD_COMMA_EXCLAMATION_SEMICOLON_TO_KAOMOJIS_FIRST, func() string {
		return FACES[rand.Intn(len(FACES))]
	}).ReplaceWithFuncSingle(PERIOD_COMMA_EXCLAMATION_SEMICOLON_TO_KAOMOJIS_SECOND, func() string {
		return FACES[rand.Intn(len(FACES))]
	})
}

func MapThatToDat(input *word.Word) *word.Word {
	return input.Replace(THAT_TO_DAT_LOWER, "dat").Replace(THAT_TO_DAT_UPPER, "Dat")
}

func MapThToF(input *word.Word) *word.Word {
	return input.Replace(TH_TO_F_LOWER, "f").Replace(TH_TO_F_UPPER, "F")
}

func MapLeToWal(input *word.Word) *word.Word {
	return input.Replace(LE_TO_WAL, "wal")
}

func MapVeToWe(input *word.Word) *word.Word {
	return input.Replace(VE_TO_WE_LOWER, "we").Replace(VE_TO_WE_UPPER, "We")
}

func MapRyToWwy(input *word.Word) *word.Word {
	return input.Replace(RY_TO_WWY, "wwy")
}

func MapROrLToW(input *word.Word) *word.Word {
	return input.Replace(RORL_TO_W_LOWER, "w").Replace(RORL_TO_W_UPPER, "W")
}


func MapLlToWw(input *word.Word) *word.Word {
	return input.Replace(LL_TO_WW, "ww")
}

func MapVowelOrRExceptOLToWl(input *word.Word) *word.Word {
	return input.Replace(VOWEL_OR_R_EXCEPT_O_L_TO_WL_LOWER, "wl").Replace(VOWEL_OR_R_EXCEPT_O_L_TO_WL_UPPER, "W${1}")
}

func MapOldToOwld(input *word.Word) *word.Word {
	return input.Replace(OLD_TO_OWLD_LOWER, "${1}wld").Replace(OLD_TO_OWLD_UPPER, "OWLD")
}

func MapOlToOwl(input *word.Word) *word.Word {
	return input.Replace(OL_TO_OWL_LOWER, "${1}wl").Replace(OL_TO_OWL_UPPER, "OWL")
}

func MapLOrROToWo(input *word.Word) *word.Word {
	return input.Replace(LORR_O_TO_WO_LOWER, "wo").Replace(LORR_O_TO_WO_UPPER, "W${1}")
}

func MapSpecificConsonantsOToLetterAndWo(input *word.Word) *word.Word {
	return input.Replace(SPECIFIC_CONSONANTS_O_TO_LETTER_AND_WO_LOWER, "${1}wo").ReplaceWithFuncMultiple(SPECIFIC_CONSONANTS_O_TO_LETTER_AND_WO_UPPER, func(str1 string, str2 string) string {
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
	return input.Replace(VORW_LE_TO_WAL, "wal")
}

func MapFiToFwi(input *word.Word) *word.Word {
	return input.Replace(FI_TO_FWI_LOWER, "${1}wi").Replace(FI_TO_FWI_UPPER, "FWI")
}

func MapVerToWer(input *word.Word) *word.Word {
	return input.Replace(VER_TO_WER, "wer")
}

func MapPoiToPwoi(input *word.Word) *word.Word {
	return input.Replace(POI_TO_PWOI, "${1}woi")
}

func MapSpecificConsonantsLeToLetterAndWal(input *word.Word) *word.Word {
	return input.Replace(SPECIFIC_CONSONANTS_LE_TO_LETTER_AND_WAL, "${1}wal")
}

func MapConsonantsRToConsonantW(input *word.Word) *word.Word {
	return input.Replace(CONSONANT_R_TO_CONSONANT_W, "${1}w")
}

func MapLyToWy(input *word.Word) *word.Word {
	return input.Replace(LY_TO_WY_LOWER, "wy").Replace(LY_TO_WY_UPPER, "Wy")
}

func MapPleToPwe(input *word.Word) *word.Word {
	return input.Replace(PLE_TO_PWE, "${1}we")
}

func MapNrToNw(input *word.Word) *word.Word {
	return input.Replace(NR_TO_NW_LOWER, "nw").Replace(NR_TO_NW_UPPER, "NW")
}

func MapFucToFwuc(input *word.Word) *word.Word {
	return input.Replace(FUC_TO_FWUC, "${1}wuc")
}

func MapMomToMwom(input *word.Word) *word.Word {
	return input.Replace(MOM_TO_MWOM, "${1}wom")
}

func MapMeToMwe(input *word.Word) *word.Word {
	return input.Replace(ME_TO_MWE, "${1}we")
}

func MapNVowelToNy(input *word.Word) *word.Word {
	return input.Replace(N_VOWEL_TO_NY_FIRST, "ny${1}").Replace(N_VOWEL_TO_NY_SECOND, "Ny${1}").Replace(N_VOWEL_TO_NY_THIRD, "NY${1}")
}

func MapOveToUv(input *word.Word) *word.Word {
	return input.Replace(OVE_TO_UV_LOWER, "uv").Replace(OVE_TO_UV_UPPER, "UV")
}

func MapHahaToHeheXd(input *word.Word) *word.Word {
	return input.Replace(HAHA_TO_HEHE_XD, "hehe xD")
}

func MapTheToTeh(input *word.Word) *word.Word {
	return input.Replace(THE_TO_TEH, "${1}eh")
}

func MapYouToU(input *word.Word) *word.Word {
	return input.Replace(YOU_TO_U_UPPER, "U").Replace(YOU_TO_U_LOWER, "u")
}

func MapTimeToTim(input *word.Word) *word.Word {
	return input.Replace(TIME_TO_TIM, "${1}im")
}

func MapOverToOwor(input *word.Word) *word.Word {
	return input.Replace(OVER_TO_OWOR, "${1}wor")
}

func MapWorseToWose(input *word.Word) *word.Word {
	return input.Replace(WORSE_TO_WOSE, "${1}ose")
}