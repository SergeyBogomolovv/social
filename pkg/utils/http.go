package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ParseJSON(r *http.Request, payload any) error {
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func GetIntQuery(param string, placeHolder int, r *http.Request) int {
	if r.URL.Query().Has(param) {
		param := r.URL.Query().Get(param)
		num, err := strconv.Atoi(param)
		if err == nil {
			return num
		}
	}

	return placeHolder
}

func GetIntParam(param string, r *http.Request) (int, error) {
	val := r.PathValue(param)
	num, err := strconv.Atoi(val)

	if err != nil {
		return 0, err
	}

	return num, nil
}
