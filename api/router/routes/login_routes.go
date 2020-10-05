package routes

import (
	"net/http"

	"github.com/alfredomendozacap/jwt-auth/api/controllers"
)

var loginRoutes = []Route{
	Route{
		URI:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
