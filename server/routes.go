package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"GetPageCount",
		"GET",
		"/page-count",
		GetPageCount,
	},
	Route{
		"GetItem",
		"GET",
		"/item",
		GetItem,
	},
	Route{
		"GetKind",
		"GET",
		"/kind/{categoryId}",
		GetKindByCategoryId,
	},
	Route{
		"GetDetailById",
		"GET",
		"/detail",
		GetDetailById,
	},
	Route{
		"AccessLog",
		"POST",
		"/log",
		AccessLog,
	},
	Route{
		"GetMenu",
		"GET",
		"/menu",
		GetMenu,
	},
	Route{
		"SendMail",
		"GET",
		"/send-verification-mail",
		RegistMail,
	},
}
