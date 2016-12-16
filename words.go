package wordsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Result struct {
	Definition   string   `json:"definition"`
	Derivation   []string `json:"derivation"`
	PartOfSpeech string   `json:"partOfSpeech"`
	Synonyms     []string `json:"synonyms"`
	TypeOf       []string `json:"typeOf"`
}

type WordResponse struct {
	Word    string   `json:"word"`
	Results []Result `json:"results"`
}

func (rm *WordResponse) fill(path string, query url.Values) (err error) {
	u := url.URL{
		Scheme:   "https",
		Host:     "wordsapiv1.p.mashape.com",
		Path:     "words/" + path,
		RawQuery: query.Encode(),
	}
	fmt.Println(u.String())

	req, _ := http.NewRequest("GET", u.String(), nil)
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

func GetRandom() (r WordResponse, e error) {
	q := url.Values{}
	q.Set("random", "true")
	return r, r.fill("", q)
}
