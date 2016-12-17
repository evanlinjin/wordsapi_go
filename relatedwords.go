package wordsapi

// Phrases of which the word is a part.
func GetAlso(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/also")
}

// Words that are implied by the original word. Usually used for verbs.
func GetEntails(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/entails")
}

// Categories of the parameter word.
func GetHasCategories(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasCategories")
}

// Words that are examples of the parameter word.
func GetHasInstances(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasInstances")
}

// Words that belong to the group defined by the parameter word.
func GetHasMembers(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasMembers")
}

// Words that are parts of the original word.
func GetHasParts(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasParts")
}

// Words that are substances of the parameter word.
func GetHasSubstances(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasSubstances")
}

// Get more specific examples of types of this word.
func GetHasTypes(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasTypes")
}

// Words that are examples of the domain the original word defines.
func GetHasUsages(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/hasUsages")
}

// The domain category to which the original word belongs.
func GetInCategory(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/inCategory")
}

// Geographical areas where the word is used.
func GetInRegion(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/inRegion")
}

// Finds word that are more general than the given word.
func GetTypeOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/typeOf")
}

// Words that the parameter word is an example of.
func GetInstanceOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/instanceOf")
}

// A group to which the word belongs.
func GetMemberOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/memberOf")
}

// The larger whole to which the word belongs.
func GetPartOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/partOf")
}

// Words to which the original word is relevant.
func GetPertainsTo(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/pertainsTo")
}

// Words used in the specified geographical area.
func GetRegionOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/regionOf")
}

// Words that similar to the original word, but are not synonyms.
func GetSimilarTo(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/similarTo")
}

// Substances to which the original word is a part of.
func GetSubstanceOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/substanceOf")
}

// Words that the original word is a domain usage of.
func GetUsageOf(word string) (r DetailsResponse, e error) {
	return r, r.fill(word + "/usageOf")
}
