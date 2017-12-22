package handlers

import (
	"net/http"
	"bytes"
	"encoding/json"
)

func toJson(w http.ResponseWriter, data interface{}) {
	bs := bytes.Buffer{}
	e := json.NewEncoder(&bs)
	e.Encode(data)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	bs.WriteTo(w)
}

func toServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
