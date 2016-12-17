package wordsapi

import (
	"strconv"
	"strings"
)

func GetRandom() (r WordResponse, e error) {
	return r, r.fill("", [2]string{"random", "true"})
}

func GetSearch(p SearchParameters) (r WordResponse, e error) {
	e = p.Check()
	if e != nil {
		return
	}
	return r, r.fill(
		"",
		[2]string{"frequencymax", strconv.FormatFloat(float64(p.FrequencyMax), 'f', 2, 32)},
		[2]string{"frequencymin", strconv.FormatFloat(float64(p.FrequencyMin), 'f', 2, 32)},
		[2]string{"hasDetails", strings.Join(p.HasDetails, ",")},
		[2]string{"frequencymax", p.LetterPattern},
		[2]string{"frequencymax", strconv.Itoa(int(p.Letters))},
		[2]string{"frequencymax", strconv.Itoa(int(p.LettersMax))},
		[2]string{"frequencymax", strconv.Itoa(int(p.LettersMin))},
		[2]string{"frequencymax", strconv.Itoa(int(p.Limit))},
		[2]string{"frequencymax", strconv.Itoa(int(p.Page))},
		[2]string{"frequencymax", p.PartOfSpeech},
		[2]string{"frequencymax", p.PronunciationPattern},
		[2]string{"frequencymax", strconv.Itoa(int(p.Sounds))},
		[2]string{"frequencymax", strconv.Itoa(int(p.SoundsMax))},
		[2]string{"frequencymax", strconv.Itoa(int(p.SoundsMin))},
		[2]string{"frequencymax", strconv.Itoa(int(p.Syllables))},
		[2]string{"frequencymax", strconv.Itoa(int(p.SyllablesMax))},
		[2]string{"frequencymax", strconv.Itoa(int(p.SyllablesMin))},
	)
}

func GetWord(word string) (r WordResponse, e error) {
	return r, r.fill(word)
}
