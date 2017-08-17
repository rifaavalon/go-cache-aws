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
	fmt.Fprintf(w, "<h1 style=\"font-familyt: Helvetica;\">Welcome to the Xero AWS Cache</h1>")
}

func InstanceIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	LogHttp(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	instances := FindAll()

	if err := json.NewEncoder(w).Encode(instances); err != nil {
		panic(err)
	}
}

func InstanceShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	LogHttp(r)
	id, err := strconv.Atoi(ps.ByName("instanceIds"))
	if err != nil {
		panic(err)
	}

	instance := FindInstance(id)

	if err := json.NewEncoder(w).Encode(instance); err != nil {

	}
}

func InstanceCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var instance Instance
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &instance); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		if err1 := json.NewEncoder(w).Encode(err); err1 != nil {
			panic(err)
		}
	}
	CreateInstance(instance)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}
