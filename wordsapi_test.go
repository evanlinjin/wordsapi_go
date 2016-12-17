package wordsapi

import "testing"

type Test struct {
	Word string
}

func TestGetAlso(t *testing.T) {
	cases := []Test{
		{"endpoint"},
	}

	for _, c := range cases {
		SetApiKeyFromFile()
		got, err := GetAlso(c.Word)
		if err != nil || got.Word != c.Word {
			t.Errorf("GetAlso(%q) -> %q ! ERR: %q", c.Word, got.Word, err)
		}
	}
}

func TestGetRandom(t *testing.T) {
	SetApiKeyFromFile()
	got, err := GetRandom()
	if err != nil {
		t.Errorf("GetRandom() -> %q ! ERR: %q", got.Word, err)
	}
}
