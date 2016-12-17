package wordsapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type WordResponse struct {
	Word    string   `json:"word"`
	Results []Result `json:"results"`
}

func (rm *WordResponse) fill(path string, queries ...[2]string) (err error) {
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
		p = make([]byte, 2000)
	default:
		p = make([]byte, resp.ContentLength)
	}

	resp.Body.Read(p)
	json.Unmarshal(p, &rm)

	return
}

type Result struct {
	Definition   string   `json:"definition"`
	Derivation   []string `json:"derivation"`
	PartOfSpeech string   `json:"partOfSpeech"`
	Synonyms     []string `json:"synonyms"`
	TypeOf       []string `json:"typeOf"`
}
