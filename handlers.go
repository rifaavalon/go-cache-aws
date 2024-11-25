package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	mux "github.com/julienschmidt/httprouter"
)

func LogHttp(r *http.Request) {
	start := time.Now()
	log.Printf("%s\t%s\t%q\t%s", r.Method, r.RequestURI, r.Header, time.Since(start))
}

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	LogHttp(r)
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Welcome to the Xero AWS Cache</h1>")
}

func InstanceIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	LogHttp(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	instances := FindAll()

	if err := json.NewEncoder(w).Encode(instances); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func InstanceShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	LogHttp(r)
	id, err := strconv.Atoi(ps.ByName("instanceIds"))
	if err != nil {
		http.Error(w, "Invalid instance ID", http.StatusBadRequest)
		return
	}

	instance := FindInstance(id)

	if err := json.NewEncoder(w).Encode(instance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func InstanceCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var instance Instance
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &instance); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err1 := json.NewEncoder(w).Encode(err); err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
		}
		return
	}

	CreateInstance(instance)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
