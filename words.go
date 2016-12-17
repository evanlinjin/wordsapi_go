package wordsapi

import (
	"strconv"
	"strings"
)

// Get antonyms (opposites) of a word.
func GetAntonyms(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/antonyms")
}

// Get definitions of a word, including the part of speech.
func GetDefinitions(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/definitions")
}

// Get examples of how the word is used.
func GetExamples(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/examples")
}

// Expands upon the frequeny score returned by the main /words/{word} endpoint.
// Returns zipf, a score indicating how common the word is in the English
// language, with a range of 1 to 7; perMillion, the number of times the word
// is likely to appear in a corpus of one million English words; and diversity,
// a 0-1 scale the shows the likelyhood of the word appearing in an English
// document that is part of a corpus.
func GetFrequency(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/frequency")
}

// How to pronounce a word, according to the International Phonetic Alphabet.
// May include multiple results if the word is pronounced differently depending
// on its part of speech.
func GetPronunciation(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/pronunciation")
}

// Retrieve a random word, optionally matching a search criteria. You can use
// the same search criteria as the "Search" endpoint.
func GetRandom() (r WordResponse, e error) {
	return r, r.fill("", [2]string{"random", "true"})
}

// Get a list of words that rhyme with the given word.
func GetRhymes(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/rhymes")
}

// Search for words matching the parameters you provide.
// https://www.wordsapi.com/docs#search
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

// Returns the word broken down into syllables.
func GetSyllables(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/syllables")
}

// Get synonyms of a word.
func GetSynonyms(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/synonyms")
}

// Retrieve information about a word. Results can include definitions,
// part of speech, synonyms, related words, syllables, and pronunciation.
// This method is useful to see which relationships are attached to which
// definition and part of speech of a word.
func GetWord(word string) (r WordResponse, e error) {
	return r, r.fill(word)
}
