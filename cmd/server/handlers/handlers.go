package handlers

import (
	"net/http"
	"bytes"
	"encoding/json"
)

func writeJSONValue(w http.ResponseWriter, data interface{}, httpStatus int) {
	writeToResponse(w, data, httpStatus)
}

func writeJSONError(w http.ResponseWriter, err error, httpStatus int) {
	e := struct {
		Message string `json:"message"`
	}{err.Error()}
	writeToResponse(w, e, httpStatus)
}

func writeToResponse(w http.ResponseWriter, data interface{}, httpStatus int) {
	bs := bytes.Buffer{}
	enc := json.NewEncoder(&bs)
	enc.Encode(data)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)
	bs.WriteTo(w)
}
