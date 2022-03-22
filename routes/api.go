package routes

import (
	"go-rest-api/controllers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var ApiRoutes = Routes{
	Route{
		"PersonalidadesIndexRedirect",
		"GET",
		"/",
		controllers.Home,
	},
	Route{
		"PersonalidadesIndex",
		"GET",
		"/api/personalidades",
		controllers.TodasPersonalidades,
	},
	Route{
		"PersonalidadeShow",
		"GET",
		"/api/personalidades/{id}",
		controllers.UnicaPersonalidade,
	},
	Route{
		"PersonalidadeAdd",
		"POST",
		"/api/personalidades",
		controllers.NovaPersonalidade,
	},
	Route{
		"PersonalidadeEdit",
		"PUT",
		"/api/personalidades/{id}",
		controllers.AtualizaPersonalidade,
	},
	Route{
		"PersonalidadeDelete",
		"DELETE",
		"/api/personalidades/{id}",
		controllers.DeletaPersonalidade,
	},
}

type Routes []Route
