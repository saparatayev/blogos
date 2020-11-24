package routes

import (
	"blogos/src/api/controllers"
	"net/http"
)

var loginRoutes = []Route{
	Route{
		Uri:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}
