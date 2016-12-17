package wordsapi

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

type Test struct {
	Word string
}

func TestGetAlso(t *testing.T) {
	cases := []Test{
		{"endpoint"},
	}

	for _, c := range cases {
		SetApiKey(GetKeyFromFile())
		got, err := GetAlso(c.Word)
		if err != nil || got.Word != c.Word {
			t.Errorf("GetAlso(%q) -> %q ! ERR: %q", c.Word, got.Word, err)
		}
	}
}

func TestGetRandom(t *testing.T) {
	SetApiKey(GetKeyFromFile())
	got, err := GetRandom()
	if err != nil {
		t.Errorf("GetRandom() -> %q ! ERR: %q", got.Word, err)
	}
}

func GetKeyFromFile() (key string) {
	fileData, e := ioutil.ReadFile("api.key")
	if e != nil {
		log.Fatalln(e)
		return
	}
	key = strings.TrimSpace(string(fileData))
	return
}
