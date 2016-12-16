package wordsapi

import "net/http"

var key string = ""
var client *http.Client = &http.Client{}

func SetApiKey(str string) { key = str }
