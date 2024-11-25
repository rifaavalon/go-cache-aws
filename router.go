package main

import (
	"github.com/julienschmidt/httprouter"
)

// Route defines a single route, e.g. a human readable name, HTTP method, and the pattern the function will be executed on.
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

// Routes is a slice of Route structs.
type Routes []Route

// NewRouter creates a new router and registers the routes.
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}

// Define your routes here
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"InstanceIndex",
		"GET",
		"/instances",
		InstanceIndex,
	},
	Route{
		"InstanceShow",
		"GET",
		"/instances/:id",
		InstanceShow,
	},
	Route{
		"InstanceCreate",
		"POST",
		"/instances",
		InstanceCreate,
	},
}
