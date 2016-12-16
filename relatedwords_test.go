package wordsapi

import "testing"

type Test struct {
	Word string
}

func TestAlso(t *testing.T) {
	cases := []Test{
		{"endpoint"},
	}

	for _, c := range cases {
		SetApiKey("")
		got, err := GetAlso(c.Word)
		if err != nil {
			t.Errorf("GetAlso(%q) -> %q!", c.Word, err)
		} else if got.Word != c.Word {
			t.Errorf("GetAlso(%q) -> %q!", c.Word, got.Word)
		}
	}
}
