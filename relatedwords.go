package wordsapi

import (
	"encoding/json"
	"net/http"
)

type Definition struct {
	Definition   string `json:"definition"`
	PartOfSpeech string `json:"partOfSpeech"`
}

type Frequency struct {
	Zipf       float32 `json:"zipf"`
	PerMillion float32 `json:"perMillion"`
	Diversity  float32 `json:"diversity"`
}

type Pronunciation struct {
	All  string `json:"all"`
	Noun string `json:"noun"`
	Verb string `json:"verb"`
}

type Rhymes struct {
	All []string `json:"all"`
}

type Syllables struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}

type DetailsResponse struct {
	Word          string        `json:"word"`
	Also          []string      `json:"also"`          //
	Entails       []string      `json:"entails"`       //
	HasCategories []string      `json:"hasCategories"` //
	HasInstances  []string      `json:"hasInstances"`  //
	HasMembers    []string      `json:"hasMembers"`    //
	HasParts      []string      `json:"hasParts"`      //
	HasSubstances []string      `json:"hasSubstances"` //
	HasTypes      []string      `json:"hasTypes"`      //
	HasUsages     []string      `json:"hasUsages"`     //
	InCategory    []string      `json:"inCategory"`    //
	InRegion      []string      `json:"inRegion"`      //
	TypeOf        []string      `json:"typeOf"`        //
	InstanceOf    []string      `json:"instanceOf"`    //
	MemberOf      []string      `json:"memberOf"`      //
	PartOf        []string      `json:"partOf"`        //
	PertainsTo    []string      `json:"pertainsTo"`    //
	RegionOf      []string      `json:"regionOf"`      //
	SimilarTo     []string      `json:"similarTo"`     //
	SubstanceOf   []string      `json:"substanceOf"`   //
	UsageOf       []string      `json:"usageOf"`       //
	Antonyms      []string      `json:"antonyms"`      //
	Definitions   []Definition  `json:"definitions"`   //
	Examples      []string      `json:"examples"`
	Frequency     Frequency     `json:"frequency"`
	Pronunciation Pronunciation `json:"pronunciation"`
	Rhymes        Rhymes        `json:"rhymes"`
	Syllables     Syllables     `json:"syllables"`
	Synonyms      []string      `json:"synonyms"` //
	// TODO: Word, Random, Search
}

func (rm *DetailsResponse) fill(url string) (err error) {
	url = "https://wordsapiv1.p.mashape.com/words/" + url

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Mashape-Key", key)
	req.Header.Add("Accept", "application/json")

	resp, _ := client.Do(req)
	if resp.StatusCode != 200 {
		err = &ServerResponseError{resp.StatusCode}
	}

	p := make([]byte, resp.ContentLength)
	resp.Body.Read(p)
	json.Unmarshal(p, &rm)

	return
}

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
