package handlers

import (
	"net/http"
	"bytes"
	"encoding/json"
)

func toOk(w http.ResponseWriter, data interface{}) {
	bs := bytes.Buffer{}
	e := json.NewEncoder(&bs)
	e.Encode(data)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	bs.WriteTo(w)
}

func writeJsonError(w http.ResponseWriter, err error, httpStatus int) {
	e := struct {
		Message string `json:"message"`
	}{err.Error()}
	bs := bytes.Buffer{}
	enc := json.NewEncoder(&bs)
	enc.Encode(e)

	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	bs.WriteTo(w)
}
