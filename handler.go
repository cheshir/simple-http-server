package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func handler(logger RequestLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			payload = []byte("error: " + err.Error())
		}

		payload = tryFormatJSON(payload)
		logger.Log(r.URL, r.Method, payload)
	}
}

func tryFormatJSON(payload []byte) []byte {
	var objmap map[string]json.RawMessage

	if err := json.Unmarshal(payload, &objmap); err != nil {
		println(err.Error())
		return payload
	}

	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(objmap); err != nil {
		println(err.Error())
		return payload
	}

	return buf.Bytes()
}
