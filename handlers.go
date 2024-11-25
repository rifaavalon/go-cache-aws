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

// LogHttp logs the details of the HTTP request
func LogHttp(r *http.Request) {
	start := time.Now()
	log.Printf("%s\t%s\t%q\t%s", r.Method, r.RequestURI, r.Header, time.Since(start))
}

// Index handles the root endpoint
func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	LogHttp(r)
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Welcome to the Xero AWS Cache</h1>")
}

// InstanceIndex handles the /instances endpoint
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

// InstanceShow handles the /instances/:id endpoint
func InstanceShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	LogHttp(r)
	id, err := strconv.Atoi(ps.ByName("instanceId"))
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

// InstanceCreate handles the creation of a new instance
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
