package wordsapi

func GetAlso(word string) (out Response, err error) {
	return out, out.fill(word + "/also")
}
