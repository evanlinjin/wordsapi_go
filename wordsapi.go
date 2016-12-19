package wordsapi

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var key string = ""
var client *http.Client = &http.Client{}

// Sets the api key to 'str'.
func SetApiKey(str string) { key = str }

func SetApiKeyFromFile() {
	fileData, e := ioutil.ReadFile("api.key")
	if e != nil {
		log.Fatalln(e)
		return
	}
	key = strings.TrimSpace(string(fileData))
}

type Response interface {
	fill(string, ...[2]string) error
}

// Manually reads in the http.Response data if ContentLength is undefined.
func manualReadHttpResponse(resp *http.Response) (out []byte) {
	p := make([]byte, 1)
	for {
		_, e := resp.Body.Read(p)
		if e != nil {
			break
		}
		out = append(out, p[0])
	}
	return
}
