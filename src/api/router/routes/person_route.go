package routes

import (
	"net/http"
	"teste1/src/api/controllers"
)

var personRoutes = []Route{
	Route{
		Url:    "/api/persons/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetPerson,
		AuthRequired: true,
	},
	Route{
		Url:    "/api/create",
		Method: http.MethodPost,
		Handler: controllers.CreatePerson,
		AuthRequired: true,
	},
}