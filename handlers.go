package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func getInt64(r *http.Request, key string) (id int64, ok bool) {
	v := chi.URLParam(r, key)
	if v == "" {
		ok = false
		return
	}

	var err error
	id, err = strconv.ParseInt(v, 10, 64)
	if err != nil {
		return
	}
	ok = true
	return
}

func jsonResponse(w http.ResponseWriter, v interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		fmt.Println("For some reason writing response failed:")
		fmt.Println(err.Error())
		fmt.Println()
	}
}

func jsonRequest(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		fmt.Println("Failed to unmarshal request body JSON:")
		fmt.Println(err.Error())
		fmt.Println()
	}
	return err
}
