package wordsapi

type SearchParameters struct {
	FrequencyMax         float32  // The max freq score. Range: 1.74 - 8.03.
	FrequencyMin         float32  // The min freq score. Range: 1.74 - 8.03.
	HasDetails           []string // List of detail types the word has.
	LetterPattern        string   // RE the word's letters need to match.
	Letters              uint     // Number of letters the word must have.
	LettersMax           uint     // Max number of letters the word can have.
	LettersMin           uint     // Min number of letters the word can have.
	Limit                uint     // Number of results to display per page.
	Page                 uint     // Page of results to return. Default: 1.
	PartOfSpeech         string   // Word must have 1 definition with this POS.
	PronunciationPattern string   // RE the word's pronounciation needs to match.
	Sounds               uint     // Number of phonemes the word must have.
	SoundsMax            uint     // Max number of phonemes the word can have.
	SoundsMin            uint     // Min number of phonemes the word can have.
	Syllables            uint     // Number of syllables the word must have.
	SyllablesMax         uint     // Max number of syllables the word can have.
	SyllablesMin         uint     // Min number of syllables the word can have.
}

func (p *SearchParameters) Check() (e error) {
	isOfRange := func(v, rL, rH float32) bool {
		return (v == 0) || (v >= rL && v <= rH)
	}

	freqMaxOk := isOfRange(p.FrequencyMax, 1.74, 8.03)
	freqMinOk := isOfRange(p.FrequencyMin, 1.74, 8.03)
	pageOk := p.Page > 0

	if !(freqMaxOk && freqMinOk && pageOk) {
		e = &SearchParametersError{freqMaxOk, freqMinOk, pageOk}
	}
	return
}
