package wordsapi

const LO_FREQ_SCORE float32 = 1.74
const HI_FREQ_SCORE float32 = 8.03

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

// Checks SearchParameters to see if parameters are valid.
func (p *SearchParameters) Check() (e error) {
	isOfRange := func(v, rL, rH float32) bool {
		return (v == 0) || (v >= rL && v <= rH)
	}

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
