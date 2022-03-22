package helpers

import (
	"go-rest-api/middleware"
	"go-rest-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.ContentTypeMiddleware)
	for _, route := range routes.ApiRoutes {
		var handler http.Handler
	
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	
	return router
}
