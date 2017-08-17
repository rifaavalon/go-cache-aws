package main

import mux "github.com/julienschmidt/httprouter"

//Route struct
type Route struct {
	Method  string
	Pattern string
	Name    string
	Handle  mux.Handle
}

//Routes slice
type Routes []Route

//NewRouter for the API
var routes = Routes{

	Route{
		"GET",
		"/",
		"Index",
		Index,
	},

	Route{
		"GET",
		"/instances",
		"InstanceIndex",
		InstanceIndex,
	},
	Route{

		"GET",
		"/instances/:id",
		"InstanceShow",
		InstanceShow,
	},
}
