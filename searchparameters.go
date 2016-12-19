package wordsapi

const LO_FREQ_SCORE float32 = 1.74
const HI_FREQ_SCORE float32 = 8.03

var detail_types = [...]string{
	"hasCategories",
	"hasInstances",
	"hasMembers",
	"hasParts",
	"hasSubstances",
	"hasTypes",
	"hasUsages",
	"inCategory",
	"inRegion",
	"typeOf",
	"InstanceOf",
	"MemberOf",
	"PartOf",
	"PertainsTo",
	"RegionOf",
	"SimilarTo",
	"SubstanceOf",
	"UsageOf",
}

func IsValidDetalType(str string) bool {
	for _, v := range detail_types {
		if str == v {
			return true
		}
	}
	return false
}

func MakeSearchParameters() (p SearchParameters) {
	p.Fix()
	return
}

type SearchParameters struct {
	FrequencyMax         float32  // The max freq score. RANGE: 1.74 - 8.03.
	FrequencyMin         float32  // The min freq score. RANGE: 1.74 - 8.03.
	HasDetails           []string // List of detail types the word has.
	LetterPattern        string   // RE the word's letters need to match.
	Letters              uint     // Number of letters the word must have.
	LettersMax           uint     // Max number of letters the word can have.
	LettersMin           uint     // Min number of letters the word can have.
	Limit                uint     // Number of results to display per page.
	Page                 uint     // Page of results to return. DEFAULT: 1.
	PartOfSpeech         string   // Word must have 1 definition with this POS.
	PronunciationPattern string   // RE the word's pronounciation needs to match.
	Sounds               uint     // Number of phonemes the word must have.
	SoundsMax            uint     // Max number of phonemes the word can have.
	SoundsMin            uint     // Min number of phonemes the word can have.
	Syllables            uint     // Number of syllables the word must have.
	SyllablesMax         uint     // Max number of syllables the word can have.
	SyllablesMin         uint     // Min number of syllables the word can have.
}

// Sets the FrequencyMax parameter.
func (p *SearchParameters) SetFrequencyMax(v float32) (e error) {
	if !isOfRange(v, LO_FREQ_SCORE, HI_FREQ_SCORE) {
		e = &ParameterError{
			pType:     "FrequencyMax",
			pCrrValue: p.FrequencyMax,
			pErrValue: v,
			pValidMin: LO_FREQ_SCORE,
			pValidMax: HI_FREQ_SCORE,
		}
		return
	}
	p.FrequencyMax = v
	if p.FrequencyMin > p.FrequencyMax {
		p.FrequencyMin = p.FrequencyMax
	}
	return
}

// Sets the FrequencyMin parameter.
func (p *SearchParameters) SetFrequencyMin(v float32) (e error) {
	if !isOfRange(v, LO_FREQ_SCORE, HI_FREQ_SCORE) {
		e = &ParameterError{
			pType:     "FrequencyMin",
			pCrrValue: p.FrequencyMin,
			pErrValue: v,
			pValidMin: LO_FREQ_SCORE,
			pValidMax: HI_FREQ_SCORE,
		}
		return
	}
	p.FrequencyMin = v
	if p.FrequencyMin > p.FrequencyMax {
		p.FrequencyMax = p.FrequencyMin
	}
	return
}

// Sets the HasDetails parameter.
func (p *SearchParameters) SetHasDetails(details ...string) (e error) {
	validDetails, allDetails := []string{}, []string{}
	for _, d := range details {
		if IsValidDetalType(d) == true {
			validDetails = append(validDetails, d)
			allDetails = append(allDetails, d)
		} else {
			allDetails = append(allDetails, "(INVALID)"+d)
		}
	}
	p.HasDetails = validDetails
	if len(validDetails) != len(allDetails) {
		e = &ParameterError{
			pType:     "HasDetails",
			pCrrValue: p.HasDetails,
			pErrValue: allDetails,
			pValidMax: detail_types,
		}
	}
	return
}

// Sets the LetterPattern parameter.
func (p *SearchParameters) SetLetterPattern(v string) (e error) {
	p.LetterPattern = v
	return
}

// Sets the Letters, LettersMax and LettersMin parameters.
func (p *SearchParameters) SetLetters(vals ...uint) (e error) {
	switch len(vals) {
	case 1:
		p.Letters, p.LettersMax, p.LettersMin = vals[0], 0, 0
	case 2:
		if p.Letters = 0; vals[0] > vals[1] {
			p.LettersMax, p.LettersMin = vals[0], vals[1]
		} else {
			p.LettersMax, p.LettersMin = vals[1], vals[0]
		}
	default:
		e = &ParameterError{
			pType:     "Letters",
			pCrrValue: []uint{p.Letters, p.LettersMax, p.LettersMin},
			pErrValue: vals,
		}
	}
	return
}

// Checks SearchParameters to see if parameters are valid.
func (p *SearchParameters) Check() (e error) {
	freqMaxOk := isOfRange(p.FrequencyMax, LO_FREQ_SCORE, HI_FREQ_SCORE)
	freqMinOk := isOfRange(p.FrequencyMin, LO_FREQ_SCORE, HI_FREQ_SCORE)
	pageOk := p.Page > 0

	if !(freqMaxOk && freqMinOk && pageOk) {
		e = &SearchParametersError{freqMaxOk, freqMinOk, pageOk}
	}
	return
}

// Fix SearchParameters to make all parameters valid.
func (p *SearchParameters) Fix() {
	switch {

	case p.FrequencyMax != 0:
		if p.FrequencyMax < LO_FREQ_SCORE {
			p.FrequencyMax = LO_FREQ_SCORE
		} else if p.FrequencyMax > HI_FREQ_SCORE {
			p.FrequencyMax = HI_FREQ_SCORE
		}
		fallthrough

	case p.FrequencyMin != 0:
		if p.FrequencyMin < LO_FREQ_SCORE {
			p.FrequencyMin = LO_FREQ_SCORE
		} else if p.FrequencyMin > HI_FREQ_SCORE {
			p.FrequencyMin = HI_FREQ_SCORE
		}
		fallthrough

	default:
		if p.Page < 1 {
			p.Page = 1
		}
	}
}

func isOfRange(v, rL, rH float32) bool {
	return (v == 0) || (v >= rL && v <= rH)
}
