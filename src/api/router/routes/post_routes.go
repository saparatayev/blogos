package routes

import (
	"blogos/src/api/controllers"
	"net/http"
)

var postsRoutes = []Route{
	Route{
		Uri:     "/posts",
		Method:  http.MethodGet,
		Handler: controllers.GetPosts,
	},
	Route{
		Uri:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
	},
}
