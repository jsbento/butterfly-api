package api

import (
	"net/http"

	j "github.com/homelight/json"
)

func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	b, err := j.MarshalSafeCollections(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}
