package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/fishies/controllers"
)

type Route struct {
	Name        string
	Method      string
	Url         string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	{"FishIndex", "GET", "/api/v1/fish", controllers.FishIndex},
	{"FishIndex", "POST", "/api/v1/fish", controllers.FishCreate},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Url).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
