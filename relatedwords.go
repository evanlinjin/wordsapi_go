package wordsapi

func GetAlso(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/also")
}

func GetEntails(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/entails")
}

func GetHasCategories(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasCategories")
}

func GetHasInstances(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasInstances")
}

func GetHasMembers(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasMembers")
}

func GetHasParts(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasParts")
}

func GetHasSubstances(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasSubstances")
}

func GetHasTypes(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasTypes")
}

func GetHasUsages(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasUsages")
}

func GetInCategory(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/inCategory")
}

func GetInRegion(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/inRegion")
}

func GetTypeOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/typeOf")
}

func GetInstanceOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/instanceOf")
}

func GetMemberOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/memberOf")
}

func GetPartOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/partOf")
}

func GetPertainsTo(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/pertainsTo")
}

func GetRegionOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/regionOf")
}

func GetSimilarTo(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/similarTo")
}

func GetSubstanceOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/substanceOf")
}

func GetUsageOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/usageOf")
}

func GetAntonyms(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/antonyms")
}

func GetDefinitions(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/definitions")
}

func GetExamples(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/examples")
}

func GetFrequency(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/frequency")
}

func GetPronunciation(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/pronunciation")
}

func GetRhymes(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/rhymes")
}

func GetSyllables(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/syllables")
}

func GetSynonyms(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/synonyms")
}
