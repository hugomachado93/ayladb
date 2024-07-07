package main

import (
	"encoding/json"
	"net/http"
)

type Api struct {
	*KeyValueStore
}

type Data struct {
	Key   string
	Value string
}

func (api *Api) SetValue(w http.ResponseWriter, r *http.Request) {
	var val Data
	json.NewDecoder(r.Body).Decode(&val)

	api.Set(val.Key, val.Value)

	w.WriteHeader(http.StatusOK)
}

func (api *Api) GetValue(w http.ResponseWriter, r *http.Request) {
	var val Data

	json.NewDecoder(r.Body).Decode(&val)
	value, err := api.Get(val.Key)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)
}
