package router

import (
	"github.com/Flyingmonk01/go-pro/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func SetupRouter() *mux.Router {
	router := NewRouter()

	// Applying the middleware for all the routes so get the desired logs when any api hits
	router.Use(middleware.RequestLogger)

	AuthRouter(router)
	UserRouter(router)
	return router
}
