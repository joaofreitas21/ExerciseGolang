package routes

import (
	"net/http"
	"teste1/src/api/controllers"
)

var loginRoutes = []Route{
	Route{
		Url:"/api/login",
		Method:  http.MethodPost,
		Handler: controllers.Login,
		AuthRequired: false,
	},
}

