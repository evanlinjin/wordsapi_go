package wordsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Response structure for 'Details Response'.
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
}

// Fills the DetailsResponse with data.
func (rm *DetailsResponse) fill(path string, queries ...[2]string) (err error) {
	queryMap := url.Values{}
	for _, query := range queries {
		queryMap.Set(query[0], query[1])
	}
	u := url.URL{
		Scheme:   "https",
		Host:     "wordsapiv1.p.mashape.com",
		Path:     "words/" + path,
		RawQuery: queryMap.Encode(),
	}

	fmt.Println("URL:", u.String())
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("X-Mashape-Key", key)
	req.Header.Add("Accept", "application/json")

	resp, _ := client.Do(req)
	if resp.StatusCode != 200 {
		err = &ServerResponseError{resp.StatusCode}
	}

	var p []byte
	switch resp.ContentLength {
	case -1:
		p = manualReadHttpResponse(resp)
	default:
		p = make([]byte, resp.ContentLength)
		resp.Body.Read(p)
	}
	json.Unmarshal(p, &rm)

	return
}

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
